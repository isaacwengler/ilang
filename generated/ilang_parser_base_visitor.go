// Code generated from ./ilang_parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ilang_parser

import "github.com/antlr4-go/antlr/v4"

type Baseilang_parserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *Baseilang_parserVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitIntExpr(ctx *IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitNullExpr(ctx *NullExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitArrayExpr(ctx *ArrayExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitBooleanExpr(ctx *BooleanExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitMapExpr(ctx *MapExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitArithmetic(ctx *ArithmeticContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitStringExpr(ctx *StringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitFloatExpr(ctx *FloatExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitFunctionDefExpr(ctx *FunctionDefExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitSymbolExpr(ctx *SymbolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitGroupingExpr(ctx *GroupingExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitBooleanAlgebra(ctx *BooleanAlgebraContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitFunctionDef(ctx *FunctionDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitFunctionArgs(ctx *FunctionArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitWhileLoop(ctx *WhileLoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitForeachLoop(ctx *ForeachLoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitForLoop(ctx *ForLoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitReturn(ctx *ReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitElseifStatement(ctx *ElseifStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitElseStatement(ctx *ElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitScopeBody(ctx *ScopeBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitConditionBody(ctx *ConditionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitNot(ctx *NotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitSymbol(ctx *SymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitMapLiteral(ctx *MapLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitMapLiteralItem(ctx *MapLiteralItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitMapKey(ctx *MapKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseilang_parserVisitor) VisitGrouping(ctx *GroupingContext) interface{} {
	return v.VisitChildren(ctx)
}
