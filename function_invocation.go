package cypher

import (
	errors "golang.org/x/xerrors"
)

type FunctionInvocation struct {
	ExpressionContainer
	functionName string
	arguments    FunctionArgumentList
	key          string
	notNil       bool
	err          error
}

func FunctionInvocationCreate1(definition FunctionDefinition) FunctionInvocation {
	return FunctionInvocationCreateWithFunctionName(definition.getImplementationName(), nil)
}

func FunctionInvocationCreateWithFunctionName(functionName string, arguments ...Expression) FunctionInvocation {
	if arguments != nil {
		for _, expression := range arguments {
			if expression.getError() != nil {
				return FunctionInvocationError(expression.getError())
			}
		}
	}
	argumentVisitables := make([]Visitable, len(arguments))
	for i := range arguments {
		argumentVisitables[i] = arguments[i]
	}
	f := FunctionInvocation{
		functionName: functionName,
		arguments:    FunctionArgumentListCreate(argumentVisitables...),
		notNil:       true,
	}
	f.key = getAddress(&f)
	f.ExpressionContainer = ExpressionWrap(f)
	return f
}

func FunctionInvocationCreate(definition FunctionDefinition, expressions ...Expression) FunctionInvocation {
	if expressions != nil {
		for _, expression := range expressions {
			if expression.getError() != nil {
				return FunctionInvocationError(expression.getError())
			}
		}
	}
	if len(expressions) == 0 || expressions[0] == nil || !expressions[0].isNotNil() {
		return FunctionInvocationError(errors.Errorf("expression for %s is required", definition.getImplementationName()))
	}
	arguments := make([]Visitable, len(expressions))
	for i := range expressions {
		arguments[i] = expressions[i]
	}
	f := FunctionInvocation{
		functionName: definition.getImplementationName(),
		arguments:    FunctionArgumentListCreate(arguments...),
		notNil:       true,
	}
	f.key = getAddress(&f)
	f.ExpressionContainer = ExpressionWrap(f)
	return f
}

func FunctionInvocationCreateWithPatternElement(definition FunctionDefinition, element PatternElement) FunctionInvocation {
	if element != nil && element.getError() != nil {
		return FunctionInvocationError(element.getError())
	}
	if element == nil || !element.isNotNil() {
		return FunctionInvocationError(errors.Errorf("the pattern for %s is required", definition.getImplementationName()))
	}
	arguments := make([]Visitable, 0)
	arguments = append(arguments, element)
	f := FunctionInvocation{
		functionName: definition.getImplementationName(),
		arguments:    FunctionArgumentListCreate(arguments...),
		notNil:       true,
	}
	f.key = getAddress(&f)
	f.ExpressionContainer = ExpressionWrap(f)
	return f
}

func FunctionInvocationCreateWithPattern(definition FunctionDefinition, pattern Pattern) FunctionInvocation {
	if pattern.getError() != nil {
		return FunctionInvocationError(pattern.getError())
	}
	if !pattern.isNotNil() {
		return FunctionInvocationError(errors.Errorf("the pattern for %s is required", definition.getImplementationName()))
	}
	arguments := make([]Visitable, len(pattern.patternElements))
	for i := range pattern.patternElements {
		arguments[i] = pattern.patternElements[i]
	}
	f := FunctionInvocation{
		functionName: definition.getImplementationName(),
		arguments:    FunctionArgumentListCreate(arguments...),
		notNil:       true,
	}
	f.key = getAddress(&f)
	f.ExpressionContainer = ExpressionWrap(f)
	return f
}

func FunctionInvocationCreateDistinct(definition FunctionDefinition, expressions ...Expression) FunctionInvocation {
	if expressions != nil {
		for _, expression := range expressions {
			if expression.getError() != nil {
				return FunctionInvocationError(expression.getError())
			}
		}
	}
	if !definition.isAggregate() {
		return FunctionInvocationError(errors.New("the distinct operator can only be applied within aggregate functions"))
	}
	if len(expressions) == 0 || expressions[0] == nil || !expressions[0].isNotNil() {
		return FunctionInvocationError(errors.Errorf("expression for %s is required", definition.getImplementationName()))
	}
	arguments := make([]Visitable, 0)
	arguments = append(arguments, DistinctExpression{
		delegate: expressions[0],
	})
	for _, expression := range expressions[1:] {
		arguments = append(arguments, expression)
	}
	f := FunctionInvocation{
		functionName: definition.getImplementationName(),
		arguments:    FunctionArgumentListCreate(arguments...),
		notNil:       true,
	}
	f.key = getAddress(&f)
	f.ExpressionContainer = ExpressionWrap(f)
	return f
}

func FunctionInvocationError(err error) FunctionInvocation {
	return FunctionInvocation{
		err: err,
	}
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
	renderer.append(f.functionName)
	renderer.append("(")
}

func (f FunctionInvocation) leave(renderer *CypherRenderer) {
	renderer.append(")")
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

type FunctionDefinitionDefault struct {
	functionName string
}

func (f FunctionDefinitionDefault) getImplementationName() string {
	return f.functionName
}

func (f FunctionDefinitionDefault) isAggregate() bool {
	return false
}
