parser grammar ilang_parser;

options { tokenVocab=ilang_lexer; }

start: block EOF;

block: (
    expr SEMICOLON |
    statement |
    return
)+;

statement: (
    assignment SEMICOLON |
    ifStatement |
    whileLoop |
    forLoop |
    foreachLoop
);

expr: 
    not #notExpr |
    symbol #symbolExpr |
    stringLiteral #stringExpr |
    intLiteral #intExpr |
    floatLiteral #floatExpr |
    nullLiteral #nullExpr |
    booleanLiteral #booleanExpr |
    arrayLiteral #arrayExpr |
    mapLiteral #mapExpr |
    grouping #groupingExpr |
    functionDef #functionDefExpr |
    expr CONDITIONAL_OP expr #condition |
    expr ARITHMETIC_OP expr #arithmetic |
    expr BOOLEAN_OP expr #booleanAlgebra |
    expr  functionArgs #functionCall
;

functionDef: FUNC functionArgs scopeBody;

functionArgs: OPEN_PAREN ((args+=expr COMMA)* args+=expr)? CLOSE_PAREN;

assignment: LET? SYMBOL EQUALS expr;

ifStatement: IF conditionBody scopeBody elseifStatement* elseStatement?;

whileLoop: WHILE conditionBody scopeBody;

foreachLoop: FOR OPEN_PAREN SYMBOL IN expr CLOSE_PAREN scopeBody;

forLoop: FOR OPEN_PAREN init=assignment SEMICOLON cond=expr SEMICOLON step=assignment CLOSE_PAREN scopeBody;

return: RETURN expr SEMICOLON;

elseifStatement: ELSE IF conditionBody scopeBody;

elseStatement: ELSE scopeBody;

scopeBody: OPEN_BRACE block CLOSE_BRACE;

conditionBody: OPEN_PAREN expr CLOSE_PAREN;

not: NOT expr;

symbol: SYMBOL (DOT symbol)?;

stringLiteral: STRING;

intLiteral: INT;

floatLiteral: FLOAT;

nullLiteral: NULL;

booleanLiteral: TRUE | FALSE;

arrayLiteral: OPEN_BRACKET ((items+=expr COMMA)* items+=expr)? CLOSE_BRACKET;

mapLiteral: OPEN_BRACE ((items+=mapLiteralItem COMMA)* items+=mapLiteralItem)? CLOSE_BRACE;

mapLiteralItem:  mapKey COLON expr;

mapKey: (SYMBOL | OPEN_BRACKET expr CLOSE_BRACKET);

grouping: OPEN_PAREN expr CLOSE_PAREN;

