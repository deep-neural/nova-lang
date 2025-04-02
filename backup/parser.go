package main // parser.go

import (
	"fmt"
	"strconv"
)

// NodeType represents the type of AST node
type NodeType int

const (
	NODE_PROGRAM NodeType = iota
	NODE_FUNCTION_DECL
	NODE_BLOCK
	NODE_VAR_DECL
	NODE_RETURN
	NODE_IF
	NODE_WHILE
	NODE_ASSIGN
	NODE_BINARY_OP
	NODE_CALL
	NODE_IDENT
	NODE_NUMBER
	NODE_FLOAT
	NODE_STRING
	NODE_BOOL
)

// Node represents an AST node
type Node interface {
	Type() NodeType
	String() string
}

// Program is the root node of the AST
type Program struct {
	Functions []*FunctionDecl
}

func (p *Program) Type() NodeType { return NODE_PROGRAM }
func (p *Program) String() string { return "Program" }

// FunctionDecl represents a function declaration
type FunctionDecl struct {
	Name       string
	Parameters []*VarDecl
	ReturnType string
	Body       *Block
}

func (f *FunctionDecl) Type() NodeType { return NODE_FUNCTION_DECL }
func (f *FunctionDecl) String() string {
	return fmt.Sprintf("FunctionDecl(%s)", f.Name)
}

// Block represents a block of statements
type Block struct {
	Statements []Statement
}

func (b *Block) Type() NodeType { return NODE_BLOCK }
func (b *Block) String() string { return "Block" }

// Statement represents a statement
type Statement interface {
	Node
	statementNode()
}

// Expression represents an expression
type Expression interface {
	Node
	expressionNode()
}

// VarDecl represents a variable declaration
type VarDecl struct {
	Name       string
	VarType    string
	Value      Expression
	IsTypeInferred bool
}

func (v *VarDecl) Type() NodeType { return NODE_VAR_DECL }
func (v *VarDecl) String() string {
	return fmt.Sprintf("VarDecl(%s: %s)", v.Name, v.VarType)
}
func (v *VarDecl) statementNode() {}

// ReturnStatement represents a return statement
type ReturnStatement struct {
	Value Expression
}

func (r *ReturnStatement) Type() NodeType { return NODE_RETURN }
func (r *ReturnStatement) String() string { return "Return" }
func (r *ReturnStatement) statementNode() {}

// IfStatement represents an if statement
type IfStatement struct {
	Condition Expression
	ThenBlock *Block
	ElseBlock *Block
}

func (i *IfStatement) Type() NodeType { return NODE_IF }
func (i *IfStatement) String() string { return "If" }
func (i *IfStatement) statementNode() {}

// WhileStatement represents a while statement
type WhileStatement struct {
	Condition Expression
	Body      *Block
}

func (w *WhileStatement) Type() NodeType { return NODE_WHILE }
func (w *WhileStatement) String() string { return "While" }
func (w *WhileStatement) statementNode() {}

// AssignStatement represents an assignment
type AssignStatement struct {
	Name  string
	Value Expression
}

func (a *AssignStatement) Type() NodeType { return NODE_ASSIGN }
func (a *AssignStatement) String() string { return fmt.Sprintf("Assign(%s)", a.Name) }
func (a *AssignStatement) statementNode() {}

// ExpressionStatement wraps an expression as a statement
type ExpressionStatement struct {
	Expression Expression
}

func (e *ExpressionStatement) Type() NodeType { return e.Expression.Type() }
func (e *ExpressionStatement) String() string { return e.Expression.String() }
func (e *ExpressionStatement) statementNode() {}

// BinaryOp represents a binary operation
type BinaryOp struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (b *BinaryOp) Type() NodeType { return NODE_BINARY_OP }
func (b *BinaryOp) String() string { return fmt.Sprintf("BinaryOp(%s)", b.Operator) }
func (b *BinaryOp) expressionNode() {}

// CallExpression represents a function call
type CallExpression struct {
	Function string
	Arguments []Expression
}

func (c *CallExpression) Type() NodeType { return NODE_CALL }
func (c *CallExpression) String() string { return fmt.Sprintf("Call(%s)", c.Function) }
func (c *CallExpression) expressionNode() {}

// Identifier represents a variable reference
type Identifier struct {
	Name string
}

func (i *Identifier) Type() NodeType { return NODE_IDENT }
func (i *Identifier) String() string { return fmt.Sprintf("Ident(%s)", i.Name) }
func (i *Identifier) expressionNode() {}

