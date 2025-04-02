package main // ir-gen.go

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// CodeGenerator generates LLVM IR from an AST
type CodeGenerator struct {
	module     *ir.Module
	functions  map[string]*ir.Func
	variables  map[string]value.Value
	currentFn  *ir.Func
	currentBlk *ir.Block
	stringLit  map[string]*ir.Global
	nextTemp   int
}

// NewCodeGenerator creates a new code generator
func NewCodeGenerator(moduleName string) *CodeGenerator {
	return &CodeGenerator{
		module:    ir.NewModule(),
		functions: make(map[string]*ir.Func),
		variables: make(map[string]value.Value),
		stringLit: make(map[string]*ir.Global),
		nextTemp:  1,
	}
}

// Generate generates LLVM IR for a program
func (g *CodeGenerator) Generate(program *Program) (*ir.Module, error) {
	// Declare built-in functions
	g.declareBuiltins()
	
	// First pass: declare all functions
	for _, fn := range program.Functions {
		if err := g.declareFunction(fn); err != nil {
			return nil, err
		}
	}
	
	// Second pass: generate function bodies
	for _, fn := range program.Functions {
		if err := g.generateFunction(fn); err != nil {
			return nil, err
		}
	}
	
	return g.module, nil
}

// SaveModule saves the LLVM IR to a file
func SaveModule(m *ir.Module, filename string) error {
	// Convert module to string
	moduleStr := m.String()
	
	// Write to file
	return ioutil.WriteFile(filename, []byte(moduleStr), 0644)
}

// declareBuiltins declares built-in functions like print
func (g *CodeGenerator) declareBuiltins() {
	// Declare printf
	printf := g.module.NewFunc("printf", types.I32, ir.NewParam("format", types.NewPointer(types.I8)))
	printf.Sig.Variadic = true
	g.functions["printf"] = printf

	// Add print functions for each type
	g.addPrintFunction("print_int", types.I32)
	g.addPrintFunction("print_float", types.Float)
	g.addPrintFunction("print_bool", types.I1)
	
	// print_string is special
	g.addPrintStringFunction()
}

// addPrintFunction adds a print function for a specific type
func (g *CodeGenerator) addPrintFunction(name string, typ types.Type) {
	// Create function
	fn := g.module.NewFunc(name, types.Void, ir.NewParam("value", typ))
	g.functions[name] = fn
	
	// Create block for function
	block := fn.NewBlock("")
	
	// Create format string
	var format string
	if typ == types.I1 { // Boolean
		format = "%s\n"
	} else if typ == types.Float {
		format = "%f\n"
	} else {
		format = "%d\n"
	}
	
	// Temporarily set currentBlk so getStringLiteral can use it
	oldBlk := g.currentBlk
	g.currentBlk = block
	
	formatStr := g.getStringLiteral(format)
	
	// Call printf
	var args []value.Value
	args = append(args, formatStr)
	
	// For booleans, convert to string "true" or "false"
	if typ == types.I1 {
		// Create string constants
		trueStr := g.getStringLiteral("true")
		falseStr := g.getStringLiteral("false")
		
		// Create if/else selection
		result := block.NewSelect(fn.Params[0], trueStr, falseStr)
		args = append(args, result)
	} else {
		// For other types, pass the value directly
		args = append(args, fn.Params[0])
	}
	
	block.NewCall(g.functions["printf"], args...)
	block.NewRet(nil)
	
	// Restore original currentBlk
	g.currentBlk = oldBlk
}

// addPrintStringFunction adds a specialized print function for strings
func (g *CodeGenerator) addPrintStringFunction() {
	// Create function
	fn := g.module.NewFunc("print_string", types.Void, ir.NewParam("value", types.NewPointer(types.I8)))
	g.functions["print_string"] = fn
	
	block := fn.NewBlock("")
	
	// Temporarily set currentBlk so getStringLiteral can use it
	oldBlk := g.currentBlk
	g.currentBlk = block
	
	// Create format string
	formatStr := g.getStringLiteral("%s\n")
	
	// Call printf with the string
	block.NewCall(g.functions["printf"], formatStr, fn.Params[0])
	block.NewRet(nil)
	
	// Restore original currentBlk
	g.currentBlk = oldBlk
}

// getStringLiteral creates a global string constant
func (g *CodeGenerator) getStringLiteral(str string) value.Value {
	// Check if we already have this string
	if global, ok := g.stringLit[str]; ok {
		// If we're currently in a function block
		if g.currentBlk != nil {
			// Return a pointer to the first character of the string
			zero := constant.NewInt(types.I32, 0)
			return g.currentBlk.NewGetElementPtr(global.ContentType, global, zero, zero)
		}
		// Otherwise, just return the global (used during init)
		return global
	}
	
	// Add null terminator
	strWithNull := str + "\x00"
	
	// Create string constant
	constStr := constant.NewCharArrayFromString(strWithNull)
	
	// Create global variable for the string
	name := fmt.Sprintf(".str.%d", len(g.stringLit))
	global := g.module.NewGlobalDef(name, constStr)
	global.Immutable = true
	
	// Store in cache
	g.stringLit[str] = global
	
	// If we're currently in a function block
	if g.currentBlk != nil {
		// Return a pointer to the first character
		zero := constant.NewInt(types.I32, 0)
		return g.currentBlk.NewGetElementPtr(global.ContentType, global, zero, zero)
	}
	
	// Otherwise, just return the global (used during init)
	return global
}

// getNextTemp gets the next temporary variable name
func (g *CodeGenerator) getNextTemp() string {
	name := fmt.Sprintf("%%tmp%d", g.nextTemp)
	g.nextTemp++
	return name
}

// llvmType converts a language type to an LLVM type
func (g *CodeGenerator) llvmType(typeName string) types.Type {
	switch typeName {
	case "int":
		return types.I32
	case "float":
		return types.Float
	case "string":
		return types.NewPointer(types.I8) // char*
	case "bool":
		return types.I1
	case "void":
		return types.Void
	default:
		// Default to int for unknown types
		fmt.Printf("Warning: Unknown type %s, using i32\n", typeName)
		return types.I32
	}
}

// declareFunction declares a function (first pass)
func (g *CodeGenerator) declareFunction(fnDecl *FunctionDecl) error {
	// Convert parameter types
	params := make([]*ir.Param, 0, len(fnDecl.Parameters))
	for _, param := range fnDecl.Parameters {
		llvmType := g.llvmType(param.VarType)
		params = append(params, ir.NewParam(param.Name, llvmType))
	}
	
	// Convert return type
	retType := g.llvmType(fnDecl.ReturnType)
	
	// Create the function
	fn := g.module.NewFunc(fnDecl.Name, retType, params...)
	
	// Save for later use
	g.functions[fnDecl.Name] = fn
	
	return nil
}

// generateFunction generates code for a function body (second pass)
func (g *CodeGenerator) generateFunction(fnDecl *FunctionDecl) error {
	// Get the function
	fn := g.functions[fnDecl.Name]
	g.currentFn = fn
	
	// Create a new block
	entry := fn.NewBlock("")
	g.currentBlk = entry
	
	// Clear the local variables
	g.variables = make(map[string]value.Value)
	
	// Add parameters to variables
	for i, param := range fnDecl.Parameters {
		// Allocate memory for the parameter, but with a distinct name
		paramAlloca := g.currentBlk.NewAlloca(fn.Params[i].Type())
		
		// Use a different naming convention for the allocation to avoid conflict
		paramAlloca.SetName(param.Name + ".addr") // Add ".addr" suffix
		
		// Store the parameter value
		g.currentBlk.NewStore(fn.Params[i], paramAlloca)
		
		// Associate the variable name with the allocated memory
		g.variables[param.Name] = paramAlloca
	}
	
	// Generate code for the body
	if err := g.generateBlock(fnDecl.Body); err != nil {
		return err
	}
	
	// Add implicit return for void functions if no terminator
	if fn.Sig.RetType == types.Void && len(fn.Blocks) > 0 {
		lastBlock := fn.Blocks[len(fn.Blocks)-1]
		if lastBlock.Term == nil {
			lastBlock.NewRet(nil)
		}
	} else if len(fn.Blocks) > 0 {
		// Make sure non-void functions have a return
		lastBlock := fn.Blocks[len(fn.Blocks)-1]
		if lastBlock.Term == nil {
			// Default return value for non-void functions
			if fn.Sig.RetType == types.I32 {
				lastBlock.NewRet(constant.NewInt(types.I32, 0))
			} else if fn.Sig.RetType == types.Float {
				lastBlock.NewRet(constant.NewFloat(types.Float, 0.0))
			} else if fn.Sig.RetType == types.I1 {
				lastBlock.NewRet(constant.False)
			} else {
				// For other types (including pointers), return appropriate zero value
				if ptrType, ok := fn.Sig.RetType.(*types.PointerType); ok {
					lastBlock.NewRet(constant.NewNull(ptrType))
				} else {
					lastBlock.NewRet(constant.NewZeroInitializer(fn.Sig.RetType))
				}
			}
		}
	}
	
	return nil
}

