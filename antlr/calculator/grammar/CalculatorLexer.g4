lexer grammar CalculatorLexer;

ADD : '+';
SUB : '-';
MUL : '*';
DIV : '/';
MOD : '%';
EXP : '**';

LEFT_PAREN  : '(';
RIGHT_PAREN : ')';
NUMBER      : [1-9] [0-9]*;
WHITESPACE  : [ \t\r\n\u000C]+ -> skip;
