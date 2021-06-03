package conv

import (
	sfeel "decisionTable/lang/sfeel/ast"
)

type SFeelBaseListener struct{}

func (l *SFeelBaseListener) ExitEmptyStatement(ctx sfeel.EmptyStatement) {}
func (l *SFeelBaseListener) ExitQualifiedName(ctx sfeel.QualifiedName)   {}
func (l *SFeelBaseListener) ExitInteger(ctx sfeel.Integer)               {}
func (l *SFeelBaseListener) ExitFloat(ctx sfeel.Float)                   {}
func (l *SFeelBaseListener) ExitBoolean(ctx sfeel.Boolean)               {}
func (l *SFeelBaseListener) ExitString(ctx sfeel.String)                 {}
func (l *SFeelBaseListener) ExitDateTime(ctx sfeel.DateTime)             {}

func (l *SFeelBaseListener) ExitInterval(ctx sfeel.Interval)     {}
func (l *SFeelBaseListener) ExitUnaryTest(ctx sfeel.UnaryTest)   {}
func (l *SFeelBaseListener) ExitUnaryTests(ctx sfeel.UnaryTests) {}