// generateBlock generates code for a block of statements
func (g *CodeGenerator) generateBlock(block *Block) error {
	for _, stmt := range block.Statements {
		if err := g.generateStatement(stmt); err != nil {
			return err
		}
	}
	return nil
}

// generateStatement generates code for a statement
func (g *CodeGenerator) generateStatement(stmt Statement) error {
	switch stmt := stmt.(type) {
	case *VarDecl:
		return g.generateVarDecl(stmt)
	case *ReturnStatement:
		return g.generateReturn(stmt)
	case *IfStatement:
		return g.generateIf(stmt)
	case *WhileStatement:
		return g.generateWhile(stmt)
	case *AssignStatement:
		return g.generateAssign(stmt)
	case *ExpressionStatement:
		// Generate the expression but ignore its value
		expr := stmt.Expression
		
		// Special case for function calls that return void
		if callExpr, ok := expr.(*CallExpression); ok {
			// Check if it's a function call
			if fn, exists := g.functions[callExpr.Function]; exists && fn.Sig.RetType == types.Void {
				// Generate the call but don't try to use its void result
				_, err := g.generateCall(callExpr)
				return err
			}
		}
		
		// For other expressions, just generate and ignore the result
		_, err := g.generateExpression(expr)
		return err
	default:
		return fmt.Errorf("unsupported statement type: %T", stmt)
	}
}

// generateVarDecl generates code for a variable declaration
func (g *CodeGenerator) generateVarDecl(varDecl *VarDecl) error {
	// Allocate space for the variable on the stack
	var varType types.Type
	
	if varDecl.IsTypeInferred {
		// Type inference from the value
		if varDecl.Value == nil {
			return fmt.Errorf("cannot infer type for variable %s without initialization", varDecl.Name)
		}
		
		// Generate code for the value and use its type
		value, err := g.generateExpression(varDecl.Value)
		if err != nil {
			return err
		}
		
		varType = value.Type()
		
		// Allocate memory for the variable
		alloca := g.currentBlk.NewAlloca(varType)
		alloca.SetName(varDecl.Name)
		g.variables[varDecl.Name] = alloca
		
		// Store the initial value
		g.currentBlk.NewStore(value, alloca)
	} else {
		// Explicitly typed variable
		varType = g.llvmType(varDecl.VarType)
		
		// Allocate memory for the variable
		alloca := g.currentBlk.NewAlloca(varType)
		alloca.SetName(varDecl.Name)
		g.variables[varDecl.Name] = alloca
		
		// Store initial value if provided
		if varDecl.Value != nil {
			value, err := g.generateExpression(varDecl.Value)
			if err != nil {
				return err
			}
			
			// Check if we need type conversion
			if !types.Equal(value.Type(), varType) {
				value = g.convertValue(value, varType)
			}
			
			g.currentBlk.NewStore(value, alloca)
		}
	}
	
	return nil
}

