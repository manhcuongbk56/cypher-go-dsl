package cypher_go_dsl

type ExpressionContainer struct {
	expression Expression
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

func (e *ExpressionContainer) IsEqualTo(rhs Expression) ExpressionContainer  {
	newExpression := NewComparison(e.expression, EQUALITY, rhs)
	e.expression = newExpression
	return e
}

func (e *ExpressionContainer) Lt(rhs Expression) ExpressionContainer  {
}

func (e *ExpressionContainer) Lte(rhs Expression) ExpressionContainer  {
}

func (e *ExpressionContainer) Gt(rhs Expression) ExpressionContainer  {
}

func (e *ExpressionContainer) Gte(rhs Expression) ExpressionContainer  {
}

func (e *ExpressionContainer) IsTrue() ExpressionContainer  {
}

func (e *ExpressionContainer) IsFalse() ExpressionContainer  {
}

func (e *ExpressionContainer) Matches(expression Expression) ExpressionContainer  {
}

func (e *ExpressionContainer) MatchesPattern(pattern string) ExpressionContainer  {
}

func (e *ExpressionContainer) StartWiths(rhs Expression) ExpressionContainer  {
}

func (e *ExpressionContainer) Contains(rhs Expression) ExpressionContainer  {
}

func (e *ExpressionContainer) EndsWith(rhs Expression) ExpressionContainer  {
}

func (e *ExpressionContainer) Concat(rhs Expression) ExpressionContainer  {
}

func (e *ExpressionContainer) Add(rhs Expression) ExpressionContainer  {
}

func (e *ExpressionContainer) Subtract(rhs Expression) ExpressionContainer  {
}

func (e *ExpressionContainer) Multiply(rhs Expression) ExpressionContainer  {
}

func (e *ExpressionContainer) Divide(rhs Expression) ExpressionContainer  {
}

func (e *ExpressionContainer) Remainder(rhs Expression) ExpressionContainer  {
}

func (e *ExpressionContainer) Pow(rhs Expression) ExpressionContainer  {
}

func (e *ExpressionContainer) IsNull() ExpressionContainer  {
}

func (e *ExpressionContainer) IsNotNull() ExpressionContainer  {
}

func (e *ExpressionContainer) In() ExpressionContainer  {
}

func (e *ExpressionContainer) IsEmpty() ExpressionContainer  {
}

func (e *ExpressionContainer) Descending() ExpressionContainer  {
	sortItem := CreateDescendingSortItem(e.expression)
	e.expression = sortItem
	return *e
}

func (e *ExpressionContainer) Ascending() ExpressionContainer  {
	sortItem := CreateAscendingSortItem(e.expression)
	e.expression = sortItem
	return *e
}

func (e *ExpressionContainer) IsNotEqualTo(rhs Expression) ExpressionContainer  {
	newExpression := NewComparison(e.expression, NOT_EQUALITY, rhs)
	e.expression = newExpression
	return e
}


type Operator struct {
	representation 	string
	operatorType 	OperatorType
}

var ADDITION = Operator{
	representation: "+",
	operatorType : BINARY,
}

var EQUALITY = Operator{
	representation: "=",
	operatorType : BINARY,
}

var NOT_EQUALITY = Operator{
	representation: "="
	operatorType: BINARY,
}


type OperatorType string

const (
	BINARY OperatorType = "binary"
	PREFIX = "prefix"
	POSTFIX = "postfix"
	PROPERTY = "property"
	LABEL = "label"
)



