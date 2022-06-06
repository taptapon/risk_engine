package core

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	//"log"
)

type Dsl struct {
	Workflow []FlowNode    `yaml:"workflow,flow"`
	Rulesets []RulesetNode `yaml:"rulesets,flow"`
	//	Conditionals    []Conditional    `yaml:"conditionals,flow"`
	Abtests []AbtestNode `yaml:"abtests,flow"`
	//	DecisionTrees   []DecisionTree   `yaml:"decisiontrees,flow"`
	//	DecisionMatrixs []DecisionMatrix `yaml:"decisionmatrixs,flow"`
	//	ScoreCards      []ScoreCard      `yaml:"scorecards,flow"`
}

//load dsl from file
func LoadDslFromFile(file string) *Dsl {
	dsl := new(Dsl)
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, dsl)

	if err != nil {
		panic(err)
	}
	return dsl
}