// generateReturn generates code for a return statement
func (g *CodeGenerator) generateReturn(ret *ReturnStatement) error {
	// Check if returning void
	if g.currentFn.Sig.RetType == types.Void {
		g.currentBlk.NewRet(nil)
		return nil
	}
	
	// Generate the return value
	if ret.Value == nil {
		// Default return values for non-void functions
		if g.currentFn.Sig.RetType == types.I32 {
			g.currentBlk.NewRet(constant.NewInt(types.I32, 0))
		} else if g.currentFn.Sig.RetType == types.Float {
			g.currentBlk.NewRet(constant.NewFloat(types.Float, 0.0))
		} else if g.currentFn.Sig.RetType == types.I1 {
			g.currentBlk.NewRet(constant.False)
		} else {
			// For other types (including pointers), return appropriate zero value
			if ptrType, ok := g.currentFn.Sig.RetType.(*types.PointerType); ok {
				g.currentBlk.NewRet(constant.NewNull(ptrType))
			} else {
				g.currentBlk.NewRet(constant.NewZeroInitializer(g.currentFn.Sig.RetType))
			}
		}
		return nil
	}
	
	value, err := g.generateExpression(ret.Value)
	if err != nil {
		return err
	}
	
	// Convert return value if needed
	if !types.Equal(value.Type(), g.currentFn.Sig.RetType) {
		value = g.convertValue(value, g.currentFn.Sig.RetType)
	}
	
	g.currentBlk.NewRet(value)
	return nil
}

// generateIf generates code for an if statement
func (g *CodeGenerator) generateIf(ifStmt *IfStatement) error {
	// Generate condition
	condValue, err := g.generateExpression(ifStmt.Condition)
	if err != nil {
		return err
	}
	
	// Convert condition to boolean if needed
	if !types.Equal(condValue.Type(), types.I1) {
		condValue = g.convertToBool(condValue)
	}
	
	// Create blocks
	thenBlock := g.currentFn.NewBlock("")
	var elseBlock *ir.Block
	if ifStmt.ElseBlock != nil {
		elseBlock = g.currentFn.NewBlock("")
	}
	mergeBlock := g.currentFn.NewBlock("")
	
	// Branch based on condition
	if elseBlock != nil {
		g.currentBlk.NewCondBr(condValue, thenBlock, elseBlock)
	} else {
		g.currentBlk.NewCondBr(condValue, thenBlock, mergeBlock)
	}
	
	// Generate code for then block
	g.currentBlk = thenBlock
	if err := g.generateBlock(ifStmt.ThenBlock); err != nil {
		return err
	}
	if g.currentBlk.Term == nil {
		g.currentBlk.NewBr(mergeBlock)
	}
	
	// Generate code for else block if it exists
	if elseBlock != nil {
		g.currentBlk = elseBlock
		if err := g.generateBlock(ifStmt.ElseBlock); err != nil {
			return err
		}
		if g.currentBlk.Term == nil {
			g.currentBlk.NewBr(mergeBlock)
		}
	}
	
	// Continue at merge point
	g.currentBlk = mergeBlock
	return nil
}

// generateWhile generates code for a while statement
func (g *CodeGenerator) generateWhile(whileStmt *WhileStatement) error {
    // Create blocks
    condBlock := g.currentFn.NewBlock("while.cond")
    bodyBlock := g.currentFn.NewBlock("while.body")
    afterBlock := g.currentFn.NewBlock("while.end")
    
    // Branch to condition block
    g.currentBlk.NewBr(condBlock)
    
    // Generate condition code
    g.currentBlk = condBlock
    condValue, err := g.generateExpression(whileStmt.Condition)
    if err != nil {
        return err
    }
    
    // Convert condition to boolean if needed
    if !types.Equal(condValue.Type(), types.I1) {
        condValue = g.convertToBool(condValue)
    }
    
    // Branch based on condition (condition is true -> execute body, false -> exit loop)
    g.currentBlk.NewBr(afterBlock)
    
    // Generate loop body
    g.currentBlk = bodyBlock
    if err := g.generateBlock(whileStmt.Body); err != nil {
        return err
    }
    
    // Jump back to condition
    if g.currentBlk.Term == nil {
        g.currentBlk.NewBr(condBlock)
    }
    
    // Continue after loop
    g.currentBlk = afterBlock
    return nil
}

// generateAssign generates code for an assignment statement
func (g *CodeGenerator) generateAssign(assign *AssignStatement) error {
    // Check if variable exists
    variable, ok := g.variables[assign.Name]
    if !ok {
        return fmt.Errorf("undefined variable: %s", assign.Name)
    }
    
    // Generate value
    value, err := g.generateExpression(assign.Value)
    if err != nil {
        return err
    }
    
    // Check if variable is an alloca instruction
    alloca, ok := variable.(*ir.InstAlloca)
    if !ok {
        return fmt.Errorf("cannot assign to %s: not a variable", assign.Name)
    }
    
    // Check if we need to convert the value
    ptrType, ok := alloca.Type().(*types.PointerType)
    if !ok {
        return fmt.Errorf("internal error: expected pointer type for variable %s", assign.Name)
    }
    
    allocaType := ptrType.ElemType
    if !types.Equal(value.Type(), allocaType) {
        value = g.convertValue(value, allocaType)
    }
    
    // Store the value
    g.currentBlk.NewStore(value, alloca)
    return nil
}

// generateExpression generates code for an expression
func (g *CodeGenerator) generateExpression(expr Expression) (value.Value, error) {
	switch expr := expr.(type) {
	case *BinaryOp:
		return g.generateBinaryOp(expr)
	case *CallExpression:
		return g.generateCall(expr)
	case *Identifier:
		return g.generateIdentifier(expr)
	case *NumberLiteral:
		return constant.NewInt(types.I32, expr.Value), nil
	case *FloatLiteral:
		return constant.NewFloat(types.Float, expr.Value), nil
	case *StringLiteral:
		return g.getStringLiteral(expr.Value), nil
	case *BoolLiteral:
		if expr.Value {
			return constant.True, nil
		}
		return constant.False, nil
	default:
		return nil, fmt.Errorf("unsupported expression type: %T", expr)
	}
}

// generateBinaryOp generates code for a binary operation
func (g *CodeGenerator) generateBinaryOp(binOp *BinaryOp) (value.Value, error) {
	// Generate code for operands
	left, err := g.generateExpression(binOp.Left)
	if err != nil {
		return nil, err
	}
	
	right, err := g.generateExpression(binOp.Right)
	if err != nil {
		return nil, err
	}
	
	// Make sure types are compatible
	leftType := left.Type()
	rightType := right.Type()
	
	// Convert types if needed
	if types.Equal(leftType, types.I32) && types.Equal(rightType, types.Float) {
		left = g.currentBlk.NewSIToFP(left, types.Float)
		leftType = types.Float
	} else if types.Equal(leftType, types.Float) && types.Equal(rightType, types.I32) {
		right = g.currentBlk.NewSIToFP(right, types.Float)
		rightType = types.Float
	}
	
	// Handle operations based on operator
	switch binOp.Operator {
	case "+":
		if types.Equal(leftType, types.Float) {
			return g.currentBlk.NewFAdd(left, right), nil
		}
		return g.currentBlk.NewAdd(left, right), nil
		
	case "-":
		if types.Equal(leftType, types.Float) {
			return g.currentBlk.NewFSub(left, right), nil
		}
		return g.currentBlk.NewSub(left, right), nil
		
	case "*":
		if types.Equal(leftType, types.Float) {
			return g.currentBlk.NewFMul(left, right), nil
		}
		return g.currentBlk.NewMul(left, right), nil
		
	case "/":
		if types.Equal(leftType, types.Float) {
			return g.currentBlk.NewFDiv(left, right), nil
		}
		return g.currentBlk.NewSDiv(left, right), nil
		
	// Comparison operators
	case "==", "!=", "<", "<=", ">", ">=":
		return g.generateComparison(binOp.Operator, left, right)
		
	default:
		return nil, fmt.Errorf("unsupported binary operator: %s", binOp.Operator)
	}
}

// generateComparison generates code for comparison operations
func (g *CodeGenerator) generateComparison(op string, left, right value.Value) (value.Value, error) {
	// Use appropriate comparison based on type
	if types.Equal(left.Type(), types.Float) {
		// Float comparison
		var pred enum.FPred
		switch op {
		case "==":
			pred = enum.FPredOEQ
		case "!=":
			pred = enum.FPredONE
		case "<":
			pred = enum.FPredOLT
		case "<=":
			pred = enum.FPredOLE
		case ">":
			pred = enum.FPredOGT
		case ">=":
			pred = enum.FPredOGE
		}
		return g.currentBlk.NewFCmp(pred, left, right), nil
	} else {
		// Integer/pointer comparison
		var pred enum.IPred
		switch op {
		case "==":
			pred = enum.IPredEQ
		case "!=":
			pred = enum.IPredNE
		case "<":
			pred = enum.IPredSLT
		case "<=":
			pred = enum.IPredSLE
		case ">":
			pred = enum.IPredSGT
		case ">=":
			pred = enum.IPredSGE
		}
		return g.currentBlk.NewICmp(pred, left, right), nil
	}
}

// generateCall generates code for a function call
func (g *CodeGenerator) generateCall(callExpr *CallExpression) (value.Value, error) {
    // Check for special built-in functions
    if callExpr.Function == "printf" {
        return g.generatePrintfCall(callExpr)
    }
    
    // Check if it's a built-in print function
    if strings.HasPrefix(callExpr.Function, "print") && callExpr.Function != "printf" {
        return g.generatePrintCall(callExpr)
    }
    
    // Check if function exists
    fn, ok := g.functions[callExpr.Function]
    if !ok {
        return nil, fmt.Errorf("undefined function: %s", callExpr.Function)
    }
    
    // Generate code for arguments
    args := make([]value.Value, 0, len(callExpr.Arguments))
    for i, arg := range callExpr.Arguments {
        argValue, err := g.generateExpression(arg)
        if err != nil {
            return nil, err
        }
        
        // Convert type if needed
        if i < len(fn.Params) {
            paramType := fn.Params[i].Type()
            if !types.Equal(argValue.Type(), paramType) {
                argValue = g.convertValue(argValue, paramType)
            }
        }
        
        args = append(args, argValue)
    }
    
    // Call the function
    callInst := g.currentBlk.NewCall(fn, args...)
    
    // For void functions, return a dummy value that won't be used
    if types.Equal(fn.Sig.RetType, types.Void) {
        // Return constant int as placeholder that will be ignored
        return constant.NewInt(types.I32, 0), nil
    }
    
    return callInst, nil
}

// generatePrintfCall generates code for printf calls
func (g *CodeGenerator) generatePrintfCall(callExpr *CallExpression) (value.Value, error) {
    if len(callExpr.Arguments) < 1 {
        return nil, fmt.Errorf("printf requires at least a format string")
    }
    
    // Generate code for all arguments
    args := make([]value.Value, 0, len(callExpr.Arguments))
    
    // Handle the format string (first argument)
    formatArg, err := g.generateExpression(callExpr.Arguments[0])
    if err != nil {
        return nil, err
    }
    
    // Make sure the format is a pointer to i8
    if !g.isStringType(formatArg.Type()) {
        return nil, fmt.Errorf("printf first argument must be a string")
    }
    
    args = append(args, formatArg)
    
    // Add the rest of the arguments
    for i := 1; i < len(callExpr.Arguments); i++ {
        // Skip void function calls - they can't be used as arguments
        if callExpr, ok := callExpr.Arguments[i].(*CallExpression); ok {
            if fn, exists := g.functions[callExpr.Function]; exists && fn.Sig.RetType == types.Void {
                // Generate the call for its side effects, but don't use the result
                g.generateCall(callExpr)
                continue
            }
        }
        
        arg, err := g.generateExpression(callExpr.Arguments[i])
        if err != nil {
            return nil, err
        }
        
        // Skip void values (shouldn't normally happen)
        if types.Equal(arg.Type(), types.Void) {
            continue
        }
        
        args = append(args, arg)
    }
    
    // Call printf
    return g.currentBlk.NewCall(g.functions["printf"], args...), nil
}

// isStringType checks if a type is a string (pointer to i8)
func (g *CodeGenerator) isStringType(typ types.Type) bool {
    ptrType, ok := typ.(*types.PointerType)
    return ok && ptrType.ElemType == types.I8
}

