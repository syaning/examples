// Code generated from CalculatorParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CalculatorParser
import "github.com/antlr4-go/antlr/v4"

// BaseCalculatorParserListener is a complete listener for a parse tree produced by CalculatorParser.
type BaseCalculatorParserListener struct{}

var _ CalculatorParserListener = &BaseCalculatorParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalculatorParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalculatorParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalculatorParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalculatorParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseCalculatorParserListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseCalculatorParserListener) ExitProg(ctx *ProgContext) {}

// EnterMulDivMod is called when production MulDivMod is entered.
func (s *BaseCalculatorParserListener) EnterMulDivMod(ctx *MulDivModContext) {}

// ExitMulDivMod is called when production MulDivMod is exited.
func (s *BaseCalculatorParserListener) ExitMulDivMod(ctx *MulDivModContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseCalculatorParserListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseCalculatorParserListener) ExitNumber(ctx *NumberContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseCalculatorParserListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseCalculatorParserListener) ExitAddSub(ctx *AddSubContext) {}

// EnterParens is called when production Parens is entered.
func (s *BaseCalculatorParserListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BaseCalculatorParserListener) ExitParens(ctx *ParensContext) {}

// EnterExp is called when production Exp is entered.
func (s *BaseCalculatorParserListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production Exp is exited.
func (s *BaseCalculatorParserListener) ExitExp(ctx *ExpContext) {}
