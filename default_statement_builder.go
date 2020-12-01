package cypher_go_dsl

type DefaultStatementBuilder struct {
	OngoingReading
	invalidReason             string
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

func (d DefaultStatementBuilder) closeCurrentOngoingMatch() {

}

func (d DefaultStatementBuilder) returning(expression ...Expression) OngoingReadingAndReturn {
	return d.returningDefault(false, expression...)
}

func (d DefaultStatementBuilder) returningDistinct(expression ...Expression) OngoingReadingAndReturn {
	return d.returningDefault(true, expression...)
}

func (d DefaultStatementBuilder) returningDefault(distinct bool, expression ...Expression) OngoingReadingAndReturn {
	withReturnBuilder := DefaultStatementWithReturnBuilder{
		distinct:                distinct,
		defaultStatementBuilder: d,
	}
	withReturnBuilder.AddExpression(expression...)
	return withReturnBuilder
}

func (builder *DefaultStatementWithReturnBuilder) AddExpression(expression ...Expression) {
	builder.returnList = append(builder.returnList, expression...)
}

type DefaultStatementWithReturnBuilder struct {
	defaultStatementBuilder DefaultStatementBuilder
	distinct                bool
	returnList              []Expression
	orderBuilder            OrderBuilder
}

func (builder DefaultStatementWithReturnBuilder) Build() Statement {
	var returning *Return
	if len(builder.returnList) > 0 {
		returnItems := ExpressionList{expressions: builder.returnList}
		returning = ReturnByMultiVariable(builder.distinct, returnItems, builder.orderBuilder.BuildOrder(), builder.orderBuilder.skip,
			builder.orderBuilder.limit)
	}
	return builder.defaultStatementBuilder.BuildImpl(returning)
}

func (d DefaultStatementBuilder) Build() Statement {
	return d.BuildImpl(nil)
}

func (d DefaultStatementBuilder) BuildImpl(returning *Return) Statement {
	singlePartQuery, _ := NewSinglePartQuery(d.BuildListOfVisitable(), returning)
	return *singlePartQuery
}

func (d DefaultStatementBuilder) BuildListOfVisitable() []Visitable {
	visitables := make([]Visitable, 0)
	copy(visitables, d.currentSinglePartElements)
	if d.currentOngoingMatch != nil {
		visitables = append(visitables, d.currentOngoingMatch.buildMatch())
		d.currentOngoingMatch = nil
	}
	d.currentSinglePartElements = d.currentSinglePartElements[:0]
	return visitables
}
