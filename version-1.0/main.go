package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// Define our custom visitor implementation
type CustomVisitor struct {
	antlr.ParseTreeVisitor
	Module        *ir.Module
	CurrentFunc   *ir.Func
	CurrentBlock  *ir.Block
	SymbolTable   map[string]value.Value
	StringCounter int
	FuncMap       map[string]*ir.Func
}

func NewCustomVisitor() *CustomVisitor {
	// Add C standard library declarations
	mod := ir.NewModule()
	
	// printf declaration - variadic function
	printfParams := []*ir.Param{ir.NewParam("format", types.NewPointer(types.I8))}
	printf := mod.NewFunc("printf", types.I32, printfParams...)
	printf.Sig.Variadic = true
	
	// system declaration
	systemParams := []*ir.Param{ir.NewParam("command", types.NewPointer(types.I8))}
	system := mod.NewFunc("system", types.I32, systemParams...)

	return &CustomVisitor{
		Module:        mod,
		SymbolTable:   make(map[string]value.Value),
		StringCounter: 0,
		FuncMap: map[string]*ir.Func{
			"printf": printf,
			"system": system,
		},
	}
}

// Convert our language types to LLVM types
func getLLVMType(typeName string) types.Type {
	switch typeName {
	case "int":
		return types.I32
	case "float":
		return types.Float
	case "bool":
		return types.I1
	case "string":
		return types.NewPointer(types.I8)
	default:
		return types.Void
	}
}

// Create a global string constant
func (v *CustomVisitor) createStringConstant(text string) value.Value {
	if v.CurrentBlock == nil {
		fmt.Println("Error: No current block when creating string constant")
		return constant.NewNull(types.NewPointer(types.I8))
	}

	// Remove quotes and handle escape sequences
	text = strings.Trim(text, "\"")
	text = strings.ReplaceAll(text, "\\n", "\n")
	text = strings.ReplaceAll(text, "\\t", "\t")
	
	// Create a unique name for the string constant
	name := fmt.Sprintf(".str.%d", v.StringCounter)
	v.StringCounter++
	
	// Add null terminator
	stringVal := constant.NewCharArrayFromString(text + "\x00")
	
	// Create global variable
	global := v.Module.NewGlobalDef(name, stringVal)
	global.Immutable = true
	
	// Create pointer to the first character
	zero := constant.NewInt(types.I32, 0)
	return v.CurrentBlock.NewGetElementPtr(global.ContentType, global, zero, zero)
}

// Implementation of antlr.ParseTreeVisitor interface
func (v *CustomVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch ctx := tree.(type) {
	case *ProgramContext:
		return v.VisitProgram(ctx)
	case *FunctionContext:
		return v.VisitFunction(ctx)
	case *BlockContext:
		return v.VisitBlock(ctx)
	case *StatementContext:
		return v.VisitStatement(ctx)
	case *VariableDeclContext:
		return v.VisitVariableDecl(ctx)
	case *AssignmentContext:
		return v.VisitAssignment(ctx)
	case *ReturnStmtContext:
		return v.VisitReturnStmt(ctx)
	case *FunctionCallContext:
		return v.VisitFunctionCall(ctx)
	case *MulDivExprContext:
		return v.VisitMulDivExpr(ctx)
	case *AddSubExprContext:
		return v.VisitAddSubExpr(ctx)
	case *FunctionCallExprContext:
		return v.VisitFunctionCallExpr(ctx)
	case *VariableExprContext:
		return v.VisitVariableExpr(ctx)
	case *LiteralExprContext:
		return v.VisitLiteralExpr(ctx)
	case *ParenExprContext:
		return v.VisitParenExpr(ctx)
	default:
		fmt.Printf("Unhandled context type: %T\n", ctx)
		return nil
	}
}

func (v *CustomVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	var result interface{}
	for _, child := range node.GetChildren() {
		if childTree, ok := child.(antlr.ParseTree); ok {
			result = v.Visit(childTree)
		}
	}
	return result
}

func (v *CustomVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return node.GetText()
}

func (v *CustomVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}

