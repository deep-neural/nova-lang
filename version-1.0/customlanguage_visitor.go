// Code generated from CustomLanguage.g4 by ANTLR 4.13.1. DO NOT EDIT.

package main // CustomLanguage
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CustomLanguageParser.
type CustomLanguageVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CustomLanguageParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#variableDecl.
	VisitVariableDecl(ctx *VariableDeclContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#FunctionCallExpr.
	VisitFunctionCallExpr(ctx *FunctionCallExprContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#MulDivExpr.
	VisitMulDivExpr(ctx *MulDivExprContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#LiteralExpr.
	VisitLiteralExpr(ctx *LiteralExprContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#VariableExpr.
	VisitVariableExpr(ctx *VariableExprContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#AddSubExpr.
	VisitAddSubExpr(ctx *AddSubExprContext) interface{}

	// Visit a parse tree produced by CustomLanguageParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
