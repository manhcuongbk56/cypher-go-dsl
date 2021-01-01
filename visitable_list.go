package cypher

type FunctionArgumentList struct {
	expressions []Visitable
	key         string
	notNil      bool
	err         error
}

func FunctionArgumentListCreate(visitables ...Visitable) FunctionArgumentList {
	argumentList := FunctionArgumentList{expressions: visitables}
	argumentList.key = getAddress(&argumentList)
	return argumentList
}

func (v FunctionArgumentList) GetError() error {
	return v.err
}

func (v FunctionArgumentList) isNotNil() bool {
	return v.notNil
}

func (v FunctionArgumentList) getKey() string {
	return v.key
}

func (v FunctionArgumentList) PrepareVisit(child Visitable) Visitable {
	expression, isExpression := child.(Expression)
	if !isExpression {
		return child
	}
	return NameOrExpression(expression)
}

func (v FunctionArgumentList) accept(visitor *CypherRenderer) {
	(*visitor).enter(v)
	for _, expression := range v.expressions {
		v.PrepareVisit(expression).accept(visitor)
	}
	(*visitor).leave(v)
}

func (v FunctionArgumentList) enter(renderer *CypherRenderer) {
}

func (v FunctionArgumentList) leave(renderer *CypherRenderer) {
}
