package cypher_go_dsl

import "runtime/trace"

type DefaultStatementBuilder struct {
	currentSinglePartElements []Visitable
	currentOngoingMatch       MatchBuilder
}

func NewDefaultBuilder() DefaultStatementBuilder {
	return DefaultStatementBuilder{
		currentSinglePartElements: make([]Visitable, 0),
	}
}

func (d DefaultStatementBuilder) Match(pattern ...cyphergodsl.PatternElement) OngoingReadingWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) OptionalMatch(pattern ...cyphergodsl.PatternElement) {
	panic("implement me")
}

func (d DefaultStatementBuilder) closeCurrentOngoingMatch()  {
	
}

func (d DefaultStatementBuilder) returning(expression ...IsExpression) OngoingReadingAndReturn {
	return d.returningDefault(false, expression...)
}

func (d DefaultStatementBuilder) returningDistinct(expression ...IsExpression) OngoingReadingAndReturn {
	panic("implement me")
}

func (d DefaultStatementBuilder) returningDefault(distinct bool, expression ...IsExpression) OngoingReadingAndReturn {
	withReturnBuilder := DefaultStatementWithReturnBuilder{
		distinct: distinct,
	}
	withReturnBuilder.AddExpression(expression...)
}

func (builder DefaultStatementWithReturnBuilder) AddExpression(expression ...IsExpression)  {
	builder.returnList = append(builder.returnList, expression...)
}

type DefaultStatementWithReturnBuilder struct {
	distinct bool
	returnList []IsExpression
}

func (b DefaultStatementWithReturnBuilder) Build() Statement {
	
}