// NumberLiteral represents an integer literal
type NumberLiteral struct {
	Value int64
}

func (n *NumberLiteral) Type() NodeType { return NODE_NUMBER }
func (n *NumberLiteral) String() string { return fmt.Sprintf("Number(%d)", n.Value) }
func (n *NumberLiteral) expressionNode() {}

// FloatLiteral represents a float literal
type FloatLiteral struct {
	Value float64
}

func (f *FloatLiteral) Type() NodeType { return NODE_FLOAT }
func (f *FloatLiteral) String() string { return fmt.Sprintf("Float(%f)", f.Value) }
func (f *FloatLiteral) expressionNode() {}

// StringLiteral represents a string literal
type StringLiteral struct {
	Value string
}

func (s *StringLiteral) Type() NodeType { return NODE_STRING }
func (s *StringLiteral) String() string { return fmt.Sprintf("String(%s)", s.Value) }
func (s *StringLiteral) expressionNode() {}

// BoolLiteral represents a boolean literal
type BoolLiteral struct {
	Value bool
}

func (b *BoolLiteral) Type() NodeType { return NODE_BOOL }
func (b *BoolLiteral) String() string { return fmt.Sprintf("Bool(%t)", b.Value) }
func (b *BoolLiteral) expressionNode() {}

// Parser parses tokens into an AST
type Parser struct {
	tokenizer *Tokenizer
	currToken Token
	peekToken Token
	errors    []string
}

// NewParser creates a new parser
func NewParser(tokenizer *Tokenizer) *Parser {
	p := &Parser{tokenizer: tokenizer}
	
	// Read two tokens to set currToken and peekToken
	p.nextToken()
	p.nextToken()
	
	return p
}

// Errors returns any errors encountered during parsing
func (p *Parser) Errors() []string {
	return p.errors
}

// nextToken advances to the next token
func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.tokenizer.NextToken()
}

// Parse parses the program
func (p *Parser) Parse() *Program {
	program := &Program{
		Functions: []*FunctionDecl{},
	}
	
	for p.currToken.Type != TOKEN_EOF {
		// Handle top-level statements (only functions for now)
		if p.currToken.Type == TOKEN_FUNC {
			if fn := p.parseFunctionDecl(); fn != nil {
				program.Functions = append(program.Functions, fn)
			}
		} else {
			p.nextToken() // Skip unexpected tokens at the top level
		}
	}
	
	return program
}

// parseFunctionDecl parses a function declaration
func (p *Parser) parseFunctionDecl() *FunctionDecl {
	// Check 'func' keyword
	if p.currToken.Type != TOKEN_FUNC {
		p.addError("expected 'func', got %s", p.currToken.Literal)
		return nil
	}
	p.nextToken()
	
	// Parse function name
	if p.currToken.Type != TOKEN_IDENT {
		p.addError("expected function name, got %s", p.currToken.Literal)
		return nil
	}
	functionName := p.currToken.Literal
	p.nextToken()
	
	// Parse parameters
	if p.currToken.Type != TOKEN_LPAREN {
		p.addError("expected '(' after function name, got %s", p.currToken.Literal)
		return nil
	}
	p.nextToken() // Skip '('
	
	params := []*VarDecl{}
	
	// Parse parameter list
	if p.currToken.Type != TOKEN_RPAREN {
		// At least one parameter
		param := p.parseParameter()
		if param != nil {
			params = append(params, param)
		}
		
		// Parse additional parameters
		for p.currToken.Type == TOKEN_COMMA {
			p.nextToken() // Skip ','
			param = p.parseParameter()
			if param != nil {
				params = append(params, param)
			}
		}
	}
	
	// Check for closing parenthesis
	if p.currToken.Type != TOKEN_RPAREN {
		p.addError("expected ')' after parameters, got %s", p.currToken.Literal)
		return nil
	}
	p.nextToken() // Skip ')'
	
	// Parse return type
	var returnType string = "void" // Default return type
	
	if p.currToken.Type == TOKEN_ARROW {
		p.nextToken() // Skip '->'
		
		switch p.currToken.Type {
		case TOKEN_TYPE_INT:
			returnType = "int"
		case TOKEN_TYPE_FLOAT:
			returnType = "float"
		case TOKEN_TYPE_STRING:
			returnType = "string"
		case TOKEN_TYPE_BOOL:
			returnType = "bool"
		case TOKEN_TYPE_VOID:
			returnType = "void"
		default:
			p.addError("expected return type, got %s", p.currToken.Literal)
			return nil
		}
		p.nextToken() // Skip return type
	}
	
	// Parse function body
	if p.currToken.Type != TOKEN_LBRACE {
		p.addError("expected '{' to start function body, got %s", p.currToken.Literal)
		return nil
	}
	
	body := p.parseBlock()
	
	return &FunctionDecl{
		Name:       functionName,
		Parameters: params,
		ReturnType: returnType,
		Body:       body,
	}
}

