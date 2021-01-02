package cypher

import "golang.org/x/xerrors"

func PredicateExists(property Property) Condition {
	return BooleanFunctionConditionCreate(FunctionInvocationCreate(EXISTS, property))
}

func PredicateExistsByPattern(pattern RelationshipPattern) Condition {
	return BooleanFunctionConditionCreate(FunctionInvocationCreateWithPatternElement(EXISTS, pattern))
}

func PredicateAll(variable string) OngoingListBasedPredicateFunction {
	return PredicateAllBySymbolicName(SymbolicNameCreate(variable))
}

func PredicateAllBySymbolicName(variable SymbolicName) OngoingListBasedPredicateFunction {
	return PredicateBuilderCreate(ALL, variable)
}

func PredicateAny(variable string) OngoingListBasedPredicateFunction {
	return PredicateAnyBySymbolicName(SymbolicNameCreate(variable))
}

func PredicateAnyBySymbolicName(variable SymbolicName) OngoingListBasedPredicateFunction {
	return PredicateBuilderCreate(ANY, variable)
}

func PredicateNone(variable string) OngoingListBasedPredicateFunction {
	return PredicateNoneBySymbolicName(SymbolicNameCreate(variable))
}

func PredicateNoneBySymbolicName(variable SymbolicName) OngoingListBasedPredicateFunction {
	return PredicateBuilderCreate(NONE, variable)
}

func PredicateSingle(variable string) OngoingListBasedPredicateFunction {
	return PredicateSingleBySymbolicName(SymbolicNameCreate(variable))
}

func PredicateSingleBySymbolicName(variable SymbolicName) OngoingListBasedPredicateFunction {
	return PredicateBuilderCreate(SINGLE, variable)
}

type OngoingListBasedPredicateFunction interface {
	In(list Expression) OngoingListBasedPredicateFunctionWithList
	GetError() error
}

type OngoingListBasedPredicateFunctionWithList interface {
	Where(condition Condition) Condition
	GetError() error
}

type PredicateBuilder struct {
	predicate      Predicate
	name           SymbolicName
	listExpression Expression
	err            error
}

func PredicateBuilderCreate(predicate Predicate, name SymbolicName) PredicateBuilder {
	if !name.isNotNil() {
		return PredicateBuilder{err: xerrors.New("name for predicate is require")}
	}
	return PredicateBuilder{
		predicate: predicate,
		name:      name,
	}
}

func (p PredicateBuilder) In(list Expression) OngoingListBasedPredicateFunctionWithList {
	if list == nil || !list.isNotNil() {
		return PredicateBuilder{err: xerrors.New("predicate builder: the list expression is required")}
	}
	p.listExpression = list
	return p
}

func (p PredicateBuilder) Where(condition Condition) Condition {
	if condition == nil || !condition.isNotNil() {
		return BooleanFunctionConditionError(xerrors.New("predicate builder where: the condition is required"))
	}
	return BooleanFunctionConditionCreate(FunctionInvocationCreate(p.predicate, ListPredicateCreate(p.name, p.listExpression, WhereCreate(condition))))
}

func (p PredicateBuilder) GetError() error {
	panic("implement me")
}
