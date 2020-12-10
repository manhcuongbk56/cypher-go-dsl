package cypher_go_dsl

import (
	"fmt"
	errors "golang.org/x/xerrors"
)

type FunctionInvocation struct {
	functionName string
	arguments    FunctionArgumentList
	key          string
	notNil       bool
	err          error
}

func FunctionInvocationCreate(definition FunctionDefinition, expressions ...Expression) (FunctionInvocation, error) {
	if len(expressions) == 0 {
		return FunctionInvocation{}, errors.New("need expression")
	}
	arguments := make([]Visitable, len(expressions))
	for _, expression := range expressions {
		arguments = append(arguments, expression)
	}
	f := FunctionInvocation{
		functionName: definition.getImplementationName(),
		arguments: FunctionArgumentList{
			expressions: arguments,
		},
	}
	f.key = getAddress(&f)
	return f, nil
}

func CreateWithPatternElement(definition FunctionDefinition, element PatternElement) (FunctionInvocation, error) {
	arguments := make([]Visitable, 1)
	arguments = append(arguments, element)
	f := FunctionInvocation{
		functionName: definition.getImplementationName(),
		arguments: FunctionArgumentList{
			expressions: arguments,
		},
	}
	f.key = getAddress(&f)
	return f, nil
}

func CreateWithPattern(definition FunctionDefinition, pattern Pattern) (FunctionInvocation, error) {
	arguments := make([]Visitable, len(pattern.patternElements))
	for _, expression := range pattern.patternElements {
		arguments = append(arguments, expression)
	}
	f := FunctionInvocation{
		functionName: definition.getImplementationName(),
		arguments: FunctionArgumentList{
			expressions: arguments,
		},
	}
	f.key = getAddress(&f)
	return f, nil
}

func CreateDistinct(definition FunctionDefinition, expressions ...Expression) (FunctionInvocation, error) {
	if !definition.isAggregate() {
		return FunctionInvocation{}, errors.New("the distinct operator can only be applied within aggregate functions")
	}
	if len(expressions) == 0 {
		return FunctionInvocation{}, errors.New("need expression")
	}
	arguments := make([]Visitable, len(expressions))
	arguments = append(arguments, DistinctExpression{
		delegate: expressions[0],
	})
	for _, expression := range expressions[1:] {
		arguments = append(arguments, expression)
	}
	f := FunctionInvocation{
		functionName: definition.getImplementationName(),
		arguments: FunctionArgumentList{
			expressions: arguments,
		},
	}
	f.key = getAddress(&f)
	return f, nil
}

func (f FunctionInvocation) getError() error {
	return f.err
}

func (f FunctionInvocation) isNotNil() bool {
	return f.notNil
}

func (f FunctionInvocation) accept(visitor *CypherRenderer) {
	visitor.enter(f)
	f.arguments.accept(visitor)
	visitor.leave(f)
}

func (f FunctionInvocation) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(f.functionName)
	renderer.builder.WriteString("(")
}

func (f FunctionInvocation) leave(renderer *CypherRenderer) {
	renderer.builder.WriteString(")")
}

func (f FunctionInvocation) getKey() string {
	return f.key
}

func (f FunctionInvocation) GetExpressionType() ExpressionType {
	return EXPRESSION
}

type FunctionDefinition interface {
	getImplementationName() string
	isAggregate() bool
}
