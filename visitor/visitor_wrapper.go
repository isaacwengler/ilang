package visitor

import (
	parser "ilang/generated"
	"ilang/logger"
	"ilang/library"

	"github.com/antlr4-go/antlr/v4"
)

/**
 * Some unfortunate boilerplate for the visitor. Every node needs two methods
 * `VisitNodeName` here and `VisitNodeNameWrapped` in visitor.go
 *
 * This is done because antlr generated the visitor to have to return `interface{}`,
 * making type safety difficult. In `visitor.go`, all methods return `types.WrappedValue`,
 * making it easier to reason with.
 *
 * TODO: move off antlr? It seems not best suited for go
 */

func (v *Visitor) VisitWrapped(t antlr.ParseTree) library.WrappedValue {
	val := v.Visit(t)
	if val == nil {
		logger.Warn("VisitWrapped got nil return value")
		return nil
	}
	return val.(library.WrappedValue)
}

func (v *Visitor) VisitChildrenWrapped(node antlr.RuleNode) library.WrappedValue {
	val := v.VisitChildren(node)
	if val == nil {
		logger.Warn("VisitChildrenWrapped got nil return value")
		return nil
	}
	return val.(library.WrappedValue)
}

func (v *Visitor) Visit(t antlr.ParseTree) interface{} {
	return t.Accept(v)
}

func (v *Visitor) VisitChildren(node antlr.RuleNode) interface{} {
	var val any
	var curr any
	for _, n := range node.GetChildren() {
		curr = v.Visit(n.(antlr.ParseTree))
		switch n.(type) {
		case *parser.ReturnContext:
			return curr
		default:
			// ignore terminal nodes
			if curr != nil {
				val = curr
			}
		}
	}

	return val
}

func (v *Visitor) VisitStart(ctx *parser.StartContext) interface{} {
	return v.VisitStartWrapped(ctx)
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	return v.VisitBlockWrapped(ctx)
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	return v.VisitStatementWrapped(ctx)
}

func (v *Visitor) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	return v.VisitIntExprWrapped(ctx)
}

func (v *Visitor) VisitNullExpr(ctx *parser.NullExprContext) interface{} {
	return v.VisitNullExprWrapped(ctx)
}

func (v *Visitor) VisitArrayExpr(ctx *parser.ArrayExprContext) interface{} {
	return v.VisitArrayExprWrapped(ctx)
}

func (v *Visitor) VisitBooleanExpr(ctx *parser.BooleanExprContext) interface{} {
	return v.VisitBooleanExprWrapped(ctx)
}

func (v *Visitor) VisitMapExpr(ctx *parser.MapExprContext) interface{} {
	return v.VisitMapExprWrapped(ctx)
}

func (v *Visitor) VisitArithmetic(ctx *parser.ArithmeticContext) interface{} {
	return v.VisitArithmeticWrapped(ctx)
}

func (v *Visitor) VisitStringExpr(ctx *parser.StringExprContext) interface{} {
	return v.VisitStringExprWrapped(ctx)
}

func (v *Visitor) VisitFloatExpr(ctx *parser.FloatExprContext) interface{} {
	return v.VisitFloatExprWrapped(ctx)
}

func (v *Visitor) VisitCondition(ctx *parser.ConditionContext) interface{} {
	return v.VisitConditionWrapped(ctx)
}

func (v *Visitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	return v.VisitNotExprWrapped(ctx)
}

func (v *Visitor) VisitFunctionDefExpr(ctx *parser.FunctionDefExprContext) interface{} {
	return v.VisitFunctionDefExprWrapped(ctx)
}

func (v *Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	return v.VisitFunctionCallWrapped(ctx)
}

func (v *Visitor) VisitSymbolExpr(ctx *parser.SymbolExprContext) interface{} {
	return v.VisitSymbolExprWrapped(ctx)
}

func (v *Visitor) VisitGroupingExpr(ctx *parser.GroupingExprContext) interface{} {
	return v.VisitGroupingExprWrapped(ctx)
}

func (v *Visitor) VisitBooleanAlgebra(ctx *parser.BooleanAlgebraContext) interface{} {
	return v.VisitBooleanAlgebraWrapped(ctx)
}

