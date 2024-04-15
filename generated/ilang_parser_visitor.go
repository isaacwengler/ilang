// Code generated from ./ilang_parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ilang_parser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ilang_parser.
type ilang_parserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ilang_parser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by ilang_parser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ilang_parser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ilang_parser#intExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by ilang_parser#nullExpr.
	VisitNullExpr(ctx *NullExprContext) interface{}

	// Visit a parse tree produced by ilang_parser#arrayExpr.
	VisitArrayExpr(ctx *ArrayExprContext) interface{}

	// Visit a parse tree produced by ilang_parser#booleanExpr.
	VisitBooleanExpr(ctx *BooleanExprContext) interface{}

	// Visit a parse tree produced by ilang_parser#mapExpr.
	VisitMapExpr(ctx *MapExprContext) interface{}

	// Visit a parse tree produced by ilang_parser#arithmetic.
	VisitArithmetic(ctx *ArithmeticContext) interface{}

	// Visit a parse tree produced by ilang_parser#stringExpr.
	VisitStringExpr(ctx *StringExprContext) interface{}

	// Visit a parse tree produced by ilang_parser#floatExpr.
	VisitFloatExpr(ctx *FloatExprContext) interface{}

	// Visit a parse tree produced by ilang_parser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by ilang_parser#notExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by ilang_parser#functionDefExpr.
	VisitFunctionDefExpr(ctx *FunctionDefExprContext) interface{}

	// Visit a parse tree produced by ilang_parser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by ilang_parser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by ilang_parser#groupingExpr.
	VisitGroupingExpr(ctx *GroupingExprContext) interface{}

	// Visit a parse tree produced by ilang_parser#symbolExpr.
	VisitSymbolExpr(ctx *SymbolExprContext) interface{}

	// Visit a parse tree produced by ilang_parser#booleanAlgebra.
	VisitBooleanAlgebra(ctx *BooleanAlgebraContext) interface{}

	// Visit a parse tree produced by ilang_parser#computedProperty.
	VisitComputedProperty(ctx *ComputedPropertyContext) interface{}

	// Visit a parse tree produced by ilang_parser#functionDef.
	VisitFunctionDef(ctx *FunctionDefContext) interface{}

	// Visit a parse tree produced by ilang_parser#functionDefArgs.
	VisitFunctionDefArgs(ctx *FunctionDefArgsContext) interface{}

	// Visit a parse tree produced by ilang_parser#functionArgs.
	VisitFunctionArgs(ctx *FunctionArgsContext) interface{}

	// Visit a parse tree produced by ilang_parser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by ilang_parser#reassignment.
	VisitReassignment(ctx *ReassignmentContext) interface{}

	// Visit a parse tree produced by ilang_parser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by ilang_parser#whileLoop.
	VisitWhileLoop(ctx *WhileLoopContext) interface{}

	// Visit a parse tree produced by ilang_parser#foreachLoop.
	VisitForeachLoop(ctx *ForeachLoopContext) interface{}

	// Visit a parse tree produced by ilang_parser#forLoop.
	VisitForLoop(ctx *ForLoopContext) interface{}

	// Visit a parse tree produced by ilang_parser#return.
	VisitReturn(ctx *ReturnContext) interface{}

	// Visit a parse tree produced by ilang_parser#elseifStatement.
	VisitElseifStatement(ctx *ElseifStatementContext) interface{}

	// Visit a parse tree produced by ilang_parser#elseStatement.
	VisitElseStatement(ctx *ElseStatementContext) interface{}

	// Visit a parse tree produced by ilang_parser#scopeBody.
	VisitScopeBody(ctx *ScopeBodyContext) interface{}

	// Visit a parse tree produced by ilang_parser#conditionBody.
	VisitConditionBody(ctx *ConditionBodyContext) interface{}

	// Visit a parse tree produced by ilang_parser#not.
	VisitNot(ctx *NotContext) interface{}

	// Visit a parse tree produced by ilang_parser#symbol.
	VisitSymbol(ctx *SymbolContext) interface{}

	// Visit a parse tree produced by ilang_parser#symbolChild.
	VisitSymbolChild(ctx *SymbolChildContext) interface{}

	// Visit a parse tree produced by ilang_parser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by ilang_parser#intLiteral.
	VisitIntLiteral(ctx *IntLiteralContext) interface{}

	// Visit a parse tree produced by ilang_parser#floatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by ilang_parser#nullLiteral.
	VisitNullLiteral(ctx *NullLiteralContext) interface{}

	// Visit a parse tree produced by ilang_parser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by ilang_parser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by ilang_parser#mapLiteral.
	VisitMapLiteral(ctx *MapLiteralContext) interface{}

	// Visit a parse tree produced by ilang_parser#mapLiteralItem.
	VisitMapLiteralItem(ctx *MapLiteralItemContext) interface{}

	// Visit a parse tree produced by ilang_parser#mapKey.
	VisitMapKey(ctx *MapKeyContext) interface{}

	// Visit a parse tree produced by ilang_parser#grouping.
	VisitGrouping(ctx *GroupingContext) interface{}
}
