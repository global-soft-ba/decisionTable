grammar SFeel;

// Validator EntryPoints

// Complete Language Grammar
entry: expression EOF;

// Datatype Language Grammar
validIntegerInput
    : equalcomparisonInteger EOF    # EqualComparisonIntegerInputRule
    | comparisonInteger EOF         # ComparisonIntegerInputRule
    | rangeInteger EOF              # RangeComparisonIntegerInputRule
    | disjunctionsInteger EOF       # DisjunctionsIntegerInputRule
    | NEGATION'('(equalcomparisonInteger|comparisonInteger|rangeInteger|disjunctionsInteger)')' EOF  # NegationIntegerInputRule
    | '-' EOF                       # EmptyIntegerInputRule
    ;

validNumberInput
    : equalcomparisonNumber EOF     # EqualComparisonNumberInputRule
    | comparisonNumber EOF          # ComparisonNumberInputRule
    | rangeNumber EOF               # RangeComparisonNumberInputRule
    | disjunctionsNumber EOF        # DisjunctionsNumberInputRule
    | NEGATION'('(equalcomparisonNumber|comparisonNumber|rangeNumber|disjunctionsNumber)')' EOF # NegationNumberInputRule
    | '-' EOF                       # EmptyNumberInputRule
    ;

validStringInput
    : equalcomparisonStrings EOF    # EqualComparisonStringInputRule
    | disjunctionsString EOF        # ComparisonStringInputRule
    | NEGATION'('(equalcomparisonStrings|disjunctionsString)')' EOF # NegationStringInputRule
    | '-' EOF                       # EmptyStringInputRule
    ;

validBoolInput
    : equalcomparisonBool EOF       # EqualComparisonBoolInputRule
    | '-' EOF                       # EmptyBoolInputRule
    ;

validDateTimeInput
    : equalcomparisonDateTime EOF   # EqualComparisonDateTimeInputRule
    | comparisonDateTime EOF        # ComparisonDateTimeInputRule
    | rangeDateTime EOF             # RangeComparisonDateTimeInputRule
    | disjunctionsDateTime EOF      # DisjunctionsDateTimeInputRule
    | NEGATION'('(equalcomparisonDateTime|comparisonDateTime|rangeDateTime|disjunctionsDateTime)')' EOF # NegationDateTimeInputRule
    | '-' EOF                       # EmptyDateTimeInputRule
    ;

validIntegerOutput
    : INTEGER EOF                   # IntegerOutputRule
    | '-' EOF                       # EmptyIntegerOutputRule
    ;

validNumberOutput
    : number EOF                    # NumberOutputRule
    | '-' EOF                       # EmptyNumberOutputRule
    ;

validStringOutput
    : strings EOF                   # StringOutputRule
    | '-' EOF                       # EmptyStringOutputRule
    ;

validBoolOutput
    : bools EOF                     # BoolOutputRule
    | '-' EOF                       # EmptyBoolOutputRule
    ;

validDateTimeOutput
    : datetime EOF                  # DateTimeOutputRule
    | '-' EOF                       # EmptyDateTimeOutputRule
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

//TODO IN GRL Integers with different Semantic http://hyperjumptech.viewdocs.io/grule-rule-engine/GRL_Literals_en/

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
