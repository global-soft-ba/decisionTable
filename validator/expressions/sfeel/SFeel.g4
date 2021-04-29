grammar SFeel;

// Datatypes
SIGN: '+'|'-';
INTEGER:  '0' | SIGN? '1'..'9''0'..'9'*;
FLOAT:    '0'.'0' | SIGN?  '0'..'9'+ '.' '0'..'9'*'1'..'9';

STRING: '"' NONCONTROL_CHAR* '"';
    fragment NONCONTROL_CHAR: LETTER | DIGIT | SYMBOL | SPACE;
    fragment LETTER: LOWER | UPPER;
    fragment LOWER: 'a'..'z'|'ö'|'ä'|'ü';
    fragment UPPER: 'A'..'Z'|'Ö'|'Ä'|'Ü';
    fragment DIGIT: '0'..'9';
    fragment SPACE: ' ' | '\t';
    fragment SYMBOL: '!' | '#'..'/' | ':'..'@' | '['..'`' | '{'..'~';

BOOL: TRUE | FALSE ;
    fragment TRUE: 'true';
    fragment FALSE: 'false';

//Comparisions Operator
LESS: '<';
LESSEQ :'<=';
GREATER: '>';
GREATEREQ:'>=';

// Ranges
RANGEIN:'[';
RANGEOUT:']';

// Functions
DISJUNCTION: ',';
NEGATION: 'not';


DATEANDTIME: 'DateAndTime("'FORMAT'")' ;
    FORMAT: YYYY'-'MM'-'DD'T'HH'-'MMM'-'SS;
    fragment YYYY: '0'..'9''0'..'9''0'..'9''1'..'9' | '0'..'9''0'..'9''1'..'9''0'..'9' | '0'..'9''1'..'9''0'..'9''0'..'9' | '1'..'9''0'..'9''0'..'9''0'..'9';
    fragment MM:  '0''1'..'9' | '1''0'..'2';
    fragment DD:  '0''1'..'9' | '1'..'2''0'..'9' | '3''0'..'1';
    fragment HH:  '0''0'..'9' | '1''0'..'9' | '2''0'..'3';
    fragment MMM: '0''0'..'9' | '1'..'5''0'..'9';
    fragment SS:  '0''0'..'9' | '1'..'5''0'..'9';

start: expression EOF;

number: (INTEGER|FLOAT);
strings: STRING;
bools: BOOL;
datetime: DATEANDTIME;

equalcomparison : (bools|strings|datetime|number) ;

comparison: comparisonnumber | comparisondatetime;
    op:(LESS | LESSEQ | GREATER | GREATEREQ);
    comparisonnumber:   op number;
    comparisondatetime: op datetime;

ranges: rangenumber | rangedatetime;
    rop: (RANGEIN | RANGEOUT);
    rangenumber: rop number '..' number rop;
    rangedatetime: rop (datetime '..' datetime) rop;


disjunctions : (disjunctionsNumber|disjunctionsString|disjunctionsDateTime);
    disjunctionsNumber
        : disjunctionsNumber DISJUNCTION disjunctionsNumber
        | (number | comparisonnumber | rangenumber)
        ;
    disjunctionsString
        : disjunctionsString DISJUNCTION disjunctionsString
        | strings
        ;
    disjunctionsDateTime
        : disjunctionsDateTime DISJUNCTION disjunctionsDateTime
        | (datetime | comparisondatetime | rangedatetime)
        ;

negation
    : NEGATION'('(equalcomparison|comparison|ranges|disjunctions)')'
    ;

expression
    : equalcomparison                               # EqualcomparisonRule
    | comparison                                    # ComparisionsRule
    | ranges                                         # RangeRule
    | disjunctions                                  # DisjunctionRule
    | negation                                      # NegationRule
    | '-'                                           # EmptyRule
    ;

