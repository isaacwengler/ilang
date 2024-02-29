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
	v := &Visitor{&antlr.BaseParseTreeVisitor{}, types.NewScope(nil)}
	v.registerGlobalFunctions()
	v.scope = types.NewScope(v.scope)
	return v
}

type Visitor struct {
	*antlr.BaseParseTreeVisitor
	scope *types.Scope
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
	first := v.VisitWrapped(ctx.Expr(0))
	second := v.VisitWrapped(ctx.Expr(1))
	op := ctx.ARITHMETIC_OP().GetText()

	switch first.(type) {
	case *types.IntValue:
		return first.(*types.IntValue).Arithmetic(op, second)
	case *types.FloatValue:
		return first.(*types.FloatValue).Arithmetic(op, second)
	case *types.StringValue:
		return first.(*types.StringValue).Arithmetic(op, second)
	default:
		err := errors.New("Arithmetic operator '" + op + "' not defined for value type")
		panic(err)
	}
}

func (v *Visitor) VisitStringExprWrapped(ctx *parser.StringExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitFloatExprWrapped(ctx *parser.FloatExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitConditionWrapped(ctx *parser.ConditionContext) types.WrappedValue {
	first := v.VisitWrapped(ctx.Expr(0))
	second := v.VisitWrapped(ctx.Expr(1))

	return first.Comparison(ctx.CONDITIONAL_OP().GetText(), second)
}

func (v *Visitor) VisitNotExprWrapped(ctx *parser.NotExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitFunctionDefExprWrapped(ctx *parser.FunctionDefExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitFunctionCallWrapped(ctx *parser.FunctionCallContext) types.WrappedValue {
	expr := v.VisitWrapped(ctx.Expr())
	switch expr.(type) {
	case *types.FunctionValue:
		function := expr.(*types.FunctionValue)

		logger.Debug("Calling function, switching to closure scope")
		currentScope := v.scope
		v.scope = function.ClosureScope

		argsCtx := ctx.FunctionArgs().GetArgs()
		if len(argsCtx) != len(function.Args) {
			logger.Error(len(argsCtx), len(function.Args))
			logger.Error(argsCtx)
			logger.Error(function.Args)
			err := errors.New("Received different number of args than expected")
			panic(err)
		}

		for i := range argsCtx {
			v.scope.SetVar(function.Args[i], v.VisitWrapped(argsCtx[i]))
		}

		res := v.VisitWrapped(function.Def)
		logger.Debug("Function called, now returning to current scope")
		v.scope = currentScope
		return res
	case *types.LibFunctionValue:
		argsCtx := ctx.FunctionArgs().GetArgs()
		args := make([]types.WrappedValue, len(argsCtx))
		for i, a := range argsCtx {
			args[i] = v.VisitWrapped(a)
		}
		return expr.(*types.LibFunctionValue).Call(args)
	default:
		err := errors.New("Attempted to invoke function on value that is not a function")
		panic(err)
	}
}

func (v *Visitor) VisitSymbolExprWrapped(ctx *parser.SymbolExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitGroupingExprWrapped(ctx *parser.GroupingExprContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitBooleanAlgebraWrapped(ctx *parser.BooleanAlgebraContext) types.WrappedValue {
	first := v.VisitWrapped(ctx.Expr(0))
	second := v.VisitWrapped(ctx.Expr(1))
	op := ctx.BOOLEAN_OP().GetText()

	var firstBool, secondBool bool
	switch first.(type) {
	case *types.BooleanValue:
		firstBool = first.(*types.BooleanValue).GetValue()
	default:
		err := errors.New("Boolean operator '" + op + "' can only be applied to boolean value")
		panic(err)
	}

	if op == "||" && firstBool {
		return types.NewBooleanValue(true)
	} else if op == "&&" && !firstBool {
		return types.NewBooleanValue(false)
	}

	switch second.(type) {
	case *types.BooleanValue:
		secondBool = second.(*types.BooleanValue).GetValue()
	default:
		err := errors.New("Boolean operator '" + op + "' can only be applied to boolean value")
		panic(err)
	}

	if op == "&&" {
		return types.NewBooleanValue(firstBool && secondBool)
	} else {
		return types.NewBooleanValue(firstBool || secondBool)
	}
}

func (v *Visitor) VisitFunctionDefWrapped(ctx *parser.FunctionDefContext) types.WrappedValue {
	symbols := ctx.FunctionDefArgs().AllSYMBOL()
	args := make([]string, len(symbols))
	for i, s := range symbols {
		args[i] = s.GetText()
	}

	return types.NewFunctionValue(args, ctx.ScopeBody(), v.scope)
}

func (v *Visitor) VisitFunctionArgsWrapped(ctx *parser.FunctionArgsContext) types.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitAssignmentWrapped(ctx *parser.AssignmentContext) types.WrappedValue {
	isReassignment := ctx.LET() == nil
	variable := ctx.SYMBOL().GetText()
	value := v.VisitWrapped(ctx.Expr())

	if isReassignment {
		v.scope.ReassignVariable(variable, value)
		return value
	}

	_, isPresentInCurrentScope := v.scope.GetVar(variable)
	if isPresentInCurrentScope {
		panic(errors.New("variable '" + variable + "' already defined in scope"))
	}
	v.scope.SetVar(variable, value)
	return value
}

func (v *Visitor) VisitIfStatementWrapped(ctx *parser.IfStatementContext) types.WrappedValue {
	condition := evalCondition(v.VisitWrapped(ctx.ConditionBody()))
	if condition {
		return v.VisitWrapped(ctx.ScopeBody())
	}

	for _, elseif := range ctx.AllElseifStatement() {
		condition := evalCondition(v.VisitWrapped(elseif.ConditionBody()))
		if condition {
			return v.VisitWrapped(elseif.ScopeBody())
		}
	}

	if ctx.ElseStatement() != nil {
		return v.VisitWrapped(ctx.ElseStatement().ScopeBody())
	}

	return types.NewNullValue()
}

func (v *Visitor) VisitWhileLoopWrapped(ctx *parser.WhileLoopContext) types.WrappedValue {
	var last types.WrappedValue = types.NewNullValue()
	for evalCondition(v.VisitWrapped(ctx.ConditionBody())) {
		last = v.VisitWrapped(ctx.ScopeBody())
	}
	return last
}

func (v *Visitor) VisitForeachLoopWrapped(ctx *parser.ForeachLoopContext) types.WrappedValue {
	logger.Debug("pushing new scope for foreach variables")
	v.scope = types.NewScope(v.scope)

	expr := v.VisitWrapped(ctx.Expr())
	switch expr.(type) {
	case *types.ArrayValue:
		items := expr.(*types.ArrayValue).GetValue()
		var last types.WrappedValue = types.NewNullValue()
		for _, item := range items {
			v.scope.SetVar(ctx.SYMBOL().GetText(), item)
			last = v.VisitWrapped(ctx.ScopeBody())
		}

		logger.Debug("popping off scope with foreach variables")
		v.scope = v.scope.ParentScope
		return last
	default:
		err := errors.New("foreach loop can only be used with array")
		panic(err)
	}
}

func (v *Visitor) VisitForLoopWrapped(ctx *parser.ForLoopContext) types.WrappedValue {
	logger.Debug("pushing new scope for for variables")
	v.scope = types.NewScope(v.scope)

	v.VisitWrapped(ctx.GetInit())
	var last types.WrappedValue = types.NewNullValue()
	for evalCondition(v.VisitWrapped(ctx.GetCond())) {
		last = v.VisitWrapped(ctx.ScopeBody())
		v.VisitWrapped(ctx.GetStep())
	}

	logger.Debug("popping off scope with for variables")
	v.scope = types.NewScope(v.scope)
	return last
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
	v.scope = types.NewScope(v.scope)

	val := v.VisitChildrenWrapped(ctx)

	logger.Debug("popping off scope")
	v.scope = v.scope.ParentScope
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
	return v.scope.ResolveVariable(ctx.GetText())
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

func evalCondition(c types.WrappedValue) bool {
	switch c.(type) {
	case *types.BooleanValue:
		return c.(*types.BooleanValue).GetValue()
	default:
		err := errors.New("Condition body must have boolean value")
		panic(err)
	}
}
