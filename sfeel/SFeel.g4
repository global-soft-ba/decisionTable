grammar SFeel;
//OMG Version 1.3


start
    : simple_unary_tests EOF // Each input entry SHALL be an instance of simple unary tests (grammar rule 12 - OMG Standard)
    | simple_expressions EOF; // Each output entry SHALL be a simple expression (grammar rule 3 - OMG Standard).


// Unaray Tests
simple_unary_tests   //Disjunctions
    : simple_positive_unary_tests                       #SimpleUnaryTests
    | 'not(' simple_positive_unary_tests ')'            #NegationSimpleUnaryTests
    | '-'                                               #EmptySimpleUnaryTests
    ;

simple_positive_unary_tests: simple_positive_unary_test ( ',' simple_positive_unary_test )* ;
simple_positive_unary_test: unary_comparison | interval ;
unary_comparison
    : (LESS | LESSEQ | GREATER | GREATEREQ) endpoint #UnaryComparison
    | endpoint                                       #EqualUnaryComparison
    ;

// Unary Interval Statements
interval: ( open_interval_start | closed_interval_start ) endpoint '..' endpoint ( open_interval_end | closed_interval_end ) ;
open_interval_start: '(' | ']' ;
closed_interval_start: '[' ;
open_interval_end: ')' | '[' ;
closed_interval_end: ']' ;

// Simple Expressions
simple_expressions: simple_expression  (',' simple_expression)* ; //Disjunctions
expression: simple_expression ;
simple_expression: arithmetic_expression | simple_value | comparison ;

// Comparison Operations
comparison
        : simple_value (LESS | LESSEQ | GREATER | GREATEREQ | EQUAL | NOTEQUAL) expression
        | arithmetic_expression (LESS | LESSEQ | GREATER | GREATEREQ | EQUAL | NOTEQUAL) expression
        | comparison (LESS | LESSEQ | GREATER | GREATEREQ | EQUAL | NOTEQUAL) expression
        ;

// Arithmetic Expressions
arithmetic_expression
   : simple_value                                                    # Value
   | '(' arithmetic_expression ')'                                   # Parentheses
   | '-' arithmetic_expression                                       # ArithmeticNegation
   | arithmetic_expression operator=POW arithmetic_expression        # Power
   | arithmetic_expression operator=(MUL|DIV) arithmetic_expression  # MultiplicationOrDivision
   | arithmetic_expression operator=(ADD|SUB) arithmetic_expression  # AdditionOrSubtraction
   ;




// Simple Primitives
endpoint: simple_value ;
simple_value: qualified_name | simple_literal ;


// Qualified Names (respects GOLANG Identifier unicode_letter | "_"  -- https://golang.org/ref/spec#letter)
qualified_name: Name ('.' Name)* ;
Name: Name_start_char (Name_part_char)*;
    fragment Name_start_char: LOWER | UPPER;
    fragment Name_part_char: Name_start_char | '_' | DIGIT;



// Simple Literals
simple_literal: numeric_literal | string_literal | boolean_literal | date_time_literal ;
date_time_literal: ('date' | 'time' | 'date and time' | 'duration' ) '(' string_literal ')' ;
numeric_literal: '-'? (integer_literal | real_literal);
integer_literal: INTEGER;
real_literal: REAL;
boolean_literal: 'true' | 'false' ;
string_literal: STRING+;

// Arithmetic Expression
POW: '**';
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';


// Comparison Operators
LESS: '<';
LESSEQ :'<=';
GREATER: '>';
GREATEREQ:'>=';
EQUAL:'=';
NOTEQUAL: '!=';

// General Datatypes
INTEGER:  '0'..'9'DIGIT*;
REAL:     DIGIT+ '.' DIGIT*'1'..'9';


STRING: '"' NONCONTROL_CHAR* '"';
    fragment NONCONTROL_CHAR: LETTER | DIGIT | SYMBOL | SPACE;
    fragment LETTER: LOWER | UPPER;
    fragment LOWER: 'a'..'z'|'ö'|'ä'|'ü';
    fragment UPPER: 'A'..'Z'|'Ö'|'Ä'|'Ü';
    fragment DIGIT: '0'..'9';
    fragment SPACE: ' ' | '\t';
    fragment SYMBOL: '!' | '#'..'/' | ':'..'@' | '['..'`' | '{'..'~';

WS: [ \t\r\n] -> skip ;