// Visit program node
func (v *CustomVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	fmt.Println("Visiting program...")
	
	// Visit all function declarations
	for _, funcCtx := range ctx.AllFunction() {
		v.Visit(funcCtx)
	}
	
	return v.Module
}

// Visit function node
func (v *CustomVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	fmt.Println("Visiting function:", ctx.ID().GetText())
	
	// Get function name and return type
	funcName := ctx.ID().GetText()
	returnType := getLLVMType(ctx.Type_().GetText())
	
	// Process parameters
	var params []*ir.Param
	if paramsCtx := ctx.Parameters(); paramsCtx != nil {
		for _, paramCtx := range paramsCtx.AllParameter() {
			paramName := paramCtx.ID().GetText()
			paramType := getLLVMType(paramCtx.Type_().GetText())
			param := ir.NewParam(paramName, paramType)
			params = append(params, param)
		}
	}
	
	// Create the function or use existing one
	var f *ir.Func
	if existing, exists := v.FuncMap[funcName]; exists {
		// Keep existing function but update if needed
		f = existing
	} else {
		f = v.Module.NewFunc(funcName, returnType, params...)
		v.FuncMap[funcName] = f
	}
	
	v.CurrentFunc = f
	
	// Create entry block
	entryBlock := f.NewBlock("entry")
	v.CurrentBlock = entryBlock
	
	// Save old symbol table
	oldSymbolTable := v.SymbolTable
	v.SymbolTable = make(map[string]value.Value)
	
	// Create allocas for parameters and store parameter values in them
	for _, param := range f.Params {
		// Create a unique name for the alloca to avoid conflicts
		allocaName := fmt.Sprintf("%s.addr", param.Name())
		
		alloca := entryBlock.NewAlloca(param.Type())
		alloca.SetName(allocaName)
		entryBlock.NewStore(param, alloca)
		
		// Store the alloca in the symbol table with the original parameter name
		v.SymbolTable[param.Name()] = alloca
	}
	
	// Process function body
	v.Visit(ctx.Block())
	
	// Make sure the function has a return if needed
	if v.CurrentBlock.Term == nil {
		if returnType == types.Void {
			v.CurrentBlock.NewRet(nil)
		} else if intType, ok := returnType.(*types.IntType); ok {
			v.CurrentBlock.NewRet(constant.NewInt(intType, 0))
		} else if ptrType, ok := returnType.(*types.PointerType); ok {
			v.CurrentBlock.NewRet(constant.NewNull(ptrType))
		} else if returnType == types.Float {
			v.CurrentBlock.NewRet(constant.NewFloat(types.Float, 0.0))
		} else {
			v.CurrentBlock.NewRet(constant.NewInt(types.I32, 0))
		}
	}
	
	// Restore symbol table
	v.SymbolTable = oldSymbolTable
	
	return nil
}

// Visit block node
func (v *CustomVisitor) VisitBlock(ctx *BlockContext) interface{} {
	fmt.Println("Visiting block...")
	
	// Visit all statements
	for _, stmt := range ctx.AllStatement() {
		v.Visit(stmt)
		
		// If we've added a terminator, stop processing
		if v.CurrentBlock.Term != nil {
			break
		}
	}
	
	return nil
}

// Visit statement node - this is the new method we need
func (v *CustomVisitor) VisitStatement(ctx *StatementContext) interface{} {
	fmt.Println("Visiting statement...")
	
	// Check which child we have and visit it
	if varDecl := ctx.VariableDecl(); varDecl != nil {
		return v.Visit(varDecl)
	} else if assign := ctx.Assignment(); assign != nil {
		return v.Visit(assign)
	} else if funcCall := ctx.FunctionCall(); funcCall != nil {
		// This is a function call statement
		result := v.Visit(funcCall)
		return result
	} else if returnStmt := ctx.ReturnStmt(); returnStmt != nil {
		return v.Visit(returnStmt)
	}
	
	fmt.Println("Warning: Unhandled statement type")
	return nil
}

