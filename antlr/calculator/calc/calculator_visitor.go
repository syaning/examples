package calc

import (
	"math"
	"strconv"

	"calculator/parsing"

	"github.com/antlr4-go/antlr/v4"
)

type CalculatorVisitor struct {
	*parsing.BaseCalculatorParserVisitor
}

func NewCalculatorVisitor() *CalculatorVisitor {
	return &CalculatorVisitor{
		BaseCalculatorParserVisitor: &parsing.BaseCalculatorParserVisitor{},
	}
}

func (v *CalculatorVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *CalculatorVisitor) VisitProg(ctx *parsing.ProgContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *CalculatorVisitor) VisitNumber(ctx *parsing.NumberContext) interface{} {
	val, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return val
}

func (v *CalculatorVisitor) VisitExp(ctx *parsing.ExpContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(int64)
	right := v.Visit(ctx.Expr(1)).(int64)
	return int64(math.Pow(float64(left), float64(right)))
}

func (v *CalculatorVisitor) VisitMulDivMod(ctx *parsing.MulDivModContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(int64)
	right := v.Visit(ctx.Expr(1)).(int64)
	if ctx.GetOp().GetTokenType() == parsing.CalculatorParserMUL {
		return left * right
	}
	if ctx.GetOp().GetTokenType() == parsing.CalculatorParserDIV {
		return left / right
	}
	return left % right
}

func (v *CalculatorVisitor) VisitAddSub(ctx *parsing.AddSubContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(int64)
	right := v.Visit(ctx.Expr(1)).(int64)
	if ctx.GetOp().GetTokenType() == parsing.CalculatorParserADD {
		return left + right
	}
	return left - right
}

func (v *CalculatorVisitor) VisitParens(ctx *parsing.ParensContext) interface{} {
	return v.Visit(ctx.Expr())
}
