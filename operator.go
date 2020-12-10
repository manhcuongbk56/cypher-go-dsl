package cypher_go_dsl

import "fmt"

type Operator struct {
	representation string
	operatorType   OperatorType
	key            string
	notNil         bool
	err error
}

func (o Operator) getError() error {
	return o.err
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
