package decisionTable

type RuleBuilderInterface interface {
	SetAnnotation(annotation string) RuleBuilderInterface
	AddInputEntry(expression string) RuleBuilderInterface
	AddOutputEntry(expression string) RuleBuilderInterface
	Build() Rule
}
