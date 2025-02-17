// Code generated from CalculatorParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CalculatorParser
import "github.com/antlr4-go/antlr/v4"

// CalculatorParserListener is a complete listener for a parse tree produced by CalculatorParser.
type CalculatorParserListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterMulDivMod is called when entering the MulDivMod production.
	EnterMulDivMod(c *MulDivModContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterExp is called when entering the Exp production.
	EnterExp(c *ExpContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitMulDivMod is called when exiting the MulDivMod production.
	ExitMulDivMod(c *MulDivModContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitExp is called when exiting the Exp production.
	ExitExp(c *ExpContext)
}
