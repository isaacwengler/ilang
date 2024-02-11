package visitor

import (
	"errors"
	parser "ilang/generated"
	"ilang/logger"
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

func (v *Visitor) Visit(t antlr.ParseTree) interface{} {
	return t.Accept(v)
}

func (v *Visitor) VisitChildren(node antlr.RuleNode) interface{} {
	var val any
	var curr any
	for _, n := range node.GetChildren() {
		 curr = v.Visit(n.(antlr.ParseTree))
         if curr != nil {
             // TODO: maybe not do this?
             val = curr
         }
	}

	return val
}

func (v *Visitor) VisitStart(ctx *parser.StartContext) interface{} {
	logger.Debug("VisitStart", ctx.GetText())
	return v.Visit(ctx.Block())
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	logger.Debug("VisitBlock", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	logger.Debug("VisitStatement", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	logger.Debug("VisitIntExpr", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitNullExpr(ctx *parser.NullExprContext) interface{} {
	logger.Debug("VisitNullExpr", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitBooleanExpr(ctx *parser.BooleanExprContext) interface{} {
	logger.Debug("VisitBooleanExpr", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArithmetic(ctx *parser.ArithmeticContext) interface{} {
	logger.Debug("VisitArithmetic", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStringExpr(ctx *parser.StringExprContext) interface{} {
	logger.Debug("VisitStringExpr", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFloatExpr(ctx *parser.FloatExprContext) interface{} {
	logger.Debug("VisitFloatExpr", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitCondition(ctx *parser.ConditionContext) interface{} {
	logger.Debug("VisitCondition", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	logger.Debug("VisitNotExpr", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctionDefExpr(ctx *parser.FunctionDefExprContext) interface{} {
	logger.Debug("VisitFucntionDefExpr", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	logger.Debug("VisitFunctionCall", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSymbolExpr(ctx *parser.SymbolExprContext) interface{} {
	logger.Debug("VisitSymbolExpr", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitGroupingExpr(ctx *parser.GroupingExprContext) interface{} {
	logger.Debug("VisitGroupingExpr", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitBooleanAlgebra(ctx *parser.BooleanAlgebraContext) interface{} {
	logger.Debug("VisitBooleanAlgebra", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctionDef(ctx *parser.FunctionDefContext) interface{} {
	logger.Debug("VisitFunctionDef", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctionArgs(ctx *parser.FunctionArgsContext) interface{} {
	logger.Debug("VisitFunctionArgs", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLetAssignment(ctx *parser.LetAssignmentContext) interface{} {
	logger.Debug("VisitLetAssignment", ctx.GetText())
	variable := ctx.SYMBOL().GetText()
	value := v.Visit(ctx.Expr())
	v.scope.setVar(variable, value)
	return nil
}

func (v *Visitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	logger.Debug("VisitIfStatement", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitWhileLoop(ctx *parser.WhileLoopContext) interface{} {
	logger.Debug("VisitWhileLoop", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitForeachLoop(ctx *parser.ForeachLoopContext) interface{} {
	logger.Debug("VisitForeachLoop", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitForLoop(ctx *parser.ForLoopContext) interface{} {
	logger.Debug("VisitForLoop", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitElseifStatement(ctx *parser.ElseifStatementContext) interface{} {
	logger.Debug("VisitElseifStatement", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitElseStatement(ctx *parser.ElseStatementContext) interface{} {
	logger.Debug("VisitElseStatement", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitScopeBody(ctx *parser.ScopeBodyContext) interface{} {
	logger.Debug("VisitScopeBody", ctx.GetText())

	logger.Debug("pushing new scope")
	v.scope = newScope(v.scope)

	v.VisitChildren(ctx)

	logger.Debug("popping off scope")
	v.scope = v.scope.parentScope
	return nil
}

func (v *Visitor) VisitConditionBody(ctx *parser.ConditionBodyContext) interface{} {
	logger.Debug("VisitConditionBody", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitNot(ctx *parser.NotContext) interface{} {
	logger.Debug("VisitNot", ctx.GetText())
	expr := v.Visit(ctx.Expr())
	switch expr.(type) {
	case bool:
		return !expr.(bool)
	default:
		err := errors.New("Not operator can only be applied to boolean. Got: " + ctx.Expr().GetText())
		logger.Error(err.Error())
		panic(err)

	}
}

func (v *Visitor) VisitSymbol(ctx *parser.SymbolContext) interface{} {
	logger.Debug("VisitSymbol", ctx.GetText())
	// TODO: support sub-bindings
	return v.scope.resolveVariable(ctx.GetText())
}

func (v *Visitor) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	logger.Debug("VisitStringLiteral", ctx.GetText())
	stringWithQuotes := ctx.STRING().GetText()
	return stringWithQuotes[1 : len(stringWithQuotes)-1]
}

func (v *Visitor) VisitIntLiteral(ctx *parser.IntLiteralContext) interface{} {
	logger.Debug("VisitIntLiteral", ctx.GetText())
	val, err := strconv.ParseInt(ctx.GetText(), 0, 64)
	if err != nil {
		logger.Error("could not parse float", ctx.GetText(), err.Error())
		panic(err)
	}
	return val
}

func (v *Visitor) VisitFloatLiteral(ctx *parser.FloatLiteralContext) interface{} {
	logger.Debug("VisitFloatLiteral", ctx.GetText())
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		logger.Error("could not parse float", ctx.GetText(), err.Error())
		panic(err)
	}
	return val
}

func (v *Visitor) VisitNullLiteral(ctx *parser.NullLiteralContext) interface{} {
	logger.Debug("VisitNullLiteral", ctx.GetText())
	return nil
}

func (v *Visitor) VisitBooleanLiteral(ctx *parser.BooleanLiteralContext) interface{} {
	logger.Debug("VisitBooleanLiteral", ctx.GetText())
	if ctx.TRUE() != nil {
		return true
	}
	return false
}

func (v *Visitor) VisitGrouping(ctx *parser.GroupingContext) interface{} {
	logger.Debug("VisitGrouping", ctx.GetText())
	return v.Visit(ctx.Expr())
}
