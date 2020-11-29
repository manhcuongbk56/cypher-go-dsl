package cypher_go_dsl

type CompoundCondition struct {
	ConditionContainer
}

func (c CompoundCondition) create(condition TestCondition)  {
	
}

func (c CompoundCondition) isTestCondition() bool {
	return true
}

type TestCondition interface {
	isTestCondition() bool
}


type ConditionContainer struct {
	condition TestCondition
}

func (c ConditionContainer) and(container ConditionContainer) ConditionContainer  {

}