parser grammar ilang_parser;

options { tokenVocab=ilang_lexer; }

start: block EOF;

block: (
    expr SEMICOLON |
    statement
)+;

statement: (
    let_assignment |
    if_statement |
    while_loop |
    for_loop |
    foreach_loop
);

expr: 
    not #not_expr |
    symbol #symbol_expr |
    string_literal #string_expr |
    int_literal #int_expr |
    float_literal #float_expr |
    null_literal #null_expr |
    boolean_literal #boolean_expr |
    grouping #grouping_expr |
    function_def #function_def_expr |
    expr CONDITIONAL_OP expr #condition |
    expr ARITHMETIC_OP expr #arithmetic |
    expr BOOLEAN_OP expr #boolean_algebra |
    expr  function_args #function_call
;

function_def: FUNC function_args scope_body;

function_args: OPEN_PAREN ((args+=expr COMMA)* args+=expr)? CLOSE_PAREN;

let_assignment: LET SYMBOL EQUALS expr SEMICOLON;

if_statement: IF condition_body scope_body elseif_statement* else_statement?;

while_loop: WHILE condition_body scope_body;

foreach_loop: FOR OPEN_PAREN SYMBOL IN expr CLOSE_PAREN scope_body;

for_loop: FOR OPEN_PAREN init=expr SEMICOLON cond=expr SEMICOLON step=expr CLOSE_PAREN scope_body;

elseif_statement: ELSE IF condition_body scope_body;

else_statement: ELSE scope_body;

scope_body: OPEN_BRACE block CLOSE_BRACE;

condition_body: OPEN_PAREN expr CLOSE_PAREN;

not: NOT expr;

symbol: SYMBOL (DOT symbol)?;

string_literal: STRING;

int_literal: INT;

float_literal: FLOAT;

null_literal: NULL;

boolean_literal: BOOLEAN_OP;

grouping: OPEN_PAREN expr CLOSE_PAREN;
