package cypher_go_dsl

import "errors"

type DefaultStatementBuilder struct {
	invalidReason string
	currentSinglePartElements []Visitable
	currentOngoingMatch       *MatchBuilder
}

func NewDefaultBuilder() DefaultStatementBuilder {
	return DefaultStatementBuilder{
		currentSinglePartElements: make([]Visitable, 0),
	}
}

func (d DefaultStatementBuilder) Match(pattern ...PatternElement) OngoingReadingWithoutWhere {
	if pattern == nil || len(pattern) == 0 {
		return DefaultStatementBuilder{invalidReason: "patterns to match is required"}
	}
	if d.currentOngoingMatch != nil {
		d.currentSinglePartElements = append(d.currentSinglePartElements, d.currentOngoingMatch.buildMatch())
	}
	d.currentOngoingMatch = &MatchBuilder{
		optional: false,
	}
	d.currentOngoingMatch.patternList = append(d.currentOngoingMatch.patternList, pattern...)
	return d
}

func (d DefaultStatementBuilder) OptionalMatch(pattern ...PatternElement) {
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
	d DefaultStatementBuilder
	distinct bool
	returnList []IsExpression
	orderBuilder OrderBuilder
}

func (b DefaultStatementWithReturnBuilder) Build() Statement {
	var returning *Return
	if len(b.returnList) > 0 {
		returnItems := ExpressionList{b.returnList}
		var distinctInstance *Distinct
		if b.distinct {
			distinctInstance = &Distinct{}
		}
		returning = ReturnByMultiVariable(b.distinct, returnItems, b.orderBuilder.BuildOrder(), &b.orderBuilder.skip,
			&b.orderBuilder.limit)
	}
	return b.d.BuildImpl(returning)
}

func (d DefaultStatementBuilder) Build() Statement {
	return d.BuildImpl(nil)
}

func (d DefaultStatementBuilder) BuildImpl(returning *Return) Statement {
	singlePartQuery, _ := NewSinglePartQuery(d.BuildListOfVisitable(), returning)
	return singlePartQuery
}

func (d DefaultStatementBuilder) BuildListOfVisitable() []Visitable  {
	visitables := make([]Visitable, 0)
	copy(visitables, d.currentSinglePartElements)
	if d.currentOngoingMatch != nil {
		visitables = append(visitables, d.currentOngoingMatch.buildMatch())
		d.currentOngoingMatch = nil;
	}
	d.currentSinglePartElements = d.currentSinglePartElements[:0]
	return visitables
}




