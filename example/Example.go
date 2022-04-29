package main

import (
	"fmt"
	"github.com/global-soft-ba/decisionTable"
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"time"
)

type KnowledgeLib struct {
	Library *ast.KnowledgeLibrary
	Builder *builder.RuleBuilder
}

func CreateKnowledgeLibrary() *KnowledgeLib {
	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

	return &KnowledgeLib{knowledgeLibrary, ruleBuilder}
}

func (rb *KnowledgeLib) AddRule(rule string, knowledgeBase string) error {
	// Add the rule definition above into the library and name it 'TutorialRules'  version '0.0.1'
	bs := pkg.NewBytesResource([]byte(rule))
	err := rb.Builder.BuildRuleFromResource(knowledgeBase, "0.0.1", bs)
	if err != nil {
		return err
	}
	return nil
}

type Claim struct {
	TypeOfClaim        string
	ExpenditureOfClaim int
	TimeOfClaim        time.Time
}

type Employee struct {
	ResponsibleEmployee string
	FourEyesPrinciple   bool
	LastTime            time.Time
}

func main() {

	table, _ := decisionTable.NewDecisionTableBuilder().
		SetID("determineEmployee").
		SetName("Determine Employee").
		SetHitPolicy(data.Unique).
		SetExpressionLanguage(data.SFEEL).
		SetStandard(data.GRULE).
		AddInputField(data.TestField{Name: "Claim", Key: "TypeOfClaim", Type: data.String}).
		AddInputField(data.TestField{Name: "Claim", Key: "ExpenditureOfClaim", Type: data.Integer}).
		AddOutputField(data.TestField{Name: "Employee", Key: "ResponsibleEmployee", Type: data.String}).
		AddOutputField(data.TestField{Name: "Employee", Key: "FourEyesPrinciple", Type: data.Boolean}).
		AddRule(decisionTable.NewRuleBuilder().SetAnnotation("R1").
			AddInputEntry(`"Car Accident"`).
			AddInputEntry("<1000").
			AddOutputEntry(`"Müller"`).
			AddOutputEntry("false").
			Build(),
		).
		AddRule(decisionTable.NewRuleBuilder().SetAnnotation("R2").
			AddInputEntry(`"Car Accident"`).
			AddInputEntry("[1000..10000]").
			AddOutputEntry(`"Schulz"`).
			AddOutputEntry("false").
			Build(),
		).
		AddRule(decisionTable.NewRuleBuilder().SetAnnotation("R3").
			AddInputEntry("-").
			AddInputEntry(">=10000").
			AddOutputEntry("-").
			AddOutputEntry("true").
			Build(),
		).
		Build()

	// ConvertToGrlAst Table Into Grule Rules
	rules, err := table.Convert(data.GRULE)
	if err != nil {
		fmt.Print("Error:", err)
	}

	//Load Library and Insert rules
	fmt.Println("--------------GRL-RUlES------------------------")
	kl := CreateKnowledgeLibrary()
	result := rules.([]string)
	for _, rule := range result {
		fmt.Print(rule)
		addErr := kl.AddRule(rule, "#exampleBase")
		if addErr != nil {
			fmt.Print("Error:", addErr)
		}
	}

	// Create Example Data
	timeVal, err := time.Parse("2006-01-02T15:04:05", "2021-01-04T12:00:00")
	claim := Claim{
		TypeOfClaim:        "Car Accident",
		ExpenditureOfClaim: 100,
		TimeOfClaim:        timeVal,
	}
	employee := Employee{}

	// CreateEngine Instance
	ruleEngine := engine.NewGruleEngine()
	kb := kl.Library.NewKnowledgeBaseInstance("#exampleBase", "0.0.1")

	now := time.Now()
	// Load example
	dataCtx := ast.NewDataContext()
	err = dataCtx.Add("Claim", &claim)
	err = dataCtx.Add("Employee", &employee)
	if err != nil {
		fmt.Println("Error:", err)
	}

	//Execution
	err = ruleEngine.Execute(dataCtx, kb)
	fmt.Println("--------------OutCome------------------------")
	fmt.Println("time elapse:", time.Since(now))

	fmt.Println("Input Claim =", claim)
	fmt.Println("Responsible Employee =", employee)
}