// parseParameter parses a function parameter
func (p *Parser) parseParameter() *VarDecl {
	// Get parameter name
	if p.currToken.Type != TOKEN_IDENT {
		p.addError("expected parameter name, got %s", p.currToken.Literal)
		return nil
	}
	paramName := p.currToken.Literal
	p.nextToken()
	
	// Check for type separator
	if p.currToken.Type != TOKEN_COLON {
		p.addError("expected ':' after parameter name, got %s", p.currToken.Literal)
		return nil
	}
	p.nextToken() // Skip ':'
	
	// Get parameter type
	var paramType string
	switch p.currToken.Type {
	case TOKEN_TYPE_INT:
		paramType = "int"
	case TOKEN_TYPE_FLOAT:
		paramType = "float"
	case TOKEN_TYPE_STRING:
		paramType = "string"
	case TOKEN_TYPE_BOOL:
		paramType = "bool"
	default:
		p.addError("expected type, got %s", p.currToken.Literal)
		return nil
	}
	p.nextToken() // Skip type
	
	return &VarDecl{
		Name:    paramName,
		VarType: paramType,
	}
}

// parseBlock parses a block of statements
func (p *Parser) parseBlock() *Block {
	block := &Block{
		Statements: []Statement{},
	}
	
	if p.currToken.Type != TOKEN_LBRACE {
		p.addError("expected '{', got %s", p.currToken.Literal)
		return block
	}
	p.nextToken() // Skip '{'
	
	for p.currToken.Type != TOKEN_RBRACE && p.currToken.Type != TOKEN_EOF {
		// Skip comments
		if p.currToken.Type == TOKEN_COMMENT {
			p.nextToken()
			continue
		}
		
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
	}
	
	if p.currToken.Type != TOKEN_RBRACE {
		p.addError("expected '}', got %s", p.currToken.Literal)
		return block
	}
	p.nextToken() // Skip '}'
	
	return block
}

// parseStatement parses a statement
func (p *Parser) parseStatement() Statement {
	switch p.currToken.Type {
	case TOKEN_TYPE_INT, TOKEN_TYPE_FLOAT, TOKEN_TYPE_STRING, TOKEN_TYPE_BOOL:
		return p.parseVarDecl()
	case TOKEN_VAR:
		return p.parseInferredVarDecl()
	case TOKEN_RETURN:
		return p.parseReturnStatement()
	case TOKEN_IF:
		return p.parseIfStatement()
	case TOKEN_WHILE:
		return p.parseWhileStatement()
	case TOKEN_IDENT:
		// This could be an assignment or a function call
		if p.peekToken.Type == TOKEN_EQUALS {
			return p.parseAssignStatement()
		} else {
			expr := p.parseExpression()
			if p.currToken.Type == TOKEN_SEMICOLON {
				p.nextToken() // Skip ';'
			}
			return &ExpressionStatement{Expression: expr}
		}
	default:
		// Try to parse as expression (function call, etc.)
		expr := p.parseExpression()
		if p.currToken.Type == TOKEN_SEMICOLON {
			p.nextToken() // Skip ';'
		}
		return &ExpressionStatement{Expression: expr}
	}
}

// parseVarDecl parses a variable declaration with explicit type
func (p *Parser) parseVarDecl() *VarDecl {
	// Get type
	var varType string
	switch p.currToken.Type {
	case TOKEN_TYPE_INT:
		varType = "int"
	case TOKEN_TYPE_FLOAT:
		varType = "float"
	case TOKEN_TYPE_STRING:
		varType = "string"
	case TOKEN_TYPE_BOOL:
		varType = "bool"
	default:
		p.addError("expected type, got %s", p.currToken.Literal)
		return nil
	}
	p.nextToken() // Skip type
	
	// Get variable name
	if p.currToken.Type != TOKEN_IDENT {
		p.addError("expected variable name, got %s", p.currToken.Literal)
		return nil
	}
	varName := p.currToken.Literal
	p.nextToken() // Skip variable name
	
	// Check for initialization
	var value Expression
	if p.currToken.Type == TOKEN_EQUALS {
		p.nextToken() // Skip '='
		value = p.parseExpression()
	}
	
	// Check for semicolon
	if p.currToken.Type == TOKEN_SEMICOLON {
		p.nextToken() // Skip ';'
	}
	
	return &VarDecl{
		Name:    varName,
		VarType: varType,
		Value:   value,
	}
}

