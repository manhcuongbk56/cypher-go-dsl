package cypher_go_dsl

type DefaultStatementWithUpdateBuilder struct {
	returnList   []Expression
	orderBuilder OrderBuilder
	distinct     bool
	builder      UpdatingClauseBuilder
	notNil       bool
}

func (d DefaultStatementWithUpdateBuilder) isNotNil() bool {
	return d.notNil
}

func (d DefaultStatementWithUpdateBuilder) and(expression Expression) TerminalOngoingOrderDefinition {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) descending() OngoingMatchAndReturnWithOrder {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) ascending() OngoingMatchAndReturnWithOrder {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) orderBySortItem(sortItem ...SortItem) OngoingMatchAndReturnWithOrder {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) orderBy(expression Expression) TerminalOngoingOrderDefinition {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) skip(number int) TerminalExposesLimitAndBuildableStatement {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) limit(number int) BuildableStatement {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) build() Statement {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) returningByString(variables ...string) OngoingReadingAndReturn {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) returningByNamed(variables ...Named) OngoingReadingAndReturn {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) returning(expression ...Expression) OngoingReadingAndReturn {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) returningDistinctByString(variables ...string) OngoingReadingAndReturn {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) returningDistinctByNamed(variables ...Named) OngoingReadingAndReturn {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) returningDistinct(expression ...Expression) OngoingReadingAndReturn {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) withByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) withByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) with(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) withDistinctByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) withDistinctByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) withDistinct(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) deleteByString(variables ...string) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) deleteByNamed(variables ...Named) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) delete(expressions ...Expression) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) detachDeleteByString(variables ...string) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) detachDeleteByNamed(variables ...Named) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) detachDelete(expressions ...Expression) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) merge(pattern ...PatternElement) {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) set(expressions ...Expression) BuildableStatementAndOngoingMatchAndUpdate {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) setWithNamed(variable Named, expression Expression) BuildableStatementAndOngoingMatchAndUpdate {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) setByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) removeByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) remove(properties ...Property) BuildableStatementAndOngoingMatchAndUpdate {
	panic("implement me")
}

func (d DefaultStatementWithUpdateBuilder) create(element ...PatternElement) {
	panic("implement me")
}
