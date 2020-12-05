package cypher_go_dsl

type StatementBuilder interface {
	ExposesMatch
	ExposesCreate
}

type BuildableStatement interface {
	build() Statement
}

type BuildableStatementAndOngoingMatchAndUpdate interface {
	BuildableStatement
	OngoingMatchAndUpdate
}

type OngoingReadingWithoutWhere interface {
	OngoingReading
	ExposesWhere
	ExposesMatch
	ExposesExistentialSubqueryCall
}

type OngoingReading interface {
	ExposesReturning
	ExposesWith
	ExposesUpdatingClause
	ExposesUnwind
	ExposesCreate
	ExposesSubqueryCall

	callExposes(namespaceAndProcedure ...string) OngoingInQueryCallWithoutArguments
}

type OngoingReadingAndReturn interface {
	TerminalExposesOrderBy
	TerminalExposesSkip
	TerminalExposesLimit
	BuildableStatement
}

type OngoingInQueryCallWithoutArguments interface {
	withArgs(arguments ...Expression) OngoingInQueryCallWithArguments
	yieldSymbolic(name SymbolicName) OngoingInQueryCallWithReturnFields
	yieldString(yieldedItems ...string) OngoingInQueryCallWithReturnFields
	yield(aliasedResultFields ...AliasedExpression) OngoingInQueryCallWithReturnFields
}

type OngoingInQueryCallWithArguments interface {
	yieldSymbolic(name SymbolicName) OngoingInQueryCallWithReturnFields
	yieldString(yieldedItems ...string) OngoingInQueryCallWithReturnFields
	yield(aliasedResultFields ...AliasedExpression) OngoingInQueryCallWithReturnFields
}

type OngoingInQueryCallWithReturnFields interface {
	ExposesWhere
	ExposesReturning
	ExposesWith
	ExposesSubqueryCall
}

type OngoingReadingAndWith interface {
	OngoingReading
	ExposesMatch
}

type OngoingReadingWithWhere interface {
	OngoingReading
	ExposesMatch
	ExposesExistentialSubqueryCall
	and(condition Condition) OngoingReadingWithWhere
	andPattern(pattern RelationshipPattern) OngoingReadingWithWhere
	or(condition Condition) OngoingReadingWithWhere
	orPattern(pattern RelationshipPattern) OngoingReadingWithWhere
}

type OngoingUnwind interface {
	as(variable string) OngoingReading
}

type OngoingUpdate interface {
	BuildableStatement
	ExposesCreate
	ExposesMerge
	ExposesDelete
	ExposesReturning
	ExposesWith
	ExposesMergeAction
}

type ExposesLimitAndOngoingReadingAndWith interface {
}

type OrderableOngoingReadingAndWithWithoutWhere interface {
	OrderableOngoingReadingAndWith
	where(condition Condition) OrderableOngoingReadingAndWithWithWhere
	wherePattern(pattern RelationshipPattern) OrderableOngoingReadingAndWithWithWhere
}

type OrderableOngoingReadingAndWithWithWhere interface {
	OrderableOngoingReadingAndWith
	and(condition Condition) OrderableOngoingReadingAndWithWithWhere
	andPattern(pattern RelationshipPattern) OrderableOngoingReadingAndWithWithWhere
	or(condition Condition) OrderableOngoingReadingAndWithWithWhere
	orPattern(pattern RelationshipPattern) OrderableOngoingReadingAndWithWithWhere
}

type ExposesExistentialSubqueryCall interface {
	asCondition() Expression
}

type OrderableOngoingReadingAndWith interface {
	ExposesOrderBy
	ExposesSkip
	ExposesLimit
	OngoingReadingAndWith
}

type OngoingReadingAndWithWithWhereAndOrder interface {
	ExposesSkip
	ExposesLimit
	OngoingReadingAndWith
	and1(expression Expression) OngoingOrderDefinition
}

type OngoingOrderDefinition interface {
	ExposesSkip
	ExposesLimit
	descending() OngoingReadingAndWithWithWhereAndOrder
	ascending() OngoingReadingAndWithWithWhereAndOrder
}

type ExposesOrderBy interface {
	orderBySortItem(sortItem ...SortItem) OrderableOngoingReadingAndWithWithWhere
	orderByExpression(expression Expression) OngoingOrderDefinition
}

type ExposesSkip interface {
	skip(number int) ExposesLimitAndOngoingReadingAndWith
}

type ExposesLimit interface {
	limit(number int) OngoingReadingAndWith
}

type ExposesUpdatingClause interface {
	ExposesDelete
	ExposesMerge
	ExposesSetAndRemove
}

type ExposesDelete interface {
	deleteByString(variables ...string) OngoingUpdate
	deleteByNamed(variables ...Named) OngoingUpdate
	delete(expressions ...Expression) OngoingUpdate
	detachDeleteByString(variables ...string) OngoingUpdate
	detachDeleteByNamed(variables ...Named) OngoingUpdate
	detachDelete(expressions ...Expression) OngoingUpdate
}

type ExposesWith interface {
	withByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere
	withByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere
	with(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere
	withDistinctByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere
	withDistinctByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere
	withDistinct(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere
}

type ExposesSet interface {
	set(expressions ...Expression) BuildableStatementAndOngoingMatchAndUpdate
	setWithNamed(variable Named, expression Expression) BuildableStatementAndOngoingMatchAndUpdate
}

type ExposesSetAndRemove interface {
	ExposesSet
	setByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate
	removeByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate
	remove(properties ...Property) BuildableStatementAndOngoingMatchAndUpdate
}

type OngoingMatchAndUpdate interface {
	ExposesReturning
	ExposesWith
	ExposesUpdatingClause
	ExposesCreate
}

type ExposesMergeAction interface {
	onCreate() OngoingMergeAction
	onMatch() OngoingMergeAction
}

type OngoingMatchAndUpdateAndBuildableStatementAndExposesMergeAction interface {
	OngoingMatchAndUpdate
	BuildableStatement
	ExposesMergeAction
}

type OngoingMergeAction interface {
	set(expressions ...Expression) OngoingMatchAndUpdateAndBuildableStatementAndExposesMergeAction
	setWithNamed(variable Named, expression Expression) OngoingMatchAndUpdateAndBuildableStatementAndExposesMergeAction
}

type OngoingMatchAndReturnWithOrder interface {
	TerminalExposesSkip
	TerminalExposesLimit
	BuildableStatement
	and(expression Expression) TerminalOngoingOrderDefinition
}

type TerminalExposesLimit interface {
	limit(number int) BuildableStatement
}

type UpdatingClauseBuilder interface {
	build() UpdatingClause
}

type TerminalOngoingOrderDefinition interface {
	TerminalExposesSkip
	TerminalExposesLimit
	BuildableStatement
	descending() OngoingMatchAndReturnWithOrder
	ascending() OngoingMatchAndReturnWithOrder
}

type TerminalExposesOrderBy interface {
	orderBySortItem(sortItem ...SortItem) OngoingMatchAndReturnWithOrder
	orderBy(expression Expression) TerminalOngoingOrderDefinition
}

type TerminalExposesLimitAndBuildableStatement interface {
	TerminalExposesLimit
	BuildableStatement
}

type TerminalExposesSkip interface {
	skip(number int) TerminalExposesLimitAndBuildableStatement
}
