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
	return v.VisitWrapped(ctx.Block())
}

func (v *Visitor) VisitBlockWrapped(ctx *parser.BlockContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitStatementWrapped(ctx *parser.StatementContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitIntExprWrapped(ctx *parser.IntExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitNullExprWrapped(ctx *parser.NullExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitArrayExprWrapped(ctx *parser.ArrayExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitBooleanExprWrapped(ctx *parser.BooleanExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitMapExprWrapped(ctx *parser.MapExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitArithmeticWrapped(ctx *parser.ArithmeticContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitStringExprWrapped(ctx *parser.StringExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitFloatExprWrapped(ctx *parser.FloatExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitConditionWrapped(ctx *parser.ConditionContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitNotExprWrapped(ctx *parser.NotExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitFunctionDefExprWrapped(ctx *parser.FunctionDefExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitFunctionCallWrapped(ctx *parser.FunctionCallContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitSymbolExprWrapped(ctx *parser.SymbolExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitGroupingExprWrapped(ctx *parser.GroupingExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitBooleanAlgebraWrapped(ctx *parser.BooleanAlgebraContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitFunctionDefWrapped(ctx *parser.FunctionDefContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitFunctionArgsWrapped(ctx *parser.FunctionArgsContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitAssignmentWrapped(ctx *parser.AssignmentContext) types.WrappedValue {
	isReassignment := ctx.LET() == nil
	variable := ctx.SYMBOL().GetText()
	value := v.VisitWrapped(ctx.Expr())

	if isReassignment {
		v.scope.reassignVariable(variable, value)
		return value
	}

	_, isPresentInCurrentScope := v.scope.getVar(variable)
	if isPresentInCurrentScope {
		panic(errors.New("variable '" + variable + "' already defined in scope"))
	}
	v.scope.setVar(variable, value)
	return value
}

func (v *Visitor) VisitIfStatementWrapped(ctx *parser.IfStatementContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitWhileLoopWrapped(ctx *parser.WhileLoopContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitForeachLoopWrapped(ctx *parser.ForeachLoopContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitForLoopWrapped(ctx *parser.ForLoopContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitReturnWrapped(ctx *parser.ReturnContext) types.WrappedValue {
	return v.VisitWrapped(ctx.Expr())
}

func (v *Visitor) VisitElseifStatementWrapped(ctx *parser.ElseifStatementContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitElseStatementWrapped(ctx *parser.ElseStatementContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitScopeBodyWrapped(ctx *parser.ScopeBodyContext) types.WrappedValue {
	logger.Debug("pushing new scope")
	v.scope = newScope(v.scope)

	val := v.VisitChildrenWrapped(ctx)

	logger.Debug("popping off scope")
	v.scope = v.scope.parentScope
	return val
}

func (v *Visitor) VisitConditionBodyWrapped(ctx *parser.ConditionBodyContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitNotWrapped(ctx *parser.NotContext) types.WrappedValue {
	expr := v.VisitWrapped(ctx.Expr())
	switch expr.(type) {
	case *types.BooleanValue:
		val := expr.(*types.BooleanValue).GetValue()
		return types.NewBooleanValue(!val)
	default:
		err := errors.New("Not operator can only be applied to boolean. Got: " + ctx.Expr().GetText())
		panic(err)
	}
}

func (v *Visitor) VisitSymbolWrapped(ctx *parser.SymbolContext) types.WrappedValue {
	return v.scope.resolveVariable(ctx.GetText())
}

func (v *Visitor) VisitStringLiteralWrapped(ctx *parser.StringLiteralContext) types.WrappedValue {
	stringWithQuotes := ctx.STRING().GetText()
	val := stringWithQuotes[1 : len(stringWithQuotes)-1]
	return types.NewStringValue(val)
}

func (v *Visitor) VisitIntLiteralWrapped(ctx *parser.IntLiteralContext) types.WrappedValue {
	val, err := strconv.ParseInt(ctx.GetText(), 0, 64)
	if err != nil {
		logger.Error("could not parse int", ctx.GetText(), err.Error())
		panic(err)
	}
	return types.NewIntValue(val)
}

func (v *Visitor) VisitFloatLiteralWrapped(ctx *parser.FloatLiteralContext) types.WrappedValue {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		logger.Error("could not parse float", ctx.GetText(), err.Error())
		panic(err)
	}
	return types.NewFloatValue(val)
}

func (v *Visitor) VisitNullLiteralWrapped(ctx *parser.NullLiteralContext) types.WrappedValue {
	return types.NewNullValue()
}

func (v *Visitor) VisitBooleanLiteralWrapped(ctx *parser.BooleanLiteralContext) types.WrappedValue {
	if ctx.TRUE() != nil {
		return types.NewBooleanValue(true)
	}
	return types.NewBooleanValue(false)
}

func (v *Visitor) VisitArrayLiteralWrapped(ctx *parser.ArrayLiteralContext) types.WrappedValue {
	arr := []types.WrappedValue{}
	items := ctx.GetItems()
	if items != nil {
		for _, item := range ctx.GetItems() {
			arr = append(arr, v.VisitWrapped(item))
		}
	}
	return types.NewArrayValue(arr)
}

func (v *Visitor) VisitMapLiteralWrapped(ctx *parser.MapLiteralContext) types.WrappedValue {
	m := make(map[any]types.WrappedValue)

	items := ctx.GetItems()
	if items != nil {
		for _, item := range ctx.GetItems() {
			key := v.VisitWrapped(item.MapKey())
			value := v.VisitWrapped(item.Expr())

			var resolvedKey any
			switch key.(type) {
			case *types.StringValue:
				resolvedKey = key.(*types.StringValue).GetValue()
			case *types.IntValue:
				resolvedKey = key.(*types.IntValue).GetValue()
			default:
				err := errors.New("key for map must be string or int")
				panic(err)
			}

			m[resolvedKey] = value
		}
	}

	return types.NewMapValue(m)
}

func (v *Visitor) VisitMapLiteralItemWrapped(ctx *parser.MapLiteralItemContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitMapKeyWrapped(ctx *parser.MapKeyContext) types.WrappedValue {
	if ctx.SYMBOL() != nil {
		return types.NewStringValue(ctx.SYMBOL().GetText())
	}
	return v.VisitWrapped(ctx.Expr())
}

func (v *Visitor) VisitGroupingWrapped(ctx *parser.GroupingContext) types.WrappedValue {
	return v.VisitWrapped(ctx.Expr())
}
