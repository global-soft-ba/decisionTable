package conv

type SFeelBaseListener struct{}

func (l *SFeelBaseListener) ExitEmptyStatement(ctx string) {}
func (l *SFeelBaseListener) ExitQualifiedName()            {}

func (l *SFeelBaseListener) ExitInteger()  {}
func (l *SFeelBaseListener) ExitFloat()    {}
func (l *SFeelBaseListener) ExitBoolean()  {}
func (l *SFeelBaseListener) ExitString()   {}
func (l *SFeelBaseListener) ExitDateTime() {}

func (l *SFeelBaseListener) ExitInterval()   {}
func (l *SFeelBaseListener) ExitUnaryTest()  {}
func (l *SFeelBaseListener) ExitUnaryTests() {}
