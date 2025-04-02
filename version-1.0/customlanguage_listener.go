// Code generated from CustomLanguage.g4 by ANTLR 4.13.1. DO NOT EDIT.

package main // CustomLanguage
import "github.com/antlr4-go/antlr/v4"

// CustomLanguageListener is a complete listener for a parse tree produced by CustomLanguageParser.
type CustomLanguageListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVariableDecl is called when entering the variableDecl production.
	EnterVariableDecl(c *VariableDeclContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterFunctionCallExpr is called when entering the FunctionCallExpr production.
	EnterFunctionCallExpr(c *FunctionCallExprContext)

	// EnterMulDivExpr is called when entering the MulDivExpr production.
	EnterMulDivExpr(c *MulDivExprContext)

	// EnterLiteralExpr is called when entering the LiteralExpr production.
	EnterLiteralExpr(c *LiteralExprContext)

	// EnterVariableExpr is called when entering the VariableExpr production.
	EnterVariableExpr(c *VariableExprContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterAddSubExpr is called when entering the AddSubExpr production.
	EnterAddSubExpr(c *AddSubExprContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVariableDecl is called when exiting the variableDecl production.
	ExitVariableDecl(c *VariableDeclContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitFunctionCallExpr is called when exiting the FunctionCallExpr production.
	ExitFunctionCallExpr(c *FunctionCallExprContext)

	// ExitMulDivExpr is called when exiting the MulDivExpr production.
	ExitMulDivExpr(c *MulDivExprContext)

	// ExitLiteralExpr is called when exiting the LiteralExpr production.
	ExitLiteralExpr(c *LiteralExprContext)

	// ExitVariableExpr is called when exiting the VariableExpr production.
	ExitVariableExpr(c *VariableExprContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitAddSubExpr is called when exiting the AddSubExpr production.
	ExitAddSubExpr(c *AddSubExprContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