// generatePrintCall generates code for print function calls
func (g *CodeGenerator) generatePrintCall(callExpr *CallExpression) (value.Value, error) {
	if len(callExpr.Arguments) != 1 {
		return nil, fmt.Errorf("print requires exactly one argument")
	}
	
	// Generate code for the argument
	arg, err := g.generateExpression(callExpr.Arguments[0])
	if err != nil {
		return nil, err
	}
	
	// Determine which print function to call based on type
	var printFunc *ir.Func
	
	switch {
	case types.Equal(arg.Type(), types.I32):
		printFunc = g.functions["print_int"]
	case types.Equal(arg.Type(), types.Float):
		printFunc = g.functions["print_float"]
	case types.Equal(arg.Type(), types.I1):
		printFunc = g.functions["print_bool"]
	default:
		// Check if it's a pointer to i8 (string)
		if ptrType, ok := arg.Type().(*types.PointerType); ok && ptrType.ElemType == types.I8 {
			printFunc = g.functions["print_string"]
		} else {
			// Default to printing as string
			arg = g.convertToString(arg)
			printFunc = g.functions["print_string"]
		}
	}
	
	// Call the print function and return a dummy value for void functions
	g.currentBlk.NewCall(printFunc, arg)
	return constant.NewInt(types.I32, 0), nil
}

// generateIdentifier generates code for a variable reference
func (g *CodeGenerator) generateIdentifier(ident *Identifier) (value.Value, error) {
	// Check if variable exists
	variable, ok := g.variables[ident.Name]
	if !ok {
		return nil, fmt.Errorf("undefined variable: %s", ident.Name)
	}
	
	// For allocated variables, load the value
	if alloca, ok := variable.(*ir.InstAlloca); ok {
		return g.currentBlk.NewLoad(alloca.Type().(*types.PointerType).ElemType, alloca), nil
	}
	
	// For parameters or other values, use directly
	return variable, nil
}

// convertValue converts a value from one type to another
func (g *CodeGenerator) convertValue(val value.Value, targetType types.Type) value.Value {
	sourceType := val.Type()
	
	// If types are already the same, no conversion needed
	if types.Equal(sourceType, targetType) {
		return val
	}
	
	// Integer to Float
	if types.Equal(sourceType, types.I32) && types.Equal(targetType, types.Float) {
		return g.currentBlk.NewSIToFP(val, types.Float)
	}
	
	// Float to Integer
	if types.Equal(sourceType, types.Float) && types.Equal(targetType, types.I32) {
		return g.currentBlk.NewFPToSI(val, types.I32)
	}
	
	// Integer to Bool
	if types.Equal(sourceType, types.I32) && types.Equal(targetType, types.I1) {
		return g.currentBlk.NewICmp(enum.IPredNE, val, constant.NewInt(types.I32, 0))
	}
	
	// Bool to Integer
	if types.Equal(sourceType, types.I1) && types.Equal(targetType, types.I32) {
		return g.currentBlk.NewZExt(val, types.I32)
	}
	
	// String conversion would be more complex, we'll return as is for now
	return val
}

// convertToBool converts a value to boolean
func (g *CodeGenerator) convertToBool(val value.Value) value.Value {
    // If already boolean, return as is
    if types.Equal(val.Type(), types.I1) {
        return val
    }
    
    // Integer comparison with zero
    if types.Equal(val.Type(), types.I32) {
        return g.currentBlk.NewICmp(enum.IPredNE, val, constant.NewInt(types.I32, 0))
    }
    
    // Float comparison with zero
    if types.Equal(val.Type(), types.Float) {
        return g.currentBlk.NewFCmp(enum.FPredONE, val, constant.NewFloat(types.Float, 0.0))
    }
    
    // Pointer comparison with null
    if ptrType, ok := val.Type().(*types.PointerType); ok {
        nullPtr := constant.NewNull(ptrType)
        return g.currentBlk.NewICmp(enum.IPredNE, val, nullPtr)
    }
    
    // Default to true for unknown types
    return constant.True
}

// convertToString converts a value to string representation
func (g *CodeGenerator) convertToString(val value.Value) value.Value {
	// For booleans, use "true" or "false"
	if types.Equal(val.Type(), types.I1) {
		trueStr := g.getStringLiteral("true")
		falseStr := g.getStringLiteral("false")
		return g.currentBlk.NewSelect(val, trueStr, falseStr)
	}
	
	// Default to empty string for other types
	emptyStr := g.getStringLiteral("")
	return emptyStr
}