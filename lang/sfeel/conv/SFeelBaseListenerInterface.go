package conv

type SFeelBaseListenerInterface interface {
	ExitEmptyStatement(ctx string)
	ExitQualifiedName()
	ExitInteger()
	ExitFloat()
	ExitBoolean()
	ExitString()
	ExitDateTime()

	ExitInterval()
	ExitUnaryTest()
	ExitUnaryTests()
}
