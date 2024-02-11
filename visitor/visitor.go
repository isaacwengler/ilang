package visitor

import (
	"errors"
	parser "ilang/generated"
	"ilang/logger"
	"ilang/types"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

func NewVisitor() *Visitor {
	return &Visitor{&antlr.BaseParseTreeVisitor{}, newScope(nil)}
}

type Visitor struct {
	*antlr.BaseParseTreeVisitor
	scope *scope
}

func (v *Visitor) VisitStartWrapped(ctx *parser.StartContext) types.WrappedValue {
	logger.Debug("VisitStart", ctx.GetText())
	return v.VisitWrapped(ctx.Block())
}

func (v *Visitor) VisitBlockWrapped(ctx *parser.BlockContext) types.WrappedValue {
	logger.Debug("VisitBlock", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitStatementWrapped(ctx *parser.StatementContext) types.WrappedValue {
	logger.Debug("VisitStatement", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitIntExprWrapped(ctx *parser.IntExprContext) types.WrappedValue {
	logger.Debug("VisitIntExpr", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitNullExprWrapped(ctx *parser.NullExprContext) types.WrappedValue {
	logger.Debug("VisitNullExpr", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitBooleanExprWrapped(ctx *parser.BooleanExprContext) types.WrappedValue {
	logger.Debug("VisitBooleanExpr", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitArithmeticWrapped(ctx *parser.ArithmeticContext) types.WrappedValue {
	logger.Debug("VisitArithmetic", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitStringExprWrapped(ctx *parser.StringExprContext) types.WrappedValue {
	logger.Debug("VisitStringExpr", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitFloatExprWrapped(ctx *parser.FloatExprContext) types.WrappedValue {
	logger.Debug("VisitFloatExpr", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitConditionWrapped(ctx *parser.ConditionContext) types.WrappedValue {
	logger.Debug("VisitCondition", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitNotExprWrapped(ctx *parser.NotExprContext) types.WrappedValue {
	logger.Debug("VisitNotExpr", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitFunctionDefExprWrapped(ctx *parser.FunctionDefExprContext) types.WrappedValue {
	logger.Debug("VisitFucntionDefExpr", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitFunctionCallWrapped(ctx *parser.FunctionCallContext) types.WrappedValue {
	logger.Debug("VisitFunctionCall", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitSymbolExprWrapped(ctx *parser.SymbolExprContext) types.WrappedValue {
	logger.Debug("VisitSymbolExpr", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitGroupingExprWrapped(ctx *parser.GroupingExprContext) types.WrappedValue {
	logger.Debug("VisitGroupingExpr", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitBooleanAlgebraWrapped(ctx *parser.BooleanAlgebraContext) types.WrappedValue {
	logger.Debug("VisitBooleanAlgebra", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitFunctionDefWrapped(ctx *parser.FunctionDefContext) types.WrappedValue {
	logger.Debug("VisitFunctionDef", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitFunctionArgsWrapped(ctx *parser.FunctionArgsContext) types.WrappedValue {
	logger.Debug("VisitFunctionArgs", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitLetAssignmentWrapped(ctx *parser.LetAssignmentContext) types.WrappedValue {
	logger.Debug("VisitLetAssignment", ctx.GetText())
	variable := ctx.SYMBOL().GetText()
	value := v.VisitWrapped(ctx.Expr())
	v.scope.setVar(variable, value)
	return nil
}

func (v *Visitor) VisitIfStatementWrapped(ctx *parser.IfStatementContext) types.WrappedValue {
	logger.Debug("VisitIfStatement", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitWhileLoopWrapped(ctx *parser.WhileLoopContext) types.WrappedValue {
	logger.Debug("VisitWhileLoop", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitForeachLoopWrapped(ctx *parser.ForeachLoopContext) types.WrappedValue {
	logger.Debug("VisitForeachLoop", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitForLoopWrapped(ctx *parser.ForLoopContext) types.WrappedValue {
	logger.Debug("VisitForLoop", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitElseifStatementWrapped(ctx *parser.ElseifStatementContext) types.WrappedValue {
	logger.Debug("VisitElseifStatement", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitElseStatementWrapped(ctx *parser.ElseStatementContext) types.WrappedValue {
	logger.Debug("VisitElseStatement", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitScopeBodyWrapped(ctx *parser.ScopeBodyContext) types.WrappedValue {
	logger.Debug("VisitScopeBody", ctx.GetText())

	logger.Debug("pushing new scope")
	v.scope = newScope(v.scope)

	v.VisitChildrenWrapped(ctx)

	logger.Debug("popping off scope")
	v.scope = v.scope.parentScope
	return nil
}

func (v *Visitor) VisitConditionBodyWrapped(ctx *parser.ConditionBodyContext) types.WrappedValue {
	logger.Debug("VisitConditionBody", ctx.GetText())
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitNotWrapped(ctx *parser.NotContext) types.WrappedValue {
	logger.Debug("VisitNot", ctx.GetText())
	expr := v.VisitWrapped(ctx.Expr())
	switch expr.(type) {
	case *types.BooleanValue:
		val := expr.(*types.BooleanValue).PrintValue()
		return types.NewBooleanValue(!val.(bool))
	default:
		err := errors.New("Not operator can only be applied to boolean. Got: " + ctx.Expr().GetText())
		panic(err)
	}
}

func (v *Visitor) VisitSymbolWrapped(ctx *parser.SymbolContext) types.WrappedValue {
	logger.Debug("VisitSymbol", ctx.GetText())
	// TODO: support sub-bindings
	return v.scope.resolveVariable(ctx.GetText())
}

func (v *Visitor) VisitStringLiteralWrapped(ctx *parser.StringLiteralContext) types.WrappedValue {
	logger.Debug("VisitStringLiteral", ctx.GetText())
	stringWithQuotes := ctx.STRING().GetText()
	val := stringWithQuotes[1 : len(stringWithQuotes)-1]
	return types.NewStringValue(val)
}

func (v *Visitor) VisitIntLiteralWrapped(ctx *parser.IntLiteralContext) types.WrappedValue {
	logger.Debug("VisitIntLiteral", ctx.GetText())
	val, err := strconv.ParseInt(ctx.GetText(), 0, 64)
	if err != nil {
		logger.Error("could not parse int", ctx.GetText(), err.Error())
		panic(err)
	}
	return types.NewIntValue(val)
}

func (v *Visitor) VisitFloatLiteralWrapped(ctx *parser.FloatLiteralContext) types.WrappedValue {
	logger.Debug("VisitFloatLiteral", ctx.GetText())
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		logger.Error("could not parse float", ctx.GetText(), err.Error())
		panic(err)
	}
	return types.NewFloatValue(val)
}

func (v *Visitor) VisitNullLiteralWrapped(ctx *parser.NullLiteralContext) types.WrappedValue {
	logger.Debug("VisitNullLiteral", ctx.GetText())
	return types.NewNullValue()
}

func (v *Visitor) VisitBooleanLiteralWrapped(ctx *parser.BooleanLiteralContext) types.WrappedValue {
	logger.Debug("VisitBooleanLiteral", ctx.GetText())
	if ctx.TRUE() != nil {
		return types.NewBooleanValue(true)
	}
	return types.NewBooleanValue(false)
}

func (v *Visitor) VisitGroupingWrapped(ctx *parser.GroupingContext) types.WrappedValue {
	logger.Debug("VisitGrouping", ctx.GetText())
	return v.VisitWrapped(ctx.Expr())
}
