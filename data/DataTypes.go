package data

type HitPolicy string

// Single result decision tables
const (
	Unique   HitPolicy = "Unique"   // Rules do not overlap. Only a single rule can match.
	First    HitPolicy = "First"    // Rules are evaluated from top to bottom. Rules may overlap, but only the first match counts.
	Priority HitPolicy = "Priority" // Rule outputs are prioritized. Rules may overlap, but only the match with the highest output priority counts.
	Any      HitPolicy = "Any"      // Multiple matching rules must not make a difference: all matching rules must lead to the same output.
)

// Multiple result decision tables
const (
	Collect     HitPolicy = "Collect"      // All matching rules result in an arbitrarily ordered list of all the output entries.
	RuleOrder   HitPolicy = "Rule Order"   // All matching rules result in a list of outputs ordered by the sequence of those rules in the decision table.
	OutputOrder HitPolicy = "Output Order" // All matching rules result in a list of outputs ordered by their (decreasing) output priority.
)

type CollectOperator string

const (
	List  CollectOperator = "List Operator"
	Sum   CollectOperator = "Sum Operator"
	Min   CollectOperator = "Min Operator"
	Max   CollectOperator = "Max Operator"
	Count CollectOperator = "Count Operator" // TODO: Rename to "Number"?
)

type DataType string

const (
	Boolean  DataType = "boolean"
	Integer  DataType = "integer"
	Long     DataType = "long"
	Float    DataType = "float"
	Double   DataType = "double"
	String   DataType = "string"
	Date     DataType = "date"
	Time     DataType = "time"
	DateTime DataType = "datetime"
)

type ExpressionLanguage string

const (
	SFEEL ExpressionLanguage = "SFEEL"
	FEEL  ExpressionLanguage = "FEEL"
)

type Standard string

const (
	GRULE  Standard = "GRULE"
	DMN    Standard = "DMN (OMG)"
	DROOLS Standard = "DROOLS (REDHAT)"
)
