package decisionTable

import (
	"fmt"
	conv "github.com/global-soft-ba/decisionTable/converters"
	"github.com/global-soft-ba/decisionTable/model"
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

	table, _ := CreateDecisionTable().
		SetName("Determine Employee").
		SetDefinitionKey("determineEmployee").
		SetNotationStandard(model.GRULE).
		SetHitPolicy(model.Unique).
		AddInputField("TypeOfClaim", "Claim", model.String).
		AddInputField("ExpenditureOfClaim", "Claim", model.Integer).
		AddInputField("TimeOfClaim", "Claim", model.DateTime).
		AddOutputField("ResponsibleEmployee", "Employee", model.String).
		AddOutputField("FourEyesPrinciple", "Employee", model.Boolean).
		AddOutputField("LastTime", "Employee", model.DateTime).
		AddRule("R1").
		AddInputEntry(`"Car Accident"`, model.SFEEL).
		AddInputEntry("<1000", model.SFEEL).
		AddInputEntry(`>=DateAndTime("2021-01-02T12:00:00")`, model.SFEEL).
		AddOutputEntry(`"Müller"`, model.SFEEL).
		AddOutputEntry("false", model.SFEEL).
		AddOutputEntry(`DateAndTime("2023-01-03T23:59:59")`, model.SFEEL).
		BuildRule().
		AddRule("R2").
		AddInputEntry(`"Car Accident"`, model.SFEEL).
		AddInputEntry("[1000..10000]", model.SFEEL).
		AddInputEntry("-", model.SFEEL).
		AddOutputEntry(`"Schulz"`, model.SFEEL).
		AddOutputEntry("false", model.SFEEL).
		AddOutputEntry("-", model.SFEEL).
		BuildRule().
		AddRule("R3").
		AddInputEntry("-", model.SFEEL).
		AddInputEntry(">=10000", model.SFEEL).
		AddInputEntry("-", model.SFEEL).
		AddOutputEntry("-", model.SFEEL).
		AddOutputEntry("true", model.SFEEL).
		AddOutputEntry("-", model.SFEEL).
		BuildRule().
		Build()

	// Convert Table Into Grule Rules
	converter, _ := conv.CreateTableConverterFactory().GetTableConverter(model.GRULE, model.GRL)
	rules, err := table.Convert(converter)
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
