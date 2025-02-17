package calc

import (
	"fmt"
	"math"
	"strconv"

	"calculator/parsing"
)

type CalculatorListener struct {
	*parsing.BaseCalculatorParserListener
	nums []int64
}

func NewCalculatorListener() *CalculatorListener {
	return &CalculatorListener{
		BaseCalculatorParserListener: &parsing.BaseCalculatorParserListener{},
	}
}

func (l *CalculatorListener) pop() int64 {
	if len(l.nums) == 0 {
		panic("Invalid expression")
	}
	size := len(l.nums)
	val := l.nums[size-1]
	l.nums = l.nums[:size-1]
	return val
}

func (l *CalculatorListener) ExitProg(ctx *parsing.ProgContext) {
	if len(l.nums) != 1 {
		panic("Invalid expression")
	}
	fmt.Println(l.nums[0])
}

func (l *CalculatorListener) ExitNumber(ctx *parsing.NumberContext) {
	num, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err == nil {
		l.nums = append(l.nums, num)
	}
}

func (l *CalculatorListener) ExitExp(ctx *parsing.ExpContext) {
	right, left := l.pop(), l.pop()
	l.nums = append(l.nums, int64(math.Pow(float64(left), float64(right))))
}

func (l *CalculatorListener) ExitMulDivMod(ctx *parsing.MulDivModContext) {
	right, left := l.pop(), l.pop()
	switch ctx.GetOp().GetTokenType() {
	case parsing.CalculatorParserMUL:
		l.nums = append(l.nums, left*right)
	case parsing.CalculatorParserDIV:
		l.nums = append(l.nums, left/right)
	case parsing.CalculatorParserMOD:
		l.nums = append(l.nums, left%right)
	}
}

func (l *CalculatorListener) ExitAddSub(ctx *parsing.AddSubContext) {
	right, left := l.pop(), l.pop()
	switch ctx.GetOp().GetTokenType() {
	case parsing.CalculatorParserADD:
		l.nums = append(l.nums, left+right)
	case parsing.CalculatorParserSUB:
		l.nums = append(l.nums, left-right)
	}
}
