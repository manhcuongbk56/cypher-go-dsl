package cypher

type StatementBuilder interface {
	ExposesMatch
	ExposesCreate
}

type BuildableStatement interface {
	Build() (Statement, error)
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

	Call1(namespaceAndProcedure ...string) OngoingInQueryCallWithoutArguments
}

type OngoingReadingAndReturn interface {
	TerminalExposesOrderBy
	TerminalExposesSkip
	TerminalExposesLimit
	BuildableStatement
}

type OngoingReadingAndWith interface {
	OngoingReading
	ExposesMatch
}

type OngoingReadingWithWhere interface {
	OngoingReading
	ExposesMatch
	ExposesExistentialSubqueryCall
	And(condition Condition) OngoingReadingWithWhere
	AndPattern(pattern RelationshipPattern) OngoingReadingWithWhere
	Or(condition Condition) OngoingReadingWithWhere
	OrPattern(pattern RelationshipPattern) OngoingReadingWithWhere
}

type OngoingUnwind interface {
	As(variable string) OngoingReading
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
	Where(condition Condition) OrderableOngoingReadingAndWithWithWhere
	WherePattern(pattern RelationshipPattern) OrderableOngoingReadingAndWithWithWhere
}

type OrderableOngoingReadingAndWithWithWhere interface {
	OrderableOngoingReadingAndWith
	And(condition Condition) OrderableOngoingReadingAndWithWithWhere
	AndPattern(pattern RelationshipPattern) OrderableOngoingReadingAndWithWithWhere
	Or(condition Condition) OrderableOngoingReadingAndWithWithWhere
	OrPattern(pattern RelationshipPattern) OrderableOngoingReadingAndWithWithWhere
}

type ExposesExistentialSubqueryCall interface {
	AsCondition() Condition
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
	And1(expression Expression) OngoingOrderDefinition
}

type OngoingOrderDefinition interface {
	ExposesSkip
	ExposesLimit
	Descending() OngoingReadingAndWithWithWhereAndOrder
	Ascending() OngoingReadingAndWithWithWhereAndOrder
}

type ExposesOrderBy interface {
	OrderBySortItem(sortItem ...SortItem) OrderableOngoingReadingAndWithWithWhere
	OrderByExpression(expression Expression) OngoingOrderDefinition
}

type ExposesSkip interface {
	Skip(number int) ExposesLimitAndOngoingReadingAndWith
}

type ExposesLimit interface {
	Limit(number int) OngoingReadingAndWith
}

type ExposesUpdatingClause interface {
	ExposesDelete
	ExposesMerge
	ExposesSetAndRemove
}

type ExposesDelete interface {
	DeleteByString(variables ...string) OngoingUpdate
	DeleteByNamed(variables ...Named) OngoingUpdate
	Delete(expressions ...Expression) OngoingUpdate
	DetachDeleteByString(variables ...string) OngoingUpdate
	DetachDeleteByNamed(variables ...Named) OngoingUpdate
	DetachDelete(expressions ...Expression) OngoingUpdate
}

type ExposesWith interface {
	WithByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere
	WithByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere
	With(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere
	WithDistinctByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere
	WithDistinctByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere
	WithDistinct(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere
}

type ExposesSet interface {
	Set(expressions ...Expression) BuildableStatementAndOngoingMatchAndUpdate
	SetWithNamed(variable Named, expression Expression) BuildableStatementAndOngoingMatchAndUpdate
}

type ExposesSetAndRemove interface {
	ExposesSet
	SetByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate
	RemoveByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate
	Remove(properties ...Property) BuildableStatementAndOngoingMatchAndUpdate
}

type OngoingMatchAndUpdate interface {
	ExposesReturning
	ExposesWith
	ExposesUpdatingClause
	ExposesCreate
}

type ExposesMergeAction interface {
	OnCreate() OngoingMergeAction
	OnMatch() OngoingMergeAction
}

type OngoingMatchAndUpdateAndBuildableStatementAndExposesMergeAction interface {
	OngoingMatchAndUpdate
	BuildableStatement
	ExposesMergeAction
}

type OngoingUpdateAndExposesSet interface {
	OngoingUpdate
	ExposesSet
}

type OngoingMergeAction interface {
	GetErr() error
	Set(expressions ...Expression) OngoingMatchAndUpdateAndBuildableStatementAndExposesMergeAction
	SetWithNamed(variable Named, expression Expression) OngoingMatchAndUpdateAndBuildableStatementAndExposesMergeAction
}

type OngoingMatchAndReturnWithOrder interface {
	TerminalExposesSkip
	TerminalExposesLimit
	BuildableStatement
	And(expression Expression) TerminalOngoingOrderDefinition
}

type TerminalExposesLimit interface {
	Limit(number int) BuildableStatement
}

type TerminalOngoingOrderDefinition interface {
	TerminalExposesSkip
	TerminalExposesLimit
	BuildableStatement
	Descending() OngoingMatchAndReturnWithOrder
	Ascending() OngoingMatchAndReturnWithOrder
}

type TerminalExposesOrderBy interface {
	OrderBySortItem(sortItem ...SortItem) OngoingMatchAndReturnWithOrder
	OrderBy(expression Expression) TerminalOngoingOrderDefinition
}

type TerminalExposesLimitAndBuildableStatement interface {
	TerminalExposesLimit
	BuildableStatement
}

type TerminalExposesSkip interface {
	Skip(number int) TerminalExposesLimitAndBuildableStatement
}
