package cypher

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
	if name.GetError() != nil {
		return ProcedureCall{err: name.GetError()}
	}
	if arguments.GetError() != nil {
		return ProcedureCall{err: arguments.GetError()}
	}
	if yieldItems.GetError() != nil {
		return ProcedureCall{err: yieldItems.GetError()}
	}
	if optionalWhere.GetError() != nil {
		return ProcedureCall{err: optionalWhere.GetError()}
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

func (p ProcedureCall) GetError() error {
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
	Yield(aliasedResultFields ...AliasedExpression) OngoingStandaloneCallWithReturnFields
}

/**
 * The union of a buildable statement and call exposing yields.
 */

type OngoingStandaloneCallWithArguments interface {
	BuildableStatement
	AsFunction
	YieldSymbolic(name ...SymbolicName) OngoingStandaloneCallWithReturnFields
	YieldString(yieldedItems ...string) OngoingStandaloneCallWithReturnFields
	Yield(aliasedResultFields ...AliasedExpression) OngoingStandaloneCallWithReturnFields
}

/**
 * A buildable statement exposing Where and return clauses.
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

type StandaloneCallBuilder struct {
	procedureName    ProcedureName
	arguments        []Expression
	yieldItems       YieldItems
	conditionBuilder ConditionBuilder
	notNil           bool
	err              error
}

func StandaloneCallBuilderCreate(procedureName ProcedureName) StandaloneCallBuilder {
	return StandaloneCallBuilder{
		procedureName: procedureName,
		notNil:        true,
	}
}

func StandaloneCallBuilderError(err error) StandaloneCallBuilder {
	return StandaloneCallBuilder{
		err: err,
	}
}

func (s StandaloneCallBuilder) Where(condition Condition) OngoingReadingWithWhere {
	s.conditionBuilder.Where(condition)
	return DefaultStatementBuilderCreate1(s)
}

func (s StandaloneCallBuilder) WhereConditionContainer(container ConditionContainer) OngoingReadingWithWhere {
	return s.Where(container.Get())
}

func (s StandaloneCallBuilder) WherePattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	return s.Where(RelationshipPatternConditionCreate(pattern))
}

func (s StandaloneCallBuilder) ReturningByString(variables ...string) OngoingReadingAndReturn {
	return s.Returning(CreateSymbolicNameByString(variables...)...)
}

func (s StandaloneCallBuilder) ReturningByNamed(variables ...Named) OngoingReadingAndReturn {
	return s.Returning(CreateSymbolicNameByNamed(variables...)...)
}

func (s StandaloneCallBuilder) Returning(expression ...Expression) OngoingReadingAndReturn {
	return DefaultStatementBuilderCreate1(s).Returning(expression...)
}

func (s StandaloneCallBuilder) ReturningDistinctByString(variables ...string) OngoingReadingAndReturn {
	return s.ReturningDistinct(CreateSymbolicNameByString(variables...)...)
}

func (s StandaloneCallBuilder) ReturningDistinctByNamed(variables ...Named) OngoingReadingAndReturn {
	return s.ReturningDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (s StandaloneCallBuilder) ReturningDistinct(expression ...Expression) OngoingReadingAndReturn {
	return DefaultStatementBuilderCreate1(s).ReturningDistinct(expression...)
}

func (s StandaloneCallBuilder) Build() (Statement, error) {
	if s.procedureName.GetError() != nil {
		return nil, s.procedureName.GetError()
	}
	if s.yieldItems.GetError() != nil {
		return nil, s.yieldItems.GetError()
	}
	argumentsList := Arguments{}
	if s.arguments == nil || len(s.arguments) > 0 {
		argumentsList = ArgumentsCreate(s.arguments)
	}
	if argumentsList.GetError() != nil {
		return nil, argumentsList.GetError()
	}
	condition := s.conditionBuilder.buildCondition()
	if condition != nil && condition.isNotNil() {
		if condition.GetError() != nil {
			return nil, condition.GetError()
		}
		return ProcedureCallCreate(s.procedureName, argumentsList, s.yieldItems,
			WhereCreate(condition)), nil
	}
	return ProcedureCallCreate(s.procedureName, argumentsList, s.yieldItems, Where{}), nil
}

func (s StandaloneCallBuilder) isNotNil() bool {
	return s.notNil
}

func (s StandaloneCallBuilder) AsFunction() Expression {
	return FunctionInvocationCreate(FunctionDefinitionDefault{s.procedureName.getQualifiedName()}, s.arguments...)
}

func (s StandaloneCallBuilder) withArgs(arguments ...Expression) OngoingStandaloneCallWithoutArguments {
	s.arguments = arguments
	return s
}

func (s StandaloneCallBuilder) YieldSymbolic(name ...SymbolicName) OngoingStandaloneCallWithReturnFields {
	expressions := make([]Expression, len(name))
	for i := range name {
		expressions[i] = name[i]
	}
	s.yieldItems = yieldAllOf(expressions...)
	return s
}

func (s StandaloneCallBuilder) YieldString(yieldedItems ...string) OngoingStandaloneCallWithReturnFields {
	names := make([]SymbolicName, len(yieldedItems))
	for i := range yieldedItems {
		names[i] = SymbolicNameCreate(yieldedItems[i])
	}
	return s.YieldSymbolic(names...)
}

func (s StandaloneCallBuilder) Yield(aliasedResultFields ...AliasedExpression) OngoingStandaloneCallWithReturnFields {
	expressions := make([]Expression, len(aliasedResultFields))
	for i := range aliasedResultFields {
		expressions[i] = aliasedResultFields[i]
	}
	s.yieldItems = yieldAllOf(expressions...)
	return s
}

func (s StandaloneCallBuilder) WithByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return s.With(CreateSymbolicNameByString(variables...)...)
}

func (s StandaloneCallBuilder) WithByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return s.With(CreateSymbolicNameByNamed(variables...)...)
}

func (s StandaloneCallBuilder) With(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	return DefaultStatementBuilderCreate1(s).With(expressions...)
}

func (s StandaloneCallBuilder) WithDistinctByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return s.WithDistinct(CreateSymbolicNameByString(variables...)...)
}

func (s StandaloneCallBuilder) WithDistinctByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return s.WithDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (s StandaloneCallBuilder) WithDistinct(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	return DefaultStatementBuilderCreate1(s).WithDistinct(expressions...)
}

func (s StandaloneCallBuilder) Call(statement Statement) OngoingReadingWithoutWhere {
	return DefaultStatementBuilderCreate1(s).Call(statement)
}