// parseInferredVarDecl parses a variable declaration with type inference
func (p *Parser) parseInferredVarDecl() *VarDecl {
	// Skip 'var' keyword
	p.nextToken()
	
	// Get variable name
	if p.currToken.Type != TOKEN_IDENT {
		p.addError("expected variable name, got %s", p.currToken.Literal)
		return nil
	}
	varName := p.currToken.Literal
	p.nextToken() // Skip variable name
	
	// Check for initialization
	if p.currToken.Type != TOKEN_EQUALS {
		p.addError("expected '=', got %s", p.currToken.Literal)
		return nil
	}
	p.nextToken() // Skip '='
	
	value := p.parseExpression()
	
	// Check for semicolon
	if p.currToken.Type == TOKEN_SEMICOLON {
		p.nextToken() // Skip ';'
	}
	
	return &VarDecl{
		Name:           varName,
		Value:          value,
		IsTypeInferred: true,
	}
}

// parseReturnStatement parses a return statement
func (p *Parser) parseReturnStatement() *ReturnStatement {
	// Skip 'return' keyword
	p.nextToken()
	
	var value Expression
	if p.currToken.Type != TOKEN_SEMICOLON {
		value = p.parseExpression()
	}
	
	// Check for semicolon
	if p.currToken.Type == TOKEN_SEMICOLON {
		p.nextToken() // Skip ';'
	}
	
	return &ReturnStatement{Value: value}
}

// parseIfStatement parses an if statement
func (p *Parser) parseIfStatement() *IfStatement {
	// Skip 'if' keyword
	p.nextToken()
	
	// Parse condition
	if p.currToken.Type != TOKEN_LPAREN {
		p.addError("expected '(' after 'if', got %s", p.currToken.Literal)
		return nil
	}
	p.nextToken() // Skip '('
	
	condition := p.parseExpression()
	
	if p.currToken.Type != TOKEN_RPAREN {
		p.addError("expected ')' after condition, got %s", p.currToken.Literal)
		return nil
	}
	p.nextToken() // Skip ')'
	
	// Parse then block
	thenBlock := p.parseBlock()
	
	// Parse optional else block
	var elseBlock *Block
	if p.currToken.Type == TOKEN_ELSE {
		p.nextToken() // Skip 'else'
		elseBlock = p.parseBlock()
	}
	
	return &IfStatement{
		Condition: condition,
		ThenBlock: thenBlock,
		ElseBlock: elseBlock,
	}
}

// parseWhileStatement parses a while statement
func (p *Parser) parseWhileStatement() *WhileStatement {
	// Skip 'while' keyword
	p.nextToken()
	
	// Parse condition
	if p.currToken.Type != TOKEN_LPAREN {
		p.addError("expected '(' after 'while', got %s", p.currToken.Literal)
		return nil
	}
	p.nextToken() // Skip '('
	
	condition := p.parseExpression()
	
	if p.currToken.Type != TOKEN_RPAREN {
		p.addError("expected ')' after condition, got %s", p.currToken.Literal)
		return nil
	}
	p.nextToken() // Skip ')'
	
	// Parse body
	body := p.parseBlock()
	
	return &WhileStatement{
		Condition: condition,
		Body:      body,
	}
}

// parseAssignStatement parses an assignment statement
func (p *Parser) parseAssignStatement() *AssignStatement {
	// Get variable name
	varName := p.currToken.Literal
	p.nextToken() // Skip variable name
	
	// Skip '='
	if p.currToken.Type != TOKEN_EQUALS {
		p.addError("expected '=', got %s", p.currToken.Literal)
		return nil
	}
	p.nextToken() // Skip '='
	
	// Parse value
	value := p.parseExpression()
	
	// Check for semicolon
	if p.currToken.Type == TOKEN_SEMICOLON {
		p.nextToken() // Skip ';'
	}
	
	return &AssignStatement{
		Name:  varName,
		Value: value,
	}
}

// parseExpression parses an expression
func (p *Parser) parseExpression() Expression {
	left := p.parsePrimaryExpression()
	
	// Check for binary operators
	if p.isBinaryOp(p.currToken) {
		return p.parseBinaryOp(left, 0)
	}
	
	return left
}

