// Code generated from CalculatorParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CalculatorParser
import "github.com/antlr4-go/antlr/v4"

type BaseCalculatorParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCalculatorParserVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorParserVisitor) VisitMulDivMod(ctx *MulDivModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorParserVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorParserVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorParserVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorParserVisitor) VisitExp(ctx *ExpContext) interface{} {
	return v.VisitChildren(ctx)
}
