package data

type DTableStandard string

const (
	GRULE  DTableStandard = "GRULE"
	DMN    DTableStandard = "DMN (OMG)"
	DROOLS DTableStandard = "DROOLS (REDHAT)"
)

type HitPolicy string

const (
	Unique      HitPolicy = "Unique"   // Rules do not overlap. Only a single rule can match.
	First       HitPolicy = "First"    // Rules are evaluated from top to bottom. Rules may overlap, but only the first match counts.
	Priority    HitPolicy = "Priority" // Rule outputs are prioritized. Rules may overlap, but only the match with the highest output priority counts.
	Any         HitPolicy = "Any"      // Multiple matching rules must not make a difference: all matching rules must lead to the same output.
	RuleOrder   HitPolicy = "Rule Order"
	OutputOrder HitPolicy = "Output Order"
	Collect     HitPolicy = "Collect" // The output of all matching rules is aggregated by means of an operator:
)

type ExpressionLanguage string

const (
	SFEEL      ExpressionLanguage = "sFeel"
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

type DataTyp string

const (
	String   DataTyp = "string"
	Boolean  DataTyp = "boolean"
	Integer  DataTyp = "integer"
	Float    DataTyp = "float"
	Long     DataTyp = "long"
	Double   DataTyp = "double"
	DateTime DataTyp = "date"
)
