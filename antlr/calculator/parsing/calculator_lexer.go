// Code generated from CalculatorLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CalculatorLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CalculatorLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calculatorlexerLexerInit() {
	staticData := &CalculatorLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'+'", "'-'", "'*'", "'/'", "'%'", "'**'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "ADD", "SUB", "MUL", "DIV", "MOD", "EXP", "LEFT_PAREN", "RIGHT_PAREN",
		"NUMBER", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"ADD", "SUB", "MUL", "DIV", "MOD", "EXP", "LEFT_PAREN", "RIGHT_PAREN",
		"NUMBER", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 10, 52, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 5, 8, 41, 8, 8, 10, 8, 12, 8, 44,
		9, 8, 1, 9, 4, 9, 47, 8, 9, 11, 9, 12, 9, 48, 1, 9, 1, 9, 0, 0, 10, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 1, 0, 3,
		1, 0, 49, 57, 1, 0, 48, 57, 3, 0, 9, 10, 12, 13, 32, 32, 53, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 1, 21, 1, 0, 0, 0, 3, 23, 1, 0, 0, 0, 5, 25,
		1, 0, 0, 0, 7, 27, 1, 0, 0, 0, 9, 29, 1, 0, 0, 0, 11, 31, 1, 0, 0, 0, 13,
		34, 1, 0, 0, 0, 15, 36, 1, 0, 0, 0, 17, 38, 1, 0, 0, 0, 19, 46, 1, 0, 0,
		0, 21, 22, 5, 43, 0, 0, 22, 2, 1, 0, 0, 0, 23, 24, 5, 45, 0, 0, 24, 4,
		1, 0, 0, 0, 25, 26, 5, 42, 0, 0, 26, 6, 1, 0, 0, 0, 27, 28, 5, 47, 0, 0,
		28, 8, 1, 0, 0, 0, 29, 30, 5, 37, 0, 0, 30, 10, 1, 0, 0, 0, 31, 32, 5,
		42, 0, 0, 32, 33, 5, 42, 0, 0, 33, 12, 1, 0, 0, 0, 34, 35, 5, 40, 0, 0,
		35, 14, 1, 0, 0, 0, 36, 37, 5, 41, 0, 0, 37, 16, 1, 0, 0, 0, 38, 42, 7,
		0, 0, 0, 39, 41, 7, 1, 0, 0, 40, 39, 1, 0, 0, 0, 41, 44, 1, 0, 0, 0, 42,
		40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 18, 1, 0, 0, 0, 44, 42, 1, 0, 0,
		0, 45, 47, 7, 2, 0, 0, 46, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 46,
		1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 6, 9, 0, 0,
		51, 20, 1, 0, 0, 0, 3, 0, 42, 48, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CalculatorLexerInit initializes any static state used to implement CalculatorLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCalculatorLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalculatorLexerInit() {
	staticData := &CalculatorLexerLexerStaticData
	staticData.once.Do(calculatorlexerLexerInit)
}

// NewCalculatorLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCalculatorLexer(input antlr.CharStream) *CalculatorLexer {
	CalculatorLexerInit()
	l := new(CalculatorLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CalculatorLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "CalculatorLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalculatorLexer tokens.
const (
	CalculatorLexerADD         = 1
	CalculatorLexerSUB         = 2
	CalculatorLexerMUL         = 3
	CalculatorLexerDIV         = 4
	CalculatorLexerMOD         = 5
	CalculatorLexerEXP         = 6
	CalculatorLexerLEFT_PAREN  = 7
	CalculatorLexerRIGHT_PAREN = 8
	CalculatorLexerNUMBER      = 9
	CalculatorLexerWHITESPACE  = 10
)