func (v *Visitor) VisitFunctionDef(ctx *parser.FunctionDefContext) interface{} {
	return v.VisitFunctionDefWrapped(ctx)
}

func (v *Visitor) VisitFunctionDefArgs(ctx *parser.FunctionDefArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctionArgs(ctx *parser.FunctionArgsContext) interface{} {
	return v.VisitFunctionArgsWrapped(ctx)
}

func (v *Visitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	return v.VisitAssignmentWrapped(ctx)
}

func (v *Visitor) VisitReassignment(ctx *parser.ReassignmentContext) interface{} {
	return v.VisitReassignmentWrapped(ctx)
}

func (v *Visitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	return v.VisitIfStatementWrapped(ctx)
}

func (v *Visitor) VisitWhileLoop(ctx *parser.WhileLoopContext) interface{} {
	return v.VisitWhileLoopWrapped(ctx)
}

func (v *Visitor) VisitForeachLoop(ctx *parser.ForeachLoopContext) interface{} {
	return v.VisitForeachLoopWrapped(ctx)
}

func (v *Visitor) VisitForLoop(ctx *parser.ForLoopContext) interface{} {
	return v.VisitForLoopWrapped(ctx)
}

func (v *Visitor) VisitReturn(ctx *parser.ReturnContext) interface{} {
	return v.VisitReturnWrapped(ctx)
}

func (v *Visitor) VisitElseifStatement(ctx *parser.ElseifStatementContext) interface{} {
	return v.VisitElseifStatementWrapped(ctx)
}

func (v *Visitor) VisitElseStatement(ctx *parser.ElseStatementContext) interface{} {
	return v.VisitElseStatementWrapped(ctx)
}

func (v *Visitor) VisitScopeBody(ctx *parser.ScopeBodyContext) interface{} {
	return v.VisitScopeBodyWrapped(ctx)
}

func (v *Visitor) VisitConditionBody(ctx *parser.ConditionBodyContext) interface{} {
	return v.VisitConditionBodyWrapped(ctx)
}

func (v *Visitor) VisitNot(ctx *parser.NotContext) interface{} {
	return v.VisitNotWrapped(ctx)
}

func (v *Visitor) VisitSymbol(ctx *parser.SymbolContext) interface{} {
	return v.VisitSymbolWrapped(ctx)
}

func (v *Visitor) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	return v.VisitStringLiteralWrapped(ctx)
}

func (v *Visitor) VisitIntLiteral(ctx *parser.IntLiteralContext) interface{} {
	return v.VisitIntLiteralWrapped(ctx)
}

func (v *Visitor) VisitFloatLiteral(ctx *parser.FloatLiteralContext) interface{} {
	return v.VisitFloatLiteralWrapped(ctx)
}

func (v *Visitor) VisitNullLiteral(ctx *parser.NullLiteralContext) interface{} {
	return v.VisitNullLiteralWrapped(ctx)
}

func (v *Visitor) VisitBooleanLiteral(ctx *parser.BooleanLiteralContext) interface{} {
	return v.VisitBooleanLiteralWrapped(ctx)
}

func (v *Visitor) VisitArrayLiteral(ctx *parser.ArrayLiteralContext) interface{} {
	return v.VisitArrayLiteralWrapped(ctx)
}

func (v *Visitor) VisitMapLiteral(ctx *parser.MapLiteralContext) interface{} {
	return v.VisitMapLiteralWrapped(ctx)
}

func (v *Visitor) VisitMapLiteralItem(ctx *parser.MapLiteralItemContext) interface{} {
	return v.VisitMapLiteralItemWrapped(ctx)
}

func (v *Visitor) VisitMapKey(ctx *parser.MapKeyContext) interface{} {
	return v.VisitMapKeyWrapped(ctx)
}

func (v *Visitor) VisitGrouping(ctx *parser.GroupingContext) interface{} {
	return v.VisitGroupingWrapped(ctx)
}

func (v *Visitor) VisitSymbolChild(ctx *parser.SymbolChildContext) interface{} {
	return v.VisitSymbolChildWrapped(ctx)
}

