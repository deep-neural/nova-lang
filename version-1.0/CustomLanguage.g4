grammar CustomLanguage;

program: function* EOF;

function
    : FUNC ID LPAREN parameters? RPAREN ARROW type block
    ;

parameters
    : parameter (COMMA parameter)*
    ;

parameter
    : ID COLON type
    ;

type
    : INT
    | FLOAT
    | STRING
    | BOOL
    ;

block
    : LBRACE statement* RBRACE
    ;

statement
    : variableDecl
    | assignment
    | functionCall SEMI
    | returnStmt
    ;

variableDecl
    : (type | VAR) ID ASSIGN expr SEMI
    ;

assignment
    : ID ASSIGN expr SEMI
    ;

functionCall
    : ID LPAREN (expr (COMMA expr)*)? RPAREN
    ;

returnStmt
    : RETURN expr SEMI
    ;

expr
    : expr (MULT | DIV) expr       # MulDivExpr
    | expr (PLUS | MINUS) expr     # AddSubExpr
    | functionCall                 # FunctionCallExpr
    | ID                          # VariableExpr
    | literal                     # LiteralExpr
    | LPAREN expr RPAREN          # ParenExpr
    ;

literal
    : INT_LITERAL
    | FLOAT_LITERAL
    | STRING_LITERAL
    | BOOL_LITERAL
    ;

// Lexer Rules

FUNC: 'func';
INT: 'int';
FLOAT: 'float';
STRING: 'string';
BOOL: 'bool';
RETURN: 'return';
VAR: 'var';

BOOL_LITERAL: 'true' | 'false';

ARROW: '->';
ASSIGN: '=';
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
COMMA: ',';
SEMI: ';';
COLON: ':';
PLUS: '+';
MINUS: '-';
MULT: '*';
DIV: '/';

STRING_LITERAL: '"' ( ESC | ~["\\] )* '"';
fragment ESC: '\\' [btnfr"'\\];

FLOAT_LITERAL: [0-9]+ '.' [0-9]* | '.' [0-9]+;
INT_LITERAL: [0-9]+;

ID: [a-zA-Z_] [a-zA-Z0-9_]*;

WS: [ \t\r\n]+ -> skip;

LINE_COMMENT: '//' ~[\r\n]* -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;