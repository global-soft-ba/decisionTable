grammar SFeel;

// Validator EntryPoints

// Complete Language Grammar
entry: expression EOF;

// Datatype Language Grammar
validIntegerInput
    : equalcomparisonInteger EOF
    | comparisonInteger EOF
    | rangeInteger EOF
    | disjunctionsInteger EOF
    | NEGATION'('(equalcomparisonInteger|comparisonInteger|rangeInteger|disjunctionsInteger)')' EOF
    | '-' EOF
    ;

validNumberInput
    : equalcomparisonNumber EOF
    | comparisonNumber EOF
    | rangeNumber EOF
    | disjunctionsNumber EOF
    | NEGATION'('(equalcomparisonNumber|comparisonNumber|rangeNumber|disjunctionsNumber)')' EOF
    | '-' EOF
    ;

validStringInput
    : equalcomparisonStrings EOF
    | disjunctionsString EOF
    | NEGATION'('(equalcomparisonStrings|disjunctionsString)')' EOF
    | '-' EOF
    ;

validBoolInput
    : equalcomparisonBool EOF
    | '-' EOF
    ;

validDateTimeInput
    : equalcomparisonDateTime EOF
    | comparisonDateTime EOF
    | rangeDateTime EOF
    | disjunctionsDateTime EOF
    | NEGATION'('(equalcomparisonDateTime|comparisonDateTime|rangeDateTime|disjunctionsDateTime)')' EOF
    | '-' EOF
    ;

validIntegerOutput
    : INTEGER EOF
    | '-' EOF
    ;

validNumberOutput
    : number EOF
    | '-' EOF
    ;

validStringOutput
    : strings EOF
    | '-' EOF
    ;

validBoolOutput
    : bools EOF
    | '-' EOF
    ;

validDateTimeOutput
    : datetime EOF
    | '-' EOF
    ;

// Entire Language Grammar SFeel

expression
    : equalcomparison                               # EqualcomparisonRule
    | comparison                                    # ComparisionsRule
    | ranges                                        # RangeRule
    | disjunctions                                  # DisjunctionRule
    | negation                                      # NegationRule
    | '-'                                           # EmptyInputRule
    ;

// General Rules

number: (INTEGER|FLOAT);
strings: STRING;
bools: BOOL;
datetime: DATEANDTIME;

equalcomparison : (equalcomparisonNumber|equalcomparisonStrings|equalcomparisonDateTime|equalcomparisonBool) ;
    equalcomparisonInteger : INTEGER;
    equalcomparisonNumber: number;
    equalcomparisonBool: bools;
    equalcomparisonStrings: strings;
    equalcomparisonDateTime: datetime;

comparison:  (comparisonDateTime | comparisonInteger| comparisonNumber);
    comparisonOps:(LESS | LESSEQ | GREATER | GREATEREQ);
    comparisonInteger:   comparisonOps INTEGER;
    comparisonDateTime: comparisonOps datetime;
    comparisonNumber:   comparisonOps number;

ranges: (rangeInteger | rangeDateTime | rangeNumber);
    rop: (RANGEIN | RANGEOUT);
    rangeInteger: rop INTEGER '..' INTEGER rop;
    rangeNumber: rop number '..' number rop;
    rangeDateTime: rop (datetime '..' datetime) rop;

disjunctions : (disjunctionsInteger|disjunctionsNumber|disjunctionsString|disjunctionsDateTime);
    disjunctionsInteger
        : disjunctionsInteger DISJUNCTION disjunctionsInteger
        | (INTEGER | comparisonInteger | rangeInteger)
        ;
    disjunctionsNumber
        : disjunctionsNumber DISJUNCTION disjunctionsNumber
        | (number | comparisonNumber | rangeNumber)
        ;
    disjunctionsString
        : disjunctionsString DISJUNCTION disjunctionsString
        | strings
        ;
    disjunctionsDateTime
        : disjunctionsDateTime DISJUNCTION disjunctionsDateTime
        | (datetime | comparisonDateTime | rangeDateTime)
        ;

negation
    : NEGATION'('(equalcomparison|comparison|ranges|disjunctions)')'
    ;


// Lexer and Token Config

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

// Comparision Operators
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
    FORMAT: YYYY'-'MM'-'DD'T'HH':'MMM':'SS;
    fragment YYYY: '0'..'9''0'..'9''0'..'9''1'..'9' | '0'..'9''0'..'9''1'..'9''0'..'9' | '0'..'9''1'..'9''0'..'9''0'..'9' | '1'..'9''0'..'9''0'..'9''0'..'9';
    fragment MM:  '0''1'..'9' | '1''0'..'2';
    fragment DD:  '0''1'..'9' | '1'..'2''0'..'9' | '3''0'..'1';
    fragment HH:  '0''0'..'9' | '1''0'..'9' | '2''0'..'3';
    fragment MMM: '0''0'..'9' | '1'..'5''0'..'9';
    fragment SS:  '0''0'..'9' | '1'..'5''0'..'9';
