package main

import (
	conv "decisionTable/converters"
	"decisionTable/model"
	"fmt"
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
}

type Employee struct {
	ResponsibleEmployee string
	FourEyesPrinciple   bool
}

func main() {

	table, _ := CreateDecisionTable().
		SetName("Determine Employee").
		SetDefinitionKey("determineEmployee").
		SetNotationStandard(model.GRULE).
		SetHitPolicy(model.Unique).
		AddInputField("TypeOfClaim", "Claim", model.String).
		AddInputField("ExpenditureOfClaim", "Claim", model.Integer).
		AddOutputField("ResponsibleEmployee", "Employee", model.String).
		AddOutputField("FourEyesPrinciple", "Employee", model.Boolean).
		AddRule("R1").
		AddInputEntry(`"Car Accident"`, model.SFEEL).
		AddInputEntry("<1000", model.SFEEL).
		AddOutputEntry(`"Müller"`, model.SFEEL).
		AddOutputEntry("false", model.SFEEL).
		BuildRule().
		AddRule("R2").
		AddInputEntry(`"Car Accident"`, model.SFEEL).
		AddInputEntry("[1000..10000]", model.SFEEL).
		AddOutputEntry(`"Schulz"`, model.SFEEL).
		AddOutputEntry("false", model.SFEEL).
		BuildRule().
		AddRule("R3").
		AddInputEntry("-", model.SFEEL).
		AddInputEntry(">=10000", model.SFEEL).
		AddOutputEntry("-", model.SFEEL).
		AddOutputEntry("true", model.SFEEL).
		BuildRule().
		Build()

	// COnvert Table Into Grule Rules
	converter, _ := conv.CreateTableConverterFactory().GetTableConverter(model.GRULE, model.GRL)
	rules, err := table.Convert(converter)
	if err != nil {
		fmt.Print("Error:", err)
	}

	//Load Library and Insert rules
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
	claim := Claim{
		TypeOfClaim:        "Car Accident",
		ExpenditureOfClaim: 100,
	}
	employee := Employee{}

	// Load example
	dataCtx := ast.NewDataContext()
	err = dataCtx.Add("Claim", &claim)
	err = dataCtx.Add("Employee", &employee)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// CreateEngine Instance
	ruleEngine := engine.NewGruleEngine()
	kb := kl.Library.NewKnowledgeBaseInstance("#exampleBase", "0.0.1")

	//Execution
	//ruleEntries, _ := ruleEngine.FetchMatchingRules(dataCtx,kb)
	now := time.Now()
	err = ruleEngine.Execute(dataCtx, kb)
	fmt.Println("--------------OutCome------------------------")
	fmt.Println("time elapse:", time.Since(now))

	fmt.Println("Input Claim =", claim)
	fmt.Println("Responsible Employee =", employee)
}
