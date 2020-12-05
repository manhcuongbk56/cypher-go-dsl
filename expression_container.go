package cypher_go_dsl

import "fmt"

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
	e.expression = AliasedExpressionCreate(e.expression, alias)
	return *e
}

func (e *ExpressionContainer) IsEqualTo(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, EQUALITY, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) IsNotEqualTo(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, INEQUALITY, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) Lt(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, LESS_THAN, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) Lte(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, LESS_THAN_OR_EQUAL_TO, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) Gt(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, GREATER_THAN, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) Gte(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, GREATER_THAN_OR_EQUAL_TO, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) IsTrue() ConditionContainer {
	return e.IsEqualTo(BooleanLiteral{
		content: true,
	})
}

func (e *ExpressionContainer) IsFalse() ConditionContainer {
	return e.IsEqualTo(BooleanLiteral{
		content: false,
	})
}

func (e *ExpressionContainer) Matches(expression Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, MATCHES, expression)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) MatchesPattern(pattern string) ConditionContainer {
	e.expression = NewComparison(e.expression, MATCHES, StringLiteral{
		content: pattern,
	})
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) StartWiths(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, STARTS_WITH, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) Contains(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, CONTAINS, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) EndsWith(rhs Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, ENDS_WITH, rhs)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) Concat(rhs Expression) ExpressionContainer {
	e.expression = NewComparison(e.expression, CONCAT, rhs)
	return *e
}

func (e *ExpressionContainer) Add(rhs Expression) ExpressionContainer {
	e.expression = NewComparison(e.expression, ADDITION, rhs)
	return *e
}

func (e *ExpressionContainer) Subtract(rhs Expression) ExpressionContainer {
	e.expression = NewComparison(e.expression, SUBTRACTION, rhs)
	return *e
}

func (e *ExpressionContainer) Multiply(rhs Expression) ExpressionContainer {
	e.expression = NewComparison(e.expression, MULTIPLICATION, rhs)
	return *e
}

func (e *ExpressionContainer) Divide(rhs Expression) ExpressionContainer {
	e.expression = NewComparison(e.expression, DIVISION, rhs)
	return *e
}

func (e *ExpressionContainer) Remainder(rhs Expression) ExpressionContainer {
	e.expression = NewComparison(e.expression, MODULO_DIVISION, rhs)
	return *e
}

func (e *ExpressionContainer) Pow(rhs Expression) ExpressionContainer {
	e.expression = NewComparison(e.expression, EXPONENTIATION, rhs)
	return *e
}

func (e *ExpressionContainer) IsNull() ConditionContainer {
	e.expression = NewComparisonWithConstant(IS_NOT_NULL, e.expression)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) IsNotNull() ConditionContainer {
	e.expression = NewComparisonWithConstant(IS_NOT_NULL, e.expression)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) In(haystack Expression) ConditionContainer {
	e.expression = NewComparison(e.expression, IN, haystack)
	return ConditionContainer{*e}
}

func (e *ExpressionContainer) IsEmpty() ConditionContainer {
	e.expression = size(e.expression)
	return e.IsEqualTo(NumberLiteral{
		content: 0,
	})
}

func (e *ExpressionContainer) Descending() SortItem {
	return CreateDescendingSortItem(e.expression)
}

func (e *ExpressionContainer) Ascending() SortItem {
	return CreateAscendingSortItem(e.expression)
}

type Operator struct {
	representation string
	operatorType   OperatorType
	key            string
	notNil         bool
}

func (o Operator) isNotNil() bool {
	return o.notNil
}

func (o Operator) accept(visitor *CypherRenderer) {
	o.key = fmt.Sprint(&o)
	visitor.enter(o)
	visitor.leave(o)
}

func (o Operator) enter(renderer *CypherRenderer) {
	operatorType := o.operatorType
	if operatorType == LABEL {
		return
	}
	if operatorType != PREFIX && o != EXPONENTIATION {
		renderer.builder.WriteString(" ")
	}
	renderer.builder.WriteString(o.representation)
	if operatorType != POSTFIX && o != EXPONENTIATION {
		renderer.builder.WriteString(" ")
	}
}

func (o Operator) leave(renderer *CypherRenderer) {
}

func (o Operator) getKey() string {
	return o.key
}

func (o Operator) isUnary() bool {
	return o.operatorType != BINARY
}

var ADDITION = createBinaryOperator("+")
var SUBTRACTION = createBinaryOperator("-")
var MULTIPLICATION = createBinaryOperator("*")
var DIVISION = createBinaryOperator("/")
var MODULO_DIVISION = createBinaryOperator("%")
var EXPONENTIATION = createBinaryOperator("^")
var EQUALITY = createBinaryOperator("=")
var INEQUALITY = createBinaryOperator("<>")
var LESS_THAN = createBinaryOperator("<")
var GREATER_THAN = createBinaryOperator(">")
var LESS_THAN_OR_EQUAL_TO = createBinaryOperator("<=")
var GREATER_THAN_OR_EQUAL_TO = createBinaryOperator(">=")
var STARTS_WITH = createBinaryOperator("STARTS WITH")
var ENDS_WITH = createBinaryOperator("ENDS WITH")
var CONTAINS = createBinaryOperator("CONTAINS")
var AND = createBinaryOperator("AND")
var OR = createBinaryOperator("OR")
var XOR = createBinaryOperator("XOR")
var CONCAT = createBinaryOperator("+")
var MATCHES = createBinaryOperator("=~")
var IN = createBinaryOperator("IN")
var ASSIGMENT = createBinaryOperator("=")
var PIPE = createBinaryOperator("|")
var IS_NULL = createPostfixOperator("IS NULL")
var IS_NOT_NULL = createPostfixOperator("IS NOT NULL")
var NOT = createPrefixOperator("NOT")
var SET = createPropertyOperator("=")
var GET = createPropertyOperator(".")
var MUTATE = createPropertyOperator("+=")
var SET_LABEL = createLabelOperator("")
var REMOVE_LABEL = createLabelOperator("")

func createBinaryOperator(representation string) Operator {
	return Operator{
		representation: representation,
		operatorType:   BINARY,
	}
}

func createPropertyOperator(representation string) Operator {
	return Operator{
		representation: representation,
		operatorType:   PROPERTY,
	}
}

func createLabelOperator(representation string) Operator {
	return Operator{
		representation: representation,
		operatorType:   LABEL,
	}
}

func createPrefixOperator(representation string) Operator {
	return Operator{
		representation: representation,
		operatorType:   PREFIX,
	}
}

func createPostfixOperator(representation string) Operator {
	return Operator{
		representation: representation,
		operatorType:   POSTFIX,
	}
}

type OperatorType string

const (
	BINARY   OperatorType = "binary"
	PREFIX                = "prefix"
	POSTFIX               = "postfix"
	PROPERTY              = "property"
	LABEL                 = "label"
)
