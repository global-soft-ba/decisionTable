package ast

type SFeelListener struct{}

func (l *SFeelListener) ExitEmptyStatement(ctx EmptyStatement) {}
func (l *SFeelListener) ExitQualifiedName(ctx QualifiedName)   {}
func (l *SFeelListener) ExitInteger(ctx Integer)               {}
func (l *SFeelListener) ExitFloat(ctx Float)                   {}
func (l *SFeelListener) ExitBoolean(ctx Boolean)               {}
func (l *SFeelListener) ExitString(ctx String)                 {}
func (l *SFeelListener) ExitDateTime(ctx DateTime)             {}

func (l *SFeelListener) ExitInterval(ctx Interval)     {}
func (l *SFeelListener) ExitUnaryTest(ctx UnaryTest)   {}
func (l *SFeelListener) ExitUnaryTests(ctx UnaryTests) {}
