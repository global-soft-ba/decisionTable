package model

type DTableStandard string

const (
	GRULE DTableStandard = "GRULE"
	DMN   DTableStandard = "DMN (OMG)"
)

type HitPolicy string

const (
	Unique      HitPolicy = "Unique"
	First       HitPolicy = "First"
	Priority    HitPolicy = "Priority"
	Any         HitPolicy = "Any"
	RuleOrder   HitPolicy = "Rule Order"
	OutputOrder HitPolicy = "Output Order"
	Collect     HitPolicy = "Collect"
)

type ExpressionLanguage string

const (
	GRL        ExpressionLanguage = "GRL"
	FEEL       ExpressionLanguage = "Feel"
	Javascript ExpressionLanguage = "javascript"
	Python     ExpressionLanguage = "python"
	Groovy     ExpressionLanguage = "groovy"
	JRuby      ExpressionLanguage = "jruby"
	Juel       ExpressionLanguage = "juel"
)

type CollectOperator string

const (
	List  CollectOperator = "List Operator"
	Sum   CollectOperator = "Sum Operator"
	Min   CollectOperator = "Min Operator"
	Max   CollectOperator = "Max Operator"
	Count CollectOperator = "Count Operator"
)

type VariableTyp string

const (
	String  VariableTyp = "string"
	Boolean VariableTyp = "boolean"
	Integer VariableTyp = "integer"
	Float   VariableTyp = "float"
	Long    VariableTyp = "long"
	Double  VariableTyp = "double"
	Date    VariableTyp = "date"
)
