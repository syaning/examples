// Code generated from CalculatorParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CalculatorParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CalculatorParser.
type CalculatorParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalculatorParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by CalculatorParser#MulDivMod.
	VisitMulDivMod(ctx *MulDivModContext) interface{}

	// Visit a parse tree produced by CalculatorParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by CalculatorParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by CalculatorParser#Parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by CalculatorParser#Exp.
	VisitExp(ctx *ExpContext) interface{}
}
