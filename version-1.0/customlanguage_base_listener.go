// Code generated from CustomLanguage.g4 by ANTLR 4.13.1. DO NOT EDIT.

package main // CustomLanguage
import "github.com/antlr4-go/antlr/v4"

// BaseCustomLanguageListener is a complete listener for a parse tree produced by CustomLanguageParser.
type BaseCustomLanguageListener struct{}

var _ CustomLanguageListener = &BaseCustomLanguageListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCustomLanguageListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCustomLanguageListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCustomLanguageListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCustomLanguageListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseCustomLanguageListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseCustomLanguageListener) ExitProgram(ctx *ProgramContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseCustomLanguageListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseCustomLanguageListener) ExitFunction(ctx *FunctionContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseCustomLanguageListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseCustomLanguageListener) ExitParameters(ctx *ParametersContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseCustomLanguageListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseCustomLanguageListener) ExitParameter(ctx *ParameterContext) {}

// EnterType is called when production type is entered.
func (s *BaseCustomLanguageListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseCustomLanguageListener) ExitType(ctx *TypeContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseCustomLanguageListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseCustomLanguageListener) ExitBlock(ctx *BlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCustomLanguageListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCustomLanguageListener) ExitStatement(ctx *StatementContext) {}

// EnterVariableDecl is called when production variableDecl is entered.
func (s *BaseCustomLanguageListener) EnterVariableDecl(ctx *VariableDeclContext) {}

// ExitVariableDecl is called when production variableDecl is exited.
func (s *BaseCustomLanguageListener) ExitVariableDecl(ctx *VariableDeclContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseCustomLanguageListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseCustomLanguageListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseCustomLanguageListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseCustomLanguageListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseCustomLanguageListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseCustomLanguageListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterFunctionCallExpr is called when production FunctionCallExpr is entered.
func (s *BaseCustomLanguageListener) EnterFunctionCallExpr(ctx *FunctionCallExprContext) {}

// ExitFunctionCallExpr is called when production FunctionCallExpr is exited.
func (s *BaseCustomLanguageListener) ExitFunctionCallExpr(ctx *FunctionCallExprContext) {}

// EnterMulDivExpr is called when production MulDivExpr is entered.
func (s *BaseCustomLanguageListener) EnterMulDivExpr(ctx *MulDivExprContext) {}

// ExitMulDivExpr is called when production MulDivExpr is exited.
func (s *BaseCustomLanguageListener) ExitMulDivExpr(ctx *MulDivExprContext) {}

// EnterLiteralExpr is called when production LiteralExpr is entered.
func (s *BaseCustomLanguageListener) EnterLiteralExpr(ctx *LiteralExprContext) {}

// ExitLiteralExpr is called when production LiteralExpr is exited.
func (s *BaseCustomLanguageListener) ExitLiteralExpr(ctx *LiteralExprContext) {}

// EnterVariableExpr is called when production VariableExpr is entered.
func (s *BaseCustomLanguageListener) EnterVariableExpr(ctx *VariableExprContext) {}

// ExitVariableExpr is called when production VariableExpr is exited.
func (s *BaseCustomLanguageListener) ExitVariableExpr(ctx *VariableExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseCustomLanguageListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseCustomLanguageListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterAddSubExpr is called when production AddSubExpr is entered.
func (s *BaseCustomLanguageListener) EnterAddSubExpr(ctx *AddSubExprContext) {}

// ExitAddSubExpr is called when production AddSubExpr is exited.
func (s *BaseCustomLanguageListener) ExitAddSubExpr(ctx *AddSubExprContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseCustomLanguageListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseCustomLanguageListener) ExitLiteral(ctx *LiteralContext) {}
