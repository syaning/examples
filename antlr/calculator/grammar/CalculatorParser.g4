parser grammar CalculatorParser;

options {
    tokenVocab = CalculatorLexer;
}

prog: expr EOF;

expr: expr op=EXP expr            # Exp
    | expr op=(MUL|DIV|MOD) expr  # MulDivMod
    | expr op=(ADD|SUB) expr      # AddSub
    | NUMBER                      # Number
    | LEFT_PAREN expr RIGHT_PAREN # Parens
    ;
