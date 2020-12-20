package cypher_go_dsl

type ProcedureCall struct {
	name          ProcedureName
	arguments     Arguments
	yieldItems    YieldItems
	optionalWhere Where
	key           string
	notNil        bool
	err           error
}

func ProcedureCallCreate(name ProcedureName, arguments Arguments, yieldItems YieldItems, optionalWhere Where) ProcedureCall {
	if name.getError() != nil {
		return ProcedureCall{err: name.getError()}
	}
	if arguments.getError() != nil {
		return ProcedureCall{err: arguments.getError()}
	}
	if yieldItems.getError() != nil {
		return ProcedureCall{err: yieldItems.getError()}
	}
	if optionalWhere.getError() != nil {
		return ProcedureCall{err: optionalWhere.getError()}
	}
	p := ProcedureCall{
		name:          name,
		arguments:     arguments,
		yieldItems:    yieldItems,
		optionalWhere: optionalWhere,
		notNil:        true,
	}
	p.key = getAddress(&p)
	return p
}

func (p ProcedureCall) doesReturnElement() bool {
	return p.yieldItems.isNotNil()
}

func (p ProcedureCall) getError() error {
	return p.err
}

func (p ProcedureCall) accept(visitor *CypherRenderer) {
	visitor.enter(p)
	p.name.accept(visitor)
	VisitIfNotNull(p.arguments, visitor)
	VisitIfNotNull(p.yieldItems, visitor)
	VisitIfNotNull(p.optionalWhere, visitor)
	visitor.leave(p)
}

func (p ProcedureCall) enter(renderer *CypherRenderer) {
	renderer.append("CALL ")
}

func (p ProcedureCall) leave(renderer *CypherRenderer) {
	renderer.append(".")
}

func (p ProcedureCall) getKey() string {
	return p.key
}

func (p ProcedureCall) isNotNil() bool {
	return p.notNil
}

func (p ProcedureCall) doesReturnElements() bool {
	return p.yieldItems.isNotNil()
}

/**
 * The union of a buildable statement and call exposing new arguments and yields.
 */
type OngoingStandaloneCallWithoutArguments interface {
	BuildableStatement
	AsFunction
	withArgs(arguments ...Expression) OngoingStandaloneCallWithoutArguments
	YieldSymbolic(name ...SymbolicName) OngoingStandaloneCallWithReturnFields
	YieldString(yieldedItems ...string) OngoingStandaloneCallWithReturnFields
	Yield(name ...SymbolicName) OngoingStandaloneCallWithReturnFields
}

/**
 * The union of a buildable statement and call exposing yields.
 */
type OngoingStandaloneCallWithArguments interface {
	BuildableStatement
	AsFunction
	YieldSymbolic(name ...SymbolicName) OngoingStandaloneCallWithReturnFields
	YieldString(yieldedItems ...string) OngoingStandaloneCallWithReturnFields
	Yield(name ...SymbolicName) OngoingStandaloneCallWithReturnFields
}

/**
 * A buildable statement exposing where and return clauses.
 */
type OngoingStandaloneCallWithReturnFields interface {
	BuildableStatement
	ExposesWhere
	ExposesReturning
	ExposesWith
	ExposesSubqueryCall
}

type OngoingInQueryCallWithoutArguments interface {
	WithArgs(arguments ...Expression) OngoingInQueryCallWithArguments
	YieldSymbolic(name ...SymbolicName) OngoingInQueryCallWithReturnFields
	YieldString(yieldedItems ...string) OngoingInQueryCallWithReturnFields
	Yield(aliasedResultFields ...AliasedExpression) OngoingInQueryCallWithReturnFields
}

type OngoingInQueryCallWithArguments interface {
	YieldSymbolic(name ...SymbolicName) OngoingInQueryCallWithReturnFields
	YieldString(yieldedItems ...string) OngoingInQueryCallWithReturnFields
	Yield(aliasedResultFields ...AliasedExpression) OngoingInQueryCallWithReturnFields
}

type OngoingInQueryCallWithReturnFields interface {
	ExposesWhere
	ExposesReturning
	ExposesWith
	ExposesSubqueryCall
}

type ProcedureCallBuilder interface {
	ExposesWhere
	ExposesReturning
	BuildableStatement
	isNotNil() bool
}
