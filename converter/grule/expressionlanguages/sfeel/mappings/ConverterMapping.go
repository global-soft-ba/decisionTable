package mappings

const (
	// GRULE Specific Tokens
	// Logical operator token
	EQUAL    = 1
	AND      = 2
	OR       = 3
	NEGATION = 4
)

type ConverterMapping struct {
	TargetToken         map[int]string
	ComparisonOperators map[int]string
	Operations          map[int]string
	StartRanges         map[int]string
	EndRanges           map[int]string
}
