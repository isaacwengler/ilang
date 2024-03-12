package visitor

import (
	"errors"
	parser "ilang/generated"
	"ilang/library"
	"ilang/logger"
	"ilang/model"
	"ilang/scope"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

func NewVisitor() *Visitor {
	v := &Visitor{&antlr.BaseParseTreeVisitor{}, scope.NewScope(nil)}
	v.registerGlobalFunctions()
	v.scope = scope.NewScope(v.scope)
	return v
}

type Visitor struct {
	*antlr.BaseParseTreeVisitor
	scope *scope.Scope
}

func (v *Visitor) VisitStartWrapped(ctx *parser.StartContext) model.WrappedValue {
	return v.VisitWrapped(ctx.Block())
}

func (v *Visitor) VisitBlockWrapped(ctx *parser.BlockContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitStatementWrapped(ctx *parser.StatementContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitIntExprWrapped(ctx *parser.IntExprContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitNullExprWrapped(ctx *parser.NullExprContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitArrayExprWrapped(ctx *parser.ArrayExprContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitBooleanExprWrapped(ctx *parser.BooleanExprContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitMapExprWrapped(ctx *parser.MapExprContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitArithmeticWrapped(ctx *parser.ArithmeticContext) model.WrappedValue {
	first := v.VisitWrapped(ctx.Expr(0))
	second := v.VisitWrapped(ctx.Expr(1))
	op := ctx.ARITHMETIC_OP().GetText()

	switch first.(type) {
	case *library.IntValue:
		return first.(*library.IntValue).Arithmetic(op, second)
	case *library.FloatValue:
		return first.(*library.FloatValue).Arithmetic(op, second)
	case *library.StringValue:
		return first.(*library.StringValue).Arithmetic(op, second)
	default:
		err := errors.New("Arithmetic operator '" + op + "' not defined for value type")
		panic(err)
	}
}

func (v *Visitor) VisitStringExprWrapped(ctx *parser.StringExprContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitFloatExprWrapped(ctx *parser.FloatExprContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitConditionWrapped(ctx *parser.ConditionContext) model.WrappedValue {
	first := v.VisitWrapped(ctx.Expr(0))
	second := v.VisitWrapped(ctx.Expr(1))

	return first.Comparison(ctx.CONDITIONAL_OP().GetText(), second)
}

func (v *Visitor) VisitNotExprWrapped(ctx *parser.NotExprContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitFunctionDefExprWrapped(ctx *parser.FunctionDefExprContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitFunctionCallWrapped(ctx *parser.FunctionCallContext) model.WrappedValue {
	expr := v.VisitWrapped(ctx.Expr())
	switch expr.(type) {
	case *library.FunctionValue:
		function := expr.(*library.FunctionValue)

		currentScope := v.scope
        functionScope := scope.NewScope(function.ClosureScope)

		argsCtx := ctx.FunctionArgs().GetArgs()
		if len(argsCtx) != len(function.Args) {
			err := errors.New("Received different number of args than expected")
			panic(err)
		}

		for i := range argsCtx {
			functionScope.SetVar(function.Args[i], v.VisitWrapped(argsCtx[i]))
		}

		logger.Debug("Calling function, switching to closure scope")
        v.scope = functionScope

		res := v.VisitWrapped(function.Def)
		logger.Debug("Function called, now returning to current scope")
		v.scope = currentScope
		return res
	case *library.LibFunctionValue:
		argsCtx := ctx.FunctionArgs().GetArgs()
		args := make([]model.WrappedValue, len(argsCtx))
		for i, a := range argsCtx {
			args[i] = v.VisitWrapped(a)
		}
		return expr.(*library.LibFunctionValue).Call(args)
	default:
		err := errors.New("Attempted to invoke function on value that is not a function")
		panic(err)
	}
}

func (v *Visitor) VisitSymbolExprWrapped(ctx *parser.SymbolExprContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitGroupingExprWrapped(ctx *parser.GroupingExprContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitBooleanAlgebraWrapped(ctx *parser.BooleanAlgebraContext) model.WrappedValue {
	first := v.VisitWrapped(ctx.Expr(0))
	second := v.VisitWrapped(ctx.Expr(1))
	op := ctx.BOOLEAN_OP().GetText()

	var firstBool, secondBool bool
	switch first.(type) {
	case *library.BooleanValue:
		firstBool = first.(*library.BooleanValue).GetValue()
	default:
		err := errors.New("Boolean operator '" + op + "' can only be applied to boolean value")
		panic(err)
	}

	if op == "||" && firstBool {
		return library.NewBooleanValue(true)
	} else if op == "&&" && !firstBool {
		return library.NewBooleanValue(false)
	}

	switch second.(type) {
	case *library.BooleanValue:
		secondBool = second.(*library.BooleanValue).GetValue()
	default:
		err := errors.New("Boolean operator '" + op + "' can only be applied to boolean value")
		panic(err)
	}

	if op == "&&" {
		return library.NewBooleanValue(firstBool && secondBool)
	} else {
		return library.NewBooleanValue(firstBool || secondBool)
	}
}

func (v *Visitor) VisitFunctionDefWrapped(ctx *parser.FunctionDefContext) model.WrappedValue {
	symbols := ctx.FunctionDefArgs().AllSYMBOL()
	args := make([]string, len(symbols))
	for i, s := range symbols {
		args[i] = s.GetText()
	}

	return library.NewFunctionValue(args, ctx.ScopeBody(), v.scope)
}

func (v *Visitor) VisitFunctionArgsWrapped(ctx *parser.FunctionArgsContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitAssignmentWrapped(ctx *parser.AssignmentContext) model.WrappedValue {
	variable := ctx.SYMBOL().GetText()
	value := v.VisitWrapped(ctx.Expr())

	_, isPresentInCurrentScope := v.scope.GetVar(variable)
	if isPresentInCurrentScope {
		panic(errors.New("variable '" + variable + "' already defined in scope"))
	}
	v.scope.SetVar(variable, value)
	return value
}

func (v *Visitor) VisitReassignmentWrapped(ctx *parser.ReassignmentContext) model.WrappedValue {
	symbol := ctx.SYMBOL().GetText()
	value := v.VisitWrapped(ctx.Expr())

	children := make([]model.WrappedValue, 0)
	curr := ctx.SymbolChild()
	for curr != nil {
		if curr.SYMBOL() != nil {
			children = append(children, library.NewStringValue(curr.SYMBOL().GetText()))
		} else {
			children = append(children, v.VisitWrapped(curr.Expr()))
		}
		curr = curr.SymbolChild()
	}

	v.scope.ReassignVariable(symbol, value, children)
	return value
}

func (v *Visitor) VisitIfStatementWrapped(ctx *parser.IfStatementContext) model.WrappedValue {
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

	return library.NewNullValue()
}

func (v *Visitor) VisitWhileLoopWrapped(ctx *parser.WhileLoopContext) model.WrappedValue {
	var last model.WrappedValue = library.NewNullValue()
	for evalCondition(v.VisitWrapped(ctx.ConditionBody())) {
		last = v.VisitWrapped(ctx.ScopeBody())
	}
	return last
}

func (v *Visitor) VisitForeachLoopWrapped(ctx *parser.ForeachLoopContext) model.WrappedValue {
	logger.Debug("pushing new scope for foreach variables")
	v.scope = scope.NewScope(v.scope)

	expr := v.VisitWrapped(ctx.Expr())
	switch expr.(type) {
	case *library.ArrayValue:
		items := expr.(*library.ArrayValue).GetValue()
		var last model.WrappedValue = library.NewNullValue()
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

func (v *Visitor) VisitForLoopWrapped(ctx *parser.ForLoopContext) model.WrappedValue {
	logger.Debug("pushing new scope for for variables")
	v.scope = scope.NewScope(v.scope)

	v.VisitWrapped(ctx.GetInit())
	var last model.WrappedValue = library.NewNullValue()
	for evalCondition(v.VisitWrapped(ctx.GetCond())) {
		last = v.VisitWrapped(ctx.ScopeBody())
		v.VisitWrapped(ctx.GetStep())
	}

	logger.Debug("popping off scope with for variables")
	v.scope = scope.NewScope(v.scope)
	return last
}

func (v *Visitor) VisitReturnWrapped(ctx *parser.ReturnContext) model.WrappedValue {
	return v.VisitWrapped(ctx.Expr())
}

func (v *Visitor) VisitElseifStatementWrapped(ctx *parser.ElseifStatementContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitElseStatementWrapped(ctx *parser.ElseStatementContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitScopeBodyWrapped(ctx *parser.ScopeBodyContext) model.WrappedValue {
	logger.Debug("pushing new scope")
	v.scope = scope.NewScope(v.scope)

	val := v.VisitChildrenWrapped(ctx)

	logger.Debug("popping off scope")
	v.scope = v.scope.ParentScope
	return val
}

func (v *Visitor) VisitConditionBodyWrapped(ctx *parser.ConditionBodyContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitNotWrapped(ctx *parser.NotContext) model.WrappedValue {
	expr := v.VisitWrapped(ctx.Expr())
	switch expr.(type) {
	case *library.BooleanValue:
		val := expr.(*library.BooleanValue).GetValue()
		return library.NewBooleanValue(!val)
	default:
		err := errors.New("Not operator can only be applied to boolean. Got: " + ctx.Expr().GetText())
		panic(err)
	}
}

func (v *Visitor) VisitSymbolWrapped(ctx *parser.SymbolContext) model.WrappedValue {
	return v.scope.ResolveVariable(ctx.SYMBOL().GetText())
}

func (v *Visitor) VisitStringLiteralWrapped(ctx *parser.StringLiteralContext) model.WrappedValue {
	stringWithQuotes := ctx.STRING().GetText()
	val := stringWithQuotes[1 : len(stringWithQuotes)-1]
	return library.NewStringValue(val)
}

func (v *Visitor) VisitIntLiteralWrapped(ctx *parser.IntLiteralContext) model.WrappedValue {
	val, err := strconv.ParseInt(ctx.GetText(), 0, 64)
	if err != nil {
		logger.Error("could not parse int", ctx.GetText(), err.Error())
		panic(err)
	}
	return library.NewIntValue(val)
}

func (v *Visitor) VisitFloatLiteralWrapped(ctx *parser.FloatLiteralContext) model.WrappedValue {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		logger.Error("could not parse float", ctx.GetText(), err.Error())
		panic(err)
	}
	return library.NewFloatValue(val)
}

func (v *Visitor) VisitNullLiteralWrapped(ctx *parser.NullLiteralContext) model.WrappedValue {
	return library.NewNullValue()
}

func (v *Visitor) VisitBooleanLiteralWrapped(ctx *parser.BooleanLiteralContext) model.WrappedValue {
	if ctx.TRUE() != nil {
		return library.NewBooleanValue(true)
	}
	return library.NewBooleanValue(false)
}

func (v *Visitor) VisitArrayLiteralWrapped(ctx *parser.ArrayLiteralContext) model.WrappedValue {
	arr := []model.WrappedValue{}
	items := ctx.GetItems()
	if items != nil {
		for _, item := range ctx.GetItems() {
			arr = append(arr, v.VisitWrapped(item))
		}
	}
	return library.NewArrayValue(arr)
}

func (v *Visitor) VisitMapLiteralWrapped(ctx *parser.MapLiteralContext) model.WrappedValue {
	m := make(map[any]model.WrappedValue)

	items := ctx.GetItems()
	if items != nil {
		for _, item := range ctx.GetItems() {
			key := v.VisitWrapped(item.MapKey())
			value := v.VisitWrapped(item.Expr())

			var resolvedKey any
			switch key.(type) {
			case *library.StringValue:
				resolvedKey = key.(*library.StringValue).GetValue()
			case *library.IntValue:
				resolvedKey = key.(*library.IntValue).GetValue()
			default:
				err := errors.New("key for map must be string or int")
				panic(err)
			}

			m[resolvedKey] = value
		}
	}

	return library.NewMapValue(m)
}

func (v *Visitor) VisitMapLiteralItemWrapped(ctx *parser.MapLiteralItemContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}

func (v *Visitor) VisitMapKeyWrapped(ctx *parser.MapKeyContext) model.WrappedValue {
	if ctx.SYMBOL() != nil {
		return library.NewStringValue(ctx.SYMBOL().GetText())
	}
	return v.VisitWrapped(ctx.Expr())
}

func (v *Visitor) VisitGroupingWrapped(ctx *parser.GroupingContext) model.WrappedValue {
	return v.VisitWrapped(ctx.Expr())
}

func evalCondition(c model.WrappedValue) bool {
	switch c.(type) {
	case *library.BooleanValue:
		return c.(*library.BooleanValue).GetValue()
	default:
		err := errors.New("Condition body must have boolean value")
		panic(err)
	}
}

func (v *Visitor) VisitPropertyWrapped(ctx *parser.PropertyContext) model.WrappedValue {
	curr := v.VisitWrapped(ctx.Expr())
	property := ctx.SYMBOL().GetText()
	return curr.GetChild(library.NewStringValue(property))
}

func (v *Visitor) VisitComputedPropertyWrapped(ctx *parser.ComputedPropertyContext) model.WrappedValue {
	curr := v.VisitWrapped(ctx.Expr(0))
	computedProperty := v.VisitWrapped(ctx.Expr(1))
	return curr.GetChild(computedProperty)
}

func (v *Visitor) VisitSymbolChildWrapped(ctx *parser.SymbolChildContext) model.WrappedValue {
	return v.VisitChildrenWrapped(ctx)
}
