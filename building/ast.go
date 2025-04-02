package main // ast.go

import (
	"fmt"
	"strings"
)


// CompoundAssignmentExpression represents a compound assignment expression like +=, -=, etc.
type CompoundAssignmentExpression struct {
	Token    Token // The token (+=, -=, etc.)
	Left     Expression
	Operator string
	Value    Expression
}

func (cae *CompoundAssignmentExpression) expressionNode() {}
func (cae *CompoundAssignmentExpression) TokenLiteral() string { return cae.Token.Literal }
func (cae *CompoundAssignmentExpression) String() string {
	var out strings.Builder
	
	out.WriteString(cae.Left.String())
	out.WriteString(" " + cae.Operator + " ")
	out.WriteString(cae.Value.String())
	
	return out.String()
}

// Node is the base interface for all AST nodes
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is a node that represents a statement
type Statement interface {
	Node
	statementNode()
}

// Expression is a node that represents an expression
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of every AST
type Program struct {
	Imports     []*ImportStatement
	Interfaces  []*InterfaceDefinition
	Functions   []*FunctionDefinition
}

func (p *Program) TokenLiteral() string {
	if len(p.Functions) > 0 {
		return p.Functions[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out strings.Builder
	
	for _, imp := range p.Imports {
		out.WriteString(imp.String())
		out.WriteString("\n")
	}

	for _, intf := range p.Interfaces {
		out.WriteString(intf.String())
		out.WriteString("\n")
	}

	for _, fn := range p.Functions {
		out.WriteString(fn.String())
		out.WriteString("\n")
	}
	
	return out.String()
}

// ImportStatement represents an import statement
type ImportStatement struct {
	Token Token // import token
	Path  string
}

func (is *ImportStatement) statementNode() {}
func (is *ImportStatement) TokenLiteral() string { return is.Token.Literal }
func (is *ImportStatement) String() string {
	return fmt.Sprintf("import \"%s\";", is.Path)
}

// InterfaceDefinition represents an interface definition
type InterfaceDefinition struct {
	Token Token // interface token
	Name  string
	Fields []*FieldDefinition
}

func (id *InterfaceDefinition) statementNode() {}
func (id *InterfaceDefinition) TokenLiteral() string { return id.Token.Literal }
func (id *InterfaceDefinition) String() string {
	var out strings.Builder
	out.WriteString("interface ")
	out.WriteString(id.Name)
	out.WriteString(" {\n")
	
	for _, field := range id.Fields {
		out.WriteString("    ")
		out.WriteString(field.String())
		out.WriteString("\n")
	}
	
	out.WriteString("};")
	return out.String()
}

// FieldDefinition represents a field in an interface or struct
type FieldDefinition struct {
	Token Token
	Type  TypeSpecifier
	Name  string
}

func (fd *FieldDefinition) statementNode() {}
func (fd *FieldDefinition) TokenLiteral() string { return fd.Token.Literal }
func (fd *FieldDefinition) String() string {
	return fmt.Sprintf("%s %s;", fd.Type.String(), fd.Name)
}

// TypeSpecifier represents a type specification
type TypeSpecifier struct {
	Token     Token
	TypeName  string
	IsPointer bool
}

func (ts *TypeSpecifier) String() string {
	if ts.IsPointer {
		return ts.TypeName + "*"
	}
	return ts.TypeName
}

// FunctionDefinition represents a function definition
type FunctionDefinition struct {
	Token      Token // func token
	Name       string
	Parameters []*ParameterDefinition
	ReturnType TypeSpecifier
	Body       *BlockStatement
}

func (fd *FunctionDefinition) statementNode() {}
func (fd *FunctionDefinition) TokenLiteral() string { return fd.Token.Literal }
func (fd *FunctionDefinition) String() string {
	var out strings.Builder
	
	out.WriteString("func ")
	out.WriteString(fd.Name)
	out.WriteString("(")
	
	params := []string{}
	for _, p := range fd.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(strings.Join(params, ", "))
	
	out.WriteString(") -> ")
	out.WriteString(fd.ReturnType.String())
	out.WriteString(" ")
	out.WriteString(fd.Body.String())
	
	return out.String()
}

// ParameterDefinition represents a function parameter
type ParameterDefinition struct {
	Token Token
	Name  string
	Type  TypeSpecifier
}

func (pd *ParameterDefinition) String() string {
	return fmt.Sprintf("%s: %s", pd.Name, pd.Type.String())
}

// BlockStatement represents a block of statements
type BlockStatement struct {
	Token      Token // {
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out strings.Builder
	
	out.WriteString("{\n")
	for _, s := range bs.Statements {
		out.WriteString("  ")
		out.WriteString(s.String())
		out.WriteString("\n")
	}
	out.WriteString("}")
	
	return out.String()
}

// VariableDeclarationStatement represents a variable declaration
type VariableDeclarationStatement struct {
	Token Token
	Type  TypeSpecifier
	Name  string
	Value Expression
}

func (vs *VariableDeclarationStatement) statementNode() {}
func (vs *VariableDeclarationStatement) TokenLiteral() string { return vs.Token.Literal }
func (vs *VariableDeclarationStatement) String() string {
	var out strings.Builder
	
	out.WriteString(vs.Type.String())
	out.WriteString(" ")
	out.WriteString(vs.Name)
	
	if vs.Value != nil {
		out.WriteString(" = ")
		out.WriteString(vs.Value.String())
	}
	
	out.WriteString(";")
	
	return out.String()
}

// ReturnStatement represents a return statement
type ReturnStatement struct {
	Token       Token // return token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out strings.Builder
	
	out.WriteString(rs.TokenLiteral() + " ")
	
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	
	out.WriteString(";")
	
	return out.String()
}

// ExpressionStatement represents a statement that consists of an expression
type ExpressionStatement struct {
	Token      Token // first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String() + ";"
	}
	return ";"
}

// IfStatement represents an if statement
type IfStatement struct {
	Token       Token // if token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (is *IfStatement) statementNode() {}
func (is *IfStatement) TokenLiteral() string { return is.Token.Literal }
func (is *IfStatement) String() string {
	var out strings.Builder
	
	out.WriteString("if (")
	out.WriteString(is.Condition.String())
	out.WriteString(") ")
	out.WriteString(is.Consequence.String())
	
	if is.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(is.Alternative.String())
	}
	
	return out.String()
}

// SwitchStatement represents a switch statement
type SwitchStatement struct {
	Token     Token // switch token
	Value     Expression
	Cases     []*CaseStatement
	Default   *BlockStatement
}

func (ss *SwitchStatement) statementNode() {}
func (ss *SwitchStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *SwitchStatement) String() string {
	var out strings.Builder
	
	out.WriteString("switch (")
	out.WriteString(ss.Value.String())
	out.WriteString(") {\n")
	
	for _, c := range ss.Cases {
		out.WriteString(c.String())
		out.WriteString("\n")
	}
	
	if ss.Default != nil {
		out.WriteString("default:")
		out.WriteString(ss.Default.String())
		out.WriteString("\n")
	}
	
	out.WriteString("}")
	
	return out.String()
}

// CaseStatement represents a case in a switch statement
type CaseStatement struct {
	Token     Token // case token
	Value     Expression
	Block     *BlockStatement
}

func (cs *CaseStatement) statementNode() {}
func (cs *CaseStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *CaseStatement) String() string {
	var out strings.Builder
	
	out.WriteString("case ")
	out.WriteString(cs.Value.String())
	out.WriteString(":")
	out.WriteString(cs.Block.String())
	
	return out.String()
}

// WhileStatement represents a while loop
type WhileStatement struct {
	Token     Token // while token
	Condition Expression
	Body      *BlockStatement
}

func (ws *WhileStatement) statementNode() {}
func (ws *WhileStatement) TokenLiteral() string { return ws.Token.Literal }
func (ws *WhileStatement) String() string {
	var out strings.Builder
	
	out.WriteString("while (")
	out.WriteString(ws.Condition.String())
	out.WriteString(") ")
	out.WriteString(ws.Body.String())
	
	return out.String()
}

// ForStatement represents a for loop
type ForStatement struct {
	Token      Token // for token
	Init       Statement
	Condition  Expression
	Update     Expression
	Body       *BlockStatement
}

func (fs *ForStatement) statementNode() {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForStatement) String() string {
	var out strings.Builder
	
	out.WriteString("for (")
	
	if fs.Init != nil {
		out.WriteString(fs.Init.String())
	} else {
		out.WriteString(";")
	}
	
	if fs.Condition != nil {
		out.WriteString(" ")
		out.WriteString(fs.Condition.String())
	}
	out.WriteString("; ")
	
	if fs.Update != nil {
		out.WriteString(fs.Update.String())
	}
	
	out.WriteString(") ")
	out.WriteString(fs.Body.String())
	
	return out.String()
}

// BreakStatement represents a break statement
type BreakStatement struct {
	Token Token // break token
}

func (bs *BreakStatement) statementNode() {}
func (bs *BreakStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BreakStatement) String() string { return "break;" }

// ContinueStatement represents a continue statement
type ContinueStatement struct {
	Token Token // continue token
}

func (cs *ContinueStatement) statementNode() {}
func (cs *ContinueStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *ContinueStatement) String() string { return "continue;" }

// IntegerLiteral represents an integer literal
type IntegerLiteral struct {
	Token Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string { return il.Token.Literal }

// FloatLiteral represents a floating-point literal
type FloatLiteral struct {
	Token Token
	Value float64
}

func (fl *FloatLiteral) expressionNode() {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FloatLiteral) String() string { return fl.Token.Literal }

// BooleanLiteral represents a boolean literal
type BooleanLiteral struct {
	Token Token
	Value bool
}

func (bl *BooleanLiteral) expressionNode() {}
func (bl *BooleanLiteral) TokenLiteral() string { return bl.Token.Literal }
func (bl *BooleanLiteral) String() string { return bl.Token.Literal }

// CharLiteral represents a character literal
type CharLiteral struct {
	Token Token
	Value rune
}

func (cl *CharLiteral) expressionNode() {}
func (cl *CharLiteral) TokenLiteral() string { return cl.Token.Literal }
func (cl *CharLiteral) String() string { return cl.Token.Literal }

// StringLiteral represents a string literal
type StringLiteral struct {
	Token Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string { return "\"" + sl.Value + "\"" }

// Identifier represents an identifier
type Identifier struct {
	Token Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string { return i.Value }

// PrefixExpression represents a prefix operator expression
type PrefixExpression struct {
	Token    Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out strings.Builder
	
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	
	return out.String()
}

// InfixExpression represents an infix operator expression
type InfixExpression struct {
	Token    Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out strings.Builder
	
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	
	return out.String()
}

// PostfixExpression represents a postfix operator expression
type PostfixExpression struct {
	Token    Token // The postfix token, e.g. ++
	Left     Expression
	Operator string
}

func (pe *PostfixExpression) expressionNode() {}
func (pe *PostfixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PostfixExpression) String() string {
	var out strings.Builder
	
	out.WriteString("(")
	out.WriteString(pe.Left.String())
	out.WriteString(pe.Operator)
	out.WriteString(")")
	
	return out.String()
}

// CallExpression represents a function call
type CallExpression struct {
	Token     Token // The '(' token
	Function  Expression // Identifier or another expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out strings.Builder
	args := []string{}
	
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	
	return out.String()
}

// IndexExpression represents an array indexing expression
type IndexExpression struct {
	Token Token // The '[' token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out strings.Builder
	
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	
	return out.String()
}

// DotExpression represents a member access expression
type DotExpression struct {
	Token  Token // The '.' token
	Object Expression
	Member *Identifier
}

func (de *DotExpression) expressionNode() {}
func (de *DotExpression) TokenLiteral() string { return de.Token.Literal }
func (de *DotExpression) String() string {
	var out strings.Builder
	
	out.WriteString(de.Object.String())
	out.WriteString(".")
	out.WriteString(de.Member.String())
	
	return out.String()
}

// StructLiteralExpression represents a struct literal
type StructLiteralExpression struct {
	Token  Token // The '{' token
	Fields map[string]Expression
}

func (sl *StructLiteralExpression) expressionNode() {}
func (sl *StructLiteralExpression) TokenLiteral() string { return sl.Token.Literal }
func (sl *StructLiteralExpression) String() string {
	var out strings.Builder
	pairs := []string{}
	
	for key, value := range sl.Fields {
		pairs = append(pairs, key + ": " + value.String())
	}
	
	out.WriteString("{ ")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString(" }")
	
	return out.String()
}

// AssignmentExpression represents an assignment expression
type AssignmentExpression struct {
	Token Token // The = token
	Left  Expression
	Value Expression
}

func (ae *AssignmentExpression) expressionNode() {}
func (ae *AssignmentExpression) TokenLiteral() string { return ae.Token.Literal }
func (ae *AssignmentExpression) String() string {
	var out strings.Builder
	
	out.WriteString(ae.Left.String())
	out.WriteString(" = ")
	out.WriteString(ae.Value.String())
	
	return out.String()
}