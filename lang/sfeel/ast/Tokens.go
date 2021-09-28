package ast

const (
	SFeelAssignmentEqual = iota - 1
	SFeelOperatorPOW
	SFeelOperatorMUL
	SFeelOperatorDIV
	SFeelOperatorADD
	SFeelOperatorSUB
	SFeelOperatorLESS
	SFeelOperatorLESSEQ
	SFeelOperatorGREATER
	SFeelOperatorGREATEREQ
	SFeelOperatorEQUAL
	SFeelOperatorNOTEQUAL
	SFeelClosedIntervalSTART
	SFeelClosedIntervalEND
	SFeelOpenIntervalSTART
	SFeelOpenIntervalEND

	SFeelSeparatorQualifiedName = "."
)

// SFeel Tokens for converter
type Token struct {
	Type    int
	Literal string
}
