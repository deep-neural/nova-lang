


Source Code → Tokenizer → Parser → AST → Semantic Analysis → IR → Target Code

1: tokenizer
Preprocessor("#include <stdio.h>") [L2:C1]
Keyword("int") [L4:C1]
Identifier("main") [L4:C5]
Punctuator("(") [L4:C9]
Punctuator(")") [L4:C10]
Punctuator("{") [L4:C12]
Keyword("int") [L5:C5]
Identifier("x") [L5:C9]
Operator("=") [L5:C11]
Integer("0x42") [L5:C13]
Punctuator(";") [L5:C17]
Keyword("float") [L6:C5]
Identifier("y") [L6:C11]
Operator("=") [L6:C13]
Float("3.14") [L6:C15]
Punctuator(";") [L6:C19]
Keyword("if") [L7:C5]


2: Parsing: Build an Abstract Syntax Tree (AST)
- hierarchical tree structure (AST) based on your language's grammar rules

    +
   / \
  3   *
     / \
    4   2

- The parser enforces your language’s grammar. If tokens appear in an invalid order, the parser throws errors.

 

Fast parser builds AST with position info


if x > 5:
    print("ok")


lexer (or tokenizer) processes the raw source code character by character and groups sequences of characters into meaningful tokens.


IF, IDENTIFIER(x), GT, INT(5), COLON, NEWLINE, INDENT, IDENTIFIER(print), LPAREN, STRING("Hello"), RPAREN