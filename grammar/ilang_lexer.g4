lexer grammar ilang;

fragment F_LETTER: [a-zA-z];
fragment F_NUMBER: [0-9];
fragment F_QUOTE: '"';
fragment F_UNDERSCORE: '_';
fragment F_DOT: '.';
fragment F_STRING_INSIDE: ('\\"' | ~('"' | '\n' | '\r'));


OPEN_PAREN: '(';
CLOSE_PAREN: ')';
OPEN_BRACE: '{';
CLOSE_BRACE: '}';
EQUALS: '=';
AND: '&&';
OR: '||';
NOT: '!';
QUESTION: '?';
SEMI_COLON: ';';
FUNC: 'func';
INT: F_NUMBER+;
FLOAT: (F_NUMBER+)? F_DOT F_NUMBER+;
STRING: F_QUOTE F_STRING_INSIDE* F_QUOTE;
COMMENT: '//' ~('\r' | '\n')* -> skip;
SYMBOL: (F_LETTER | F_UNDERSCORE) (F_LETTER | F_NUMBER | F_UNDERSCORE)*;
WHITE_SPACE: [ \t\r\n]+ -> skip;
