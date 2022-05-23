package hitPolicy

import (
	"fmt"
	"strings"
)

const (
	ErrDecisionTableUnknownHitPolicy = "unknown hit policy"
)

type HitPolicy string

// Single result decision tables
const (
	Unique   HitPolicy = "unique"   // Rules do not overlap. Only a single rule can match.
	First    HitPolicy = "first"    // Rules are evaluated from top to bottom. Rules may overlap, but only the first match counts.
	Priority HitPolicy = "priority" // Rule outputs are prioritized. Rules may overlap, but only the match with the highest output priority counts.
	Any      HitPolicy = "any"      // Multiple matching rules must not make a difference: all matching rules must lead to the same output.
)

// Multiple result decision tables
const (
	Collect     HitPolicy = "collect"     // All matching rules result in an arbitrarily ordered list of all the output entries.
	RuleOrder   HitPolicy = "ruleOrder"   // All matching rules result in a list of outputs ordered by the sequence of those rules in the decision table.
	OutputOrder HitPolicy = "outputOrder" // All matching rules result in a list of outputs ordered by their (decreasing) output priority.
)

func (hp *HitPolicy) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	switch s {
	case "unique":
		*hp = Unique
	case "first":
		*hp = First
	case "priority":
		*hp = Priority
	case "any":
		*hp = Any
	case "collect":
		*hp = Collect
	case "ruleOrder":
		*hp = RuleOrder
	case "outputOrder":
		*hp = OutputOrder

	default:
		return fmt.Errorf("%s \"%s\"", ErrDecisionTableUnknownHitPolicy, s)
	}

	return nil
}
