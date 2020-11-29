package cypher_go_dsl

type ExpressionContainer struct {
	expression Expression
	condition Condition
}

type ABC struct {
	ExpressionContainer
	a string
}

func (A ABC) IsExpression() bool {
	panic("implement me")
}

func (e ExpressionContainer) accept(visitor *CypherRenderer) {
	e.expression.accept(visitor)
}

func (e ExpressionContainer) enter(renderer *CypherRenderer) {
	panic("can not enter expression container")
}

func (e ExpressionContainer) leave(renderer *CypherRenderer) {
	panic("can not leave expression container")
}

func (e ExpressionContainer) getKey() string {
	panic("Expression container have no key")
}

func (e *ExpressionContainer) As(alias string) ExpressionContainer {
	aliasedExpression := AliasedExpression{
		delegate: e.expression,
		alias:    alias,
	}
	e.expression = aliasedExpression
	return *e
}

func (e ExpressionContainer) IsEqualTo(rhs Expression) ExpressionContainer  {
	abc := ABC{
		a: "",
	}
	newE := ExpressionContainer{
		expression: abc,
	}

}


type Operator struct {
	representation 	string
	operatorType 	OperatorType
}

var ADDITION = Operator{
	representation: "+",
	operatorType : BINARY,
}

type OperatorType string

const (
	BINARY OperatorType = "binary"
	PREFIX = "prefix"
	POSTFIX = "postfix"
	PROPERTY = "property"
	LABEL = "label"
)



