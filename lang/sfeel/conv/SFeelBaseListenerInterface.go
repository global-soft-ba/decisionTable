package conv

import sfeel "decisionTable/lang/sfeel/ast"

type SFeelBaseListenerInterface interface {
	ExitEmptyStatement(ctx sfeel.EmptyStatement)
	ExitQualifiedName(ctx sfeel.QualifiedName)
	ExitInteger(ctx sfeel.Integer)
	ExitFloat(ctx sfeel.Float)
	ExitBoolean(ctx sfeel.Boolean)
	ExitString(ctx sfeel.String)
	ExitDateTime(ctx sfeel.DateTime)
	ExitInterval(ctx sfeel.Interval)

	ExitUnaryTest(ctx sfeel.UnaryTest)
	ExitUnaryTests(ctx sfeel.UnaryTests)
}