// parseBinaryOp parses a binary operation with precedence
func (p *Parser) parseBinaryOp(left Expression, minPrecedence int) Expression {
	for p.isBinaryOp(p.currToken) && p.getPrecedence(p.currToken) >= minPrecedence {
		op := p.currToken.Literal
		precedence := p.getPrecedence(p.currToken)
		p.nextToken() // Skip operator
		
		right := p.parsePrimaryExpression()
		
		// Check for higher precedence operators on the right
		for p.isBinaryOp(p.currToken) && p.getPrecedence(p.currToken) > precedence {
			right = p.parseBinaryOp(right, p.getPrecedence(p.currToken))
		}
		
		// Create binary operation
		left = &BinaryOp{
			Left:     left,
			Operator: op,
			Right:    right,
		}
	}
	
	return left
}

// parsePrimaryExpression parses a primary expression (literal, identifier, or function call)
func (p *Parser) parsePrimaryExpression() Expression {
	switch p.currToken.Type {
	case TOKEN_IDENT:
		// This could be a variable reference or function call
		id := p.currToken.Literal
		p.nextToken()
		
		if p.currToken.Type == TOKEN_LPAREN {
			// Function call
			p.nextToken() // Skip '('
			
			args := []Expression{}
			
			// Parse arguments
			if p.currToken.Type != TOKEN_RPAREN {
				args = append(args, p.parseExpression())
				
				for p.currToken.Type == TOKEN_COMMA {
					p.nextToken() // Skip ','
					args = append(args, p.parseExpression())
				}
			}
			
			if p.currToken.Type != TOKEN_RPAREN {
				p.addError("expected ')' after arguments, got %s", p.currToken.Literal)
				return nil
			}
			p.nextToken() // Skip ')'
			
			return &CallExpression{
				Function:  id,
				Arguments: args,
			}
		}
		
		// Variable reference
		return &Identifier{Name: id}
		
	case TOKEN_NUMBER:
		// Integer literal
		value, err := strconv.ParseInt(p.currToken.Literal, 10, 64)
		if err != nil {
			p.addError("invalid integer: %s", p.currToken.Literal)
			return nil
		}
		p.nextToken()
		return &NumberLiteral{Value: value}
		
	case TOKEN_FLOAT:
		// Float literal
		value, err := strconv.ParseFloat(p.currToken.Literal, 64)
		if err != nil {
			p.addError("invalid float: %s", p.currToken.Literal)
			return nil
		}
		p.nextToken()
		return &FloatLiteral{Value: value}
		
	case TOKEN_STRING:
		// String literal
		value := p.currToken.Literal
		p.nextToken()
		return &StringLiteral{Value: value}
		
	case TOKEN_TRUE:
		// Boolean true
		p.nextToken()
		return &BoolLiteral{Value: true}
		
	case TOKEN_FALSE:
		// Boolean false
		p.nextToken()
		return &BoolLiteral{Value: false}
		
	case TOKEN_LPAREN:
		// Parenthesized expression
		p.nextToken() // Skip '('
		expr := p.parseExpression()
		
		if p.currToken.Type != TOKEN_RPAREN {
			p.addError("expected ')', got %s", p.currToken.Literal)
			return nil
		}
		p.nextToken() // Skip ')'
		return expr
		
	default:
		p.addError("unexpected token: %s", p.currToken.Literal)
		p.nextToken() // Skip unexpected token
		return nil
	}
}

// isBinaryOp checks if the token is a binary operator
func (p *Parser) isBinaryOp(token Token) bool {
	switch token.Type {
	case TOKEN_PLUS, TOKEN_MINUS, TOKEN_STAR, TOKEN_SLASH,
		TOKEN_EQ_EQUALS, TOKEN_NOT_EQUALS, TOKEN_LESS, TOKEN_LESS_EQUALS,
		TOKEN_GREATER, TOKEN_GREATER_EQUALS:
		return true
	default:
		return false
	}
}

// getPrecedence returns the precedence of an operator
func (p *Parser) getPrecedence(token Token) int {
	switch token.Type {
	case TOKEN_STAR, TOKEN_SLASH:
		return 5
	case TOKEN_PLUS, TOKEN_MINUS:
		return 4
	case TOKEN_EQ_EQUALS, TOKEN_NOT_EQUALS, TOKEN_LESS, TOKEN_LESS_EQUALS,
		TOKEN_GREATER, TOKEN_GREATER_EQUALS:
		return 3
	default:
		return 0
	}
}

// addError adds an error message
func (p *Parser) addError(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	p.errors = append(p.errors, fmt.Sprintf("line %d: %s", p.currToken.Line, msg))
}