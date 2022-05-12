package types

import (
	"github.com/global-soft-ba/decisionTable/data/collectOperator"
	"github.com/global-soft-ba/decisionTable/data/dataType"
	"github.com/global-soft-ba/decisionTable/data/expressionLanguage"
	"github.com/global-soft-ba/decisionTable/data/hitPolicy"
	"github.com/global-soft-ba/decisionTable/data/standard"
)

type DecisionTableStandard struct {
	ExpressionLanguages map[expressionLanguage.ExpressionLanguage]bool
	HitPolicies         map[hitPolicy.HitPolicy]bool
	CollectOperators    map[collectOperator.CollectOperator]bool
	VariableTypes       map[dataType.DataType]bool
}

var DecisionTableStandards = map[standard.Standard]DecisionTableStandard{
	standard.GRULE: grule,
	standard.DMN:   dmn,
}

var dmn = DecisionTableStandard{
	ExpressionLanguages: map[expressionLanguage.ExpressionLanguage]bool{
		expressionLanguage.FEEL: true,
	},

	HitPolicies: map[hitPolicy.HitPolicy]bool{
		hitPolicy.Unique:      true,
		hitPolicy.First:       true,
		hitPolicy.Priority:    true,
		hitPolicy.Any:         true,
		hitPolicy.Collect:     true,
		hitPolicy.RuleOrder:   true,
		hitPolicy.OutputOrder: true,
	},

	CollectOperators: map[collectOperator.CollectOperator]bool{
		collectOperator.List:  true,
		collectOperator.Sum:   true,
		collectOperator.Min:   true,
		collectOperator.Max:   true,
		collectOperator.Count: true,
	},

	VariableTypes: map[dataType.DataType]bool{
		dataType.Boolean:  true,
		dataType.Integer:  true,
		dataType.Long:     true,
		dataType.Double:   true,
		dataType.String:   true,
		dataType.DateTime: true,
	},
}

var grule = DecisionTableStandard{
	ExpressionLanguages: map[expressionLanguage.ExpressionLanguage]bool{
		expressionLanguage.SFEEL: true,
	},

	HitPolicies: map[hitPolicy.HitPolicy]bool{
		hitPolicy.Unique:   true,
		hitPolicy.First:    true,
		hitPolicy.Priority: true,
	},

	CollectOperators: map[collectOperator.CollectOperator]bool{},

	VariableTypes: map[dataType.DataType]bool{
		dataType.Boolean:  true,
		dataType.Integer:  true,
		dataType.Float:    true, // In GRL, Float is REAL
		dataType.String:   true,
		dataType.DateTime: true,
	},
}
