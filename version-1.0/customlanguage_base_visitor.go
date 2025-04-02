// Code generated from CustomLanguage.g4 by ANTLR 4.13.1. DO NOT EDIT.

package main // CustomLanguage
import "github.com/antlr4-go/antlr/v4"

type BaseCustomLanguageVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCustomLanguageVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitParameters(ctx *ParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitVariableDecl(ctx *VariableDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitFunctionCallExpr(ctx *FunctionCallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitMulDivExpr(ctx *MulDivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitLiteralExpr(ctx *LiteralExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitVariableExpr(ctx *VariableExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitAddSubExpr(ctx *AddSubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomLanguageVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