// Visit variable declaration
func (v *CustomVisitor) VisitVariableDecl(ctx *VariableDeclContext) interface{} {
	fmt.Println("Visiting variable declaration:", ctx.ID().GetText())
	
	var varType types.Type
	
	if typeCtx := ctx.Type_(); typeCtx != nil {
		varType = getLLVMType(typeCtx.GetText())
	} else {
		// For 'var' declarations, default to string pointer
		varType = types.NewPointer(types.I8)
	}
	
	varName := ctx.ID().GetText()
	
	// Create alloca instruction
	alloca := v.CurrentBlock.NewAlloca(varType)
	alloca.SetName(varName)
	v.SymbolTable[varName] = alloca
	
	// Process initializer
	exprValue := v.Visit(ctx.Expr())
	if val, ok := exprValue.(value.Value); ok {
		v.CurrentBlock.NewStore(val, alloca)
	} else {
		fmt.Printf("Warning: Invalid initializer for variable %s\n", varName)
	}
	
	return nil
}

// Visit assignment statement
func (v *CustomVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	fmt.Println("Visiting assignment:", ctx.ID().GetText())
	
	varName := ctx.ID().GetText()
	alloca, exists := v.SymbolTable[varName]
	
	if !exists {
		fmt.Printf("Error: Variable %s not declared\n", varName)
		return nil
	}
	
	// Process value
	exprValue := v.Visit(ctx.Expr())
	if val, ok := exprValue.(value.Value); ok {
		v.CurrentBlock.NewStore(val, alloca)
	} else {
		fmt.Printf("Warning: Invalid value for assignment to %s\n", varName)
	}
	
	return nil
}

// Visit function call statement
func (v *CustomVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	funcName := ctx.ID().GetText()
	fmt.Println("Visiting function call:", funcName)
	
	// Get the function
	fn, exists := v.FuncMap[funcName]
	if !exists {
		fmt.Printf("Warning: Undefined function %s, creating forward declaration\n", funcName)
		// Create a generic forward declaration (variadic with i8* return)
		fn = v.Module.NewFunc(funcName, types.I32, ir.NewParam("", types.NewPointer(types.I8)))
		fn.Sig.Variadic = true
		v.FuncMap[funcName] = fn
	}
	
	// Process arguments
	var args []value.Value
	for _, expr := range ctx.AllExpr() {
		argValue := v.Visit(expr)
		if val, ok := argValue.(value.Value); ok {
			args = append(args, val)
		} else {
			args = append(args, constant.NewInt(types.I32, 0)) // Default argument
		}
	}
	
	// Create call instruction
	call := v.CurrentBlock.NewCall(fn, args...)
	return call
}

// Visit return statement
func (v *CustomVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	fmt.Println("Visiting return statement")
	
	// Process return value
	retValue := v.Visit(ctx.Expr())
	if val, ok := retValue.(value.Value); ok {
		v.CurrentBlock.NewRet(val)
	} else {
		fmt.Println("Warning: Invalid return value, using default")
		v.CurrentBlock.NewRet(constant.NewInt(types.I32, 0))
	}
	
	return nil
}

// Visit multiplication/division expression
func (v *CustomVisitor) VisitMulDivExpr(ctx *MulDivExprContext) interface{} {
	fmt.Println("Visiting multiplication/division expression")
	
	// Process left and right operands
	leftValue := v.Visit(ctx.Expr(0))
	rightValue := v.Visit(ctx.Expr(1))
	
	left, leftOk := leftValue.(value.Value)
	right, rightOk := rightValue.(value.Value)
	
	if !leftOk || !rightOk {
		fmt.Println("Warning: Invalid operands in multiplication/division")
		return constant.NewInt(types.I32, 0)
	}
	
	// Get operation type
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	
	if op == "*" {
		return v.CurrentBlock.NewMul(left, right)
	} else { // "/"
		return v.CurrentBlock.NewSDiv(left, right)
	}
}

