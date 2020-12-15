package cypher_go_dsl

import (
	"errors"
	"golang.org/x/xerrors"
)

type DefaultStatementBuilder struct {
	err                       error
	currentSinglePartElements []Visitable
	multipartElements         []MultiPartElement
	currentOngoingMatch       MatchBuilder
	currentOngoingUpdate      DefaultStatementWithUpdateBuilder
	currentOngoingCall        ProcedureCallBuilder
}

func DefaultStatementBuilderCreate() DefaultStatementBuilder {
	return DefaultStatementBuilder{
		currentSinglePartElements: make([]Visitable, 0),
		multipartElements:         make([]MultiPartElement, 0),
	}
}

func (d DefaultStatementBuilder) Where(condition Condition) OngoingReadingWithWhere {
	d.currentOngoingMatch.conditionBuilder.Where(condition)
	return d
}

func (d DefaultStatementBuilder) AddWith(with With) DefaultStatementBuilder {
	if with.isNotNil() {
		d.multipartElements = append(d.multipartElements, MultiPartElementCreate(d.BuildListOfVisitable(true), with))
	}
	return d
}

func (d DefaultStatementBuilder) WherePattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	return d.Where(RelationshipPatternConditionCreate(pattern))
}

func (d DefaultStatementBuilder) ReturningByString(variables ...string) OngoingReadingAndReturn {
	return d.Returning(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) ReturningByNamed(variables ...Named) OngoingReadingAndReturn {
	return d.Returning(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) ReturningDistinctByString(variables ...string) OngoingReadingAndReturn {
	return d.ReturningDistinct(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) ReturningDistinctByNamed(variables ...Named) OngoingReadingAndReturn {
	return d.ReturningDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) WithByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return d.With(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) WithByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return d.With(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) With(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	return d.WithDefault(false, expressions...)
}

func (d DefaultStatementBuilder) WithDistinctByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return d.WithDistinct(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) WithDistinctByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return d.WithDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) WithDistinct(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	return d.WithDefault(true, expressions...)
}

func (d DefaultStatementBuilder) WithDefault(distinct bool, expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	ongoingMatchAndWith := DefaultStatementWithWithBuilderCreate(&d, distinct)
	ongoingMatchAndWith.addExpression(expressions...)
	return ongoingMatchAndWith
}

func (d DefaultStatementBuilder) DeleteByString(variables ...string) OngoingUpdate {
	return d.Delete(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) DeleteByNamed(variables ...Named) OngoingUpdate {
	return d.Delete(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) Delete(expressions ...Expression) OngoingUpdate {
	return d.Update(UPDATE_TYPE_DELETE, ExpressionsToVisitables(expressions))
}

func (d DefaultStatementBuilder) DetachDeleteByString(variables ...string) OngoingUpdate {
	return d.DetachDelete(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) DetachDeleteByNamed(variables ...Named) OngoingUpdate {
	return d.DetachDelete(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) DetachDelete(expressions ...Expression) OngoingUpdate {
	return d.Update(UPDATE_TYPE_DETACH_DELETE, ExpressionsToVisitables(expressions))
}

func (d DefaultStatementBuilder) Merge(pattern ...PatternElement) OngoingUpdate {
	return d.Update(UPDATE_TYPE_MERGE, PatternElementsToVisitables(pattern))
}

func (d DefaultStatementBuilder) Set(expressions ...Expression) BuildableStatementAndOngoingMatchAndUpdate {
	d.CloseCurrentOngoingUpdate()
	return DefaultStatementWithUpdateBuilderCreate2(&d, UPDATE_TYPE_SET, expressions...)
}

func (d DefaultStatementBuilder) SetWithNamed(variable Named, expression Expression) BuildableStatementAndOngoingMatchAndUpdate {
	return d.Set(variable.getSymbolicName(), expression)
}

func (d DefaultStatementBuilder) SetByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	return DefaultStatementWithUpdateBuilderCreate2(&d, UPDATE_TYPE_SET, set1(node, labels...))
}

func (d DefaultStatementBuilder) RemoveByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	return DefaultStatementWithUpdateBuilderCreate2(&d, UPDATE_TYPE_REMOVE, remove(node, labels...))
}

func (d DefaultStatementBuilder) Remove(properties ...Property) BuildableStatementAndOngoingMatchAndUpdate {
	expressions := make([]Expression, len(properties))
	for i := range properties {
		expressions[i] = properties[i]
	}
	return DefaultStatementWithUpdateBuilderCreate2(&d, UPDATE_TYPE_REMOVE, expressions...)
}

func (d DefaultStatementBuilder) Unwinds(expression ...Expression) OngoingUnwind {
	return d.Unwind(ListOf(expression...))
}

func (d DefaultStatementBuilder) UnwindByString(variable string) OngoingUnwind {
	return d.Unwind(Name(variable))
}

func (d DefaultStatementBuilder) Unwind(expression Expression) OngoingUnwind {
	d.CloseCurrentOngoingMatch()
	return DefaultOngoingUnwindCreate(&d, expression)
}

func (d DefaultStatementBuilder) Call(statement Statement) OngoingReadingWithoutWhere {
	if d.err != nil {
		return d
	}
	d.CloseCurrentOngoingCall()
	d.CloseCurrentOngoingMatch()
	d.CloseCurrentOngoingUpdate()
	singlePart := SubqueryCall(statement)
	if singlePart.getError() != nil {
		d.err = singlePart.getError()
		return d
	}
	d.currentSinglePartElements = append(d.currentSinglePartElements, singlePart)
	return d
}

func (d DefaultStatementBuilder) Call1(namespaceAndProcedure ...string) OngoingInQueryCallWithoutArguments {
	d.CloseCurrentOngoingMatch()
	d.CloseCurrentOngoingCall()
	inQueryCallBuilder := InQueryCallBuilderCreate(&d, ProcedureNameCreate(namespaceAndProcedure...))
	d.currentOngoingCall = inQueryCallBuilder
	return inQueryCallBuilder
}

func (d DefaultStatementBuilder) AsCondition() Condition {
	if !d.currentOngoingMatch.notNil || len(d.currentSinglePartElements) > 0 {
		panic("only simple MATCH statements can be used as existential subqueries")
	}
	return ExistentialSubqueryExists(d.currentOngoingMatch.buildMatch())
}

func (d DefaultStatementBuilder) And(condition Condition) OngoingReadingWithWhere {
	d.currentOngoingMatch.conditionBuilder.And(condition)
	return d
}

func (d DefaultStatementBuilder) AndPattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	return d.And(RelationshipPatternConditionCreate(pattern))
}

func (d DefaultStatementBuilder) Or(condition Condition) OngoingReadingWithWhere {
	d.currentOngoingMatch.conditionBuilder.Or(condition)
	return d
}

func (d DefaultStatementBuilder) OrPattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	return d.Or(RelationshipPatternConditionCreate(pattern))
}

func (d DefaultStatementBuilder) OptionalMatch(pattern ...PatternElement) OngoingReadingWithoutWhere {
	return d.MatchDefault(true, pattern...)
}

func (d DefaultStatementBuilder) Create(element ...PatternElement) OngoingUpdate {
	return d.Update(UPDATE_TYPE_CREATE, PatternElementsToVisitables(element))
}

func (d DefaultStatementBuilder) OnCreate() OngoingMergeAction {
	return d.OngoingOnAfterMerge(ON_CREATE)
}

func (d DefaultStatementBuilder) OnMatch() OngoingMergeAction {
	return d.OngoingOnAfterMerge(ON_MATCH)
}

func (d *DefaultStatementBuilder) OngoingOnAfterMerge(mergeType MERGE_TYPE) OngoingMergeAction {
	if _, isSupports := d.currentOngoingUpdate.builder.(SupportsActionsOnTheUpdatingClause); isSupports ||
		!d.currentOngoingUpdate.notNil {
		return OngoingMergeActionInBuilderError(errors.New("merge must have been invoked before defining an event"))
	}
	return OngoingMergeActionInBuilderCreate(d, mergeType)
}

func (d DefaultStatementBuilder) Match(pattern ...PatternElement) OngoingReadingWithoutWhere {
	return d.MatchDefault(false, pattern...)
}

func (d DefaultStatementBuilder) MatchDefault(optional bool, pattern ...PatternElement) OngoingReadingWithoutWhere {
	d.CloseCurrentOngoingMatch()
	d.CloseCurrentOngoingCall()
	d.currentOngoingMatch = MatchBuilderCreate(optional)
	d.currentOngoingMatch.patternList = append(d.currentOngoingMatch.patternList, pattern...)
	return d
}

func (d *DefaultStatementBuilder) CloseCurrentOngoingMatch() {
	if !d.currentOngoingMatch.notNil {
		return
	}
	d.currentSinglePartElements = append(d.currentSinglePartElements, d.currentOngoingMatch.buildMatch())
	d.currentOngoingMatch = MatchBuilder{}
}

func (d *DefaultStatementBuilder) CloseCurrentOngoingCall() {
	if d.currentOngoingCall == nil || !d.currentOngoingCall.isNotNil() {
		return
	}
	builtCall, err := d.currentOngoingCall.Build()
	if err != nil {
		d.err = err
		return
	}
	d.currentSinglePartElements = append(d.currentSinglePartElements, builtCall)
	d.currentOngoingCall = InQueryCallBuilder{}
}

func (d *DefaultStatementBuilder) CloseCurrentOngoingUpdate() {
	if !d.currentOngoingUpdate.notNil {
		return
	}
	d.currentSinglePartElements = append(d.currentSinglePartElements, d.currentOngoingUpdate.builder.build())
	d.currentOngoingUpdate = DefaultStatementWithUpdateBuilder{}
}

func (d DefaultStatementBuilder) Returning(expression ...Expression) OngoingReadingAndReturn {
	return d.ReturningDefault(false, expression...)
}

func (d DefaultStatementBuilder) ReturningDistinct(expression ...Expression) OngoingReadingAndReturn {
	return d.ReturningDefault(true, expression...)
}

func (d DefaultStatementBuilder) ReturningDefault(distinct bool, expression ...Expression) OngoingReadingAndReturn {
	withReturnBuilder := DefaultStatementWithReturnBuilder{
		distinct:       distinct,
		defaultBuilder: &d,
	}
	withReturnBuilder.AddExpression(expression...)
	return withReturnBuilder
}

func (d DefaultStatementBuilder) Build() (Statement, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.BuildImpl(false, Return{}), nil
}

func (d DefaultStatementBuilder) BuildImpl(clearCurrentBuildSteps bool, returning Return) Statement {
	singlePartQuery, _ := SinglePartQueryCreate(d.BuildListOfVisitable(clearCurrentBuildSteps), returning)
	if len(d.multipartElements) == 0 {
		return singlePartQuery
	}
	return MultiPartQueryCreate(d.multipartElements, singlePartQuery)
}

func (d *DefaultStatementBuilder) BuildListOfVisitable(clearAfter bool) []Visitable {
	visitables := make([]Visitable, 0)
	copy(visitables, d.currentSinglePartElements)
	if d.currentOngoingMatch.notNil {
		visitables = append(visitables, d.currentOngoingMatch.buildMatch())
	}
	if d.currentOngoingUpdate.isNotNil() {
		visitables = append(visitables, d.currentOngoingUpdate.builder.build())
	}
	if d.currentOngoingCall != nil && d.currentOngoingCall.isNotNil() {
		builtCall, err := d.currentOngoingCall.Build()
		if err != nil {
			d.err = err
			return nil
		}
		visitables = append(visitables, builtCall)
	}
	if clearAfter {
		d.currentOngoingMatch = MatchBuilder{}
		d.currentOngoingUpdate = DefaultStatementWithUpdateBuilder{}
		d.currentOngoingCall = InQueryCallBuilder{}
		d.currentSinglePartElements = make([]Visitable, 0)
	}
	return visitables
}

func (d *DefaultStatementBuilder) addUpdatingClause(clause UpdatingClause) DefaultStatementBuilder {
	d.CloseCurrentOngoingMatch()
	d.currentSinglePartElements = append(d.currentSinglePartElements, clause)
	return *d
}

func (d *DefaultStatementBuilder) Update(updateType UpdateType, pattern []Visitable) DefaultStatementBuilder {
	d.CloseCurrentOngoingMatch()
	d.CloseCurrentOngoingCall()
	d.CloseCurrentOngoingUpdate()
	d.currentOngoingUpdate = DefaultStatementWithUpdateBuilderCreate1(d, updateType, pattern)
	return *d
}

func getUpdatingClauseBuilder(updateType UpdateType, patternOrExpression ...Visitable) (UpdatingClauseBuilder, error) {
	mergeOrCreate := updateType == UPDATE_TYPE_MERGE || updateType == UPDATE_TYPE_CREATE
	if mergeOrCreate {
		patternElements := make([]PatternElement, 0)
		for _, visitable := range patternOrExpression {
			patternElements := append(patternElements, visitable.(PatternElement))
			if updateType == UPDATE_TYPE_CREATE {
				return CreateBuilderCreate(patternElements), nil
			} else {
				return MergeBuilderCreate(patternElements), nil
			}
		}
	} else {
		expressions := make([]Expression, 0)
		for _, visitable := range patternOrExpression {
			expressions = append(expressions, visitable.(Expression))
		}
		var expressionList ExpressionList
		if updateType == UPDATE_TYPE_SET {
			preparedExpression, err := prepareSetExpression(expressions)
			if err != nil {
				return RemoveBuilder{}, err
			}
			expressionList = ExpressionListCreate(preparedExpression)
		} else {
			expressionList = ExpressionListCreate(expressions)
		}
		switch updateType {
		case UPDATE_TYPE_DETACH_DELETE:
			return DeleteBuilderCreate(expressionList, true), nil
		case UPDATE_TYPE_DELETE:
			return DeleteBuilderCreate(expressionList, false), nil
		case UPDATE_TYPE_SET:
			return SetBuilderCreate(expressionList), nil
		case UPDATE_TYPE_REMOVE:
			return RemoveBuilderCreate(expressionList), nil
		default:
			return RemoveBuilder{}, xerrors.Errorf("unsupported update type %s", updateType)
		}
	}
	//Just return to make compiler happy
	return RemoveBuilder{}, xerrors.Errorf("unexpected behavior")
}

func prepareSetExpression(possibleSetOperations []Expression) ([]Expression, error) {
	propertyOperations := make([]Expression, 0)
	listOfExpressions := make([]Expression, 0)
	for _, possibleSetOperation := range possibleSetOperations {
		if operation, isOperation := possibleSetOperation.(Operation); isOperation {
			propertyOperations = append(propertyOperations, operation)
		} else {
			listOfExpressions = append(listOfExpressions, possibleSetOperation)
		}
	}
	if len(listOfExpressions)%2 != 0 {
		return nil, xerrors.New("the list of expression to set must be even")
	}
	for i := 0; i < len(listOfExpressions); i += 2 {
		propertyOperations = append(propertyOperations, set(listOfExpressions[i], listOfExpressions[i+1]))
	}
	return propertyOperations, nil
}

//Implement OngoingMergeAction
type OngoingMergeActionInBuilder struct {
	defaultBuilder *DefaultStatementBuilder
	mergeType      MERGE_TYPE
	err            error
}

func OngoingMergeActionInBuilderCreate(defaultBuilder *DefaultStatementBuilder, mergeType MERGE_TYPE) OngoingMergeActionInBuilder {
	return OngoingMergeActionInBuilder{
		defaultBuilder: defaultBuilder,
		mergeType:      mergeType,
	}
}

func OngoingMergeActionInBuilderError(err error) OngoingMergeActionInBuilder {
	return OngoingMergeActionInBuilder{
		err: err,
	}
}

func (o OngoingMergeActionInBuilder) GetErr() error {
	return o.err
}

func (o OngoingMergeActionInBuilder) Set(expressions ...Expression) OngoingMatchAndUpdateAndBuildableStatementAndExposesMergeAction {
	support := o.defaultBuilder.currentOngoingUpdate.builder.(SupportsActionsOnTheUpdatingClause)
	newSupport, err := support.on(o.mergeType, expressions...)
	newMergeBuilder, _ := newSupport.(MergeBuilder)
	if err != nil {
		o.defaultBuilder.err = err
		return o.defaultBuilder
	}
	o.defaultBuilder.currentOngoingUpdate.builder = newMergeBuilder
	return o.defaultBuilder
}

func (o OngoingMergeActionInBuilder) SetWithNamed(variable Named, expression Expression) OngoingMatchAndUpdateAndBuildableStatementAndExposesMergeAction {
	return o.Set(variable.getSymbolicName(), expression)
}
