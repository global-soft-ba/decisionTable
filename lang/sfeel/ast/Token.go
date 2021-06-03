package ast

const (
	SFeelOperatorPOW = iota
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
)

// SFeel Tokens for converter
type Token struct {
	Type    int
	Literal string
}