// Visit addition/subtraction expression
func (v *CustomVisitor) VisitAddSubExpr(ctx *AddSubExprContext) interface{} {
	fmt.Println("Visiting addition/subtraction expression")
	
	// Process left and right operands
	leftValue := v.Visit(ctx.Expr(0))
	rightValue := v.Visit(ctx.Expr(1))
	
	left, leftOk := leftValue.(value.Value)
	right, rightOk := rightValue.(value.Value)
	
	if !leftOk || !rightOk {
		fmt.Println("Warning: Invalid operands in addition/subtraction")
		return constant.NewInt(types.I32, 0)
	}
	
	// Get operation type
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	
	if op == "+" {
		return v.CurrentBlock.NewAdd(left, right)
	} else { // "-"
		return v.CurrentBlock.NewSub(left, right)
	}
}

// Visit function call expression
func (v *CustomVisitor) VisitFunctionCallExpr(ctx *FunctionCallExprContext) interface{} {
	fmt.Println("Visiting function call expression")
	
	// Process as regular function call
	return v.Visit(ctx.FunctionCall())
}

// Visit variable expression
func (v *CustomVisitor) VisitVariableExpr(ctx *VariableExprContext) interface{} {
	varName := ctx.ID().GetText()
	fmt.Println("Visiting variable expression:", varName)
	
	alloca, exists := v.SymbolTable[varName]
	if !exists {
		fmt.Printf("Error: Variable %s not declared\n", varName)
		return constant.NewInt(types.I32, 0)
	}
	
	// Load the variable value
	if ptr, ok := alloca.Type().(*types.PointerType); ok {
		return v.CurrentBlock.NewLoad(ptr.ElemType, alloca)
	}
	
	fmt.Printf("Error: Variable %s has invalid type\n", varName)
	return constant.NewInt(types.I32, 0)
}

// Visit literal expression
func (v *CustomVisitor) VisitLiteralExpr(ctx *LiteralExprContext) interface{} {
	fmt.Println("Visiting literal expression")
	
	literalCtx := ctx.Literal()
	
	if intLit := literalCtx.INT_LITERAL(); intLit != nil {
		val, _ := strconv.Atoi(intLit.GetText())
		return constant.NewInt(types.I32, int64(val))
	} else if floatLit := literalCtx.FLOAT_LITERAL(); floatLit != nil {
		val, _ := strconv.ParseFloat(floatLit.GetText(), 64)
		return constant.NewFloat(types.Float, val)
	} else if strLit := literalCtx.STRING_LITERAL(); strLit != nil {
		return v.createStringConstant(strLit.GetText())
	} else if boolLit := literalCtx.BOOL_LITERAL(); boolLit != nil {
		val := boolLit.GetText() == "true"
		if val {
			return constant.NewInt(types.I1, 1)
		}
		return constant.NewInt(types.I1, 0)
	}
	
	return constant.NewInt(types.I32, 0)
}

// Visit parenthesized expression
func (v *CustomVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	fmt.Println("Visiting parenthesized expression")
	
	// Just visit the inner expression
	return v.Visit(ctx.Expr())
}

func main() {
	// Add panic recovery
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace will be printed above")
		}
	}()
	
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input-file>")
		os.Exit(1)
	}
	
	fmt.Printf("Processing file: %s\n", os.Args[1])
	
	// Create input stream
	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}
	
	// Create lexer
	lexer := NewCustomLanguageLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	
	// Create parser
	parser := NewCustomLanguageParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	
	// Parse the input
	fmt.Println("Parsing input...")
	tree := parser.Program()
	
	// Create and run visitor
	fmt.Println("Running visitor...")
	visitor := NewCustomVisitor()
	result := visitor.Visit(tree)
	
	if result == nil {
		fmt.Println("Error: Visitor returned nil")
		os.Exit(1)
	}
	
	module, ok := result.(*ir.Module)
	if !ok {
		fmt.Printf("Error: Expected *ir.Module, got %T\n", result)
		os.Exit(1)
	}
	
	// Output LLVM IR to file
	outputFile := "program.ll"
	fmt.Printf("Writing LLVM IR to %s...\n", outputFile)
	
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer file.Close()
	
	file.WriteString(module.String())
	fmt.Println("LLVM IR generation completed successfully!")
}