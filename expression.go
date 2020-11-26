package cypher_go_dsl

type Expression interface {
	Visitable
	IsExpression() bool
}

type ExpressionStruct struct {
}

type Condition struct {
	expression ExpressionStruct
}

func (lhs ExpressionStruct) IsEqualTo(rhs ExpressionStruct) Comparison {
	return Comparison{left: lhs, operator: EQUALITY, right: rhs}
}

func (lhs ExpressionStruct) Accept(visitor *CypherRenderer) {
	panic("implement me")
}

func (lhs Condition) And(rhs Condition) Condition {
	panic("Implement me")

}

func (lhs Condition) Or(rhd Condition) Condition {
	panic("Implement me")
}

type Operator string

func (o Operator) Accept(visitor *CypherRenderer) {
	panic("implement me")
}

const (
	EQUALITY = "equality"
)

type Comparison struct {
	ExpressionStruct
	left     ExpressionStruct
	operator Operator
	right    ExpressionStruct
}

func (c Comparison) IsExpression() bool {
	return true
}

func (c Comparison) Accept(visitor *CypherRenderer) {
	panic("implement me")
}

func NameOrExpression(expression Expression) Expression {
	named, isNamed := expression.(Named)
	if isNamed && named.getSymbolicName() != nil {
		return named.getSymbolicName()
	}
	return expression
}
