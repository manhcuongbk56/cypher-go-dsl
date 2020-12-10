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

func (d DefaultStatementBuilder) where(condition Condition) OngoingReadingWithWhere {
	d.currentOngoingMatch.conditionBuilder.Where(condition)
	return d
}

func (d DefaultStatementBuilder) addWith(with With) DefaultStatementBuilder {
	if with.isNotNil() {
		d.multipartElements = append(d.multipartElements, MultiPartElementCreate(d.BuildListOfVisitable(true), with))
	}
	return d
}

func (d DefaultStatementBuilder) wherePattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	return d.where(RelationshipPatternConditionCreate(pattern))
}

func (d DefaultStatementBuilder) returningByString(variables ...string) OngoingReadingAndReturn {
	return d.returning(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) returningByNamed(variables ...Named) OngoingReadingAndReturn {
	return d.returning(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) returningDistinctByString(variables ...string) OngoingReadingAndReturn {
	return d.returningDistinct(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) returningDistinctByNamed(variables ...Named) OngoingReadingAndReturn {
	return d.returningDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) withByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return d.with(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) withByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return d.with(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) with(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	return d.withDefault(false, expressions...)
}

func (d DefaultStatementBuilder) withDistinctByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return d.withDistinct(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) withDistinctByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return d.withDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) withDistinct(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	return d.withDefault(true, expressions...)
}

func (d DefaultStatementBuilder) withDefault(distinct bool, expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	ongoingMatchAndWith := DefaultStatementWithWithBuilderCreate(&d, distinct)
	ongoingMatchAndWith.addExpression(expressions...)
	return ongoingMatchAndWith
}

func (d DefaultStatementBuilder) deleteByString(variables ...string) OngoingUpdate {
	return d.delete(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) deleteByNamed(variables ...Named) OngoingUpdate {
	return d.delete(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) delete(expressions ...Expression) OngoingUpdate {
	return d.update(UPDATE_TYPE_DELETE, ExpressionsToVisitables(expressions))
}

func (d DefaultStatementBuilder) detachDeleteByString(variables ...string) OngoingUpdate {
	return d.detachDelete(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) detachDeleteByNamed(variables ...Named) OngoingUpdate {
	return d.detachDelete(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) detachDelete(expressions ...Expression) OngoingUpdate {
	return d.update(UPDATE_TYPE_DETACH_DELETE, ExpressionsToVisitables(expressions))
}

func (d DefaultStatementBuilder) merge(pattern ...PatternElement) OngoingUpdate {
	return d.update(UPDATE_TYPE_MERGE, PatternElementsToVisitables(pattern))
}

func (d DefaultStatementBuilder) set(expressions ...Expression) BuildableStatementAndOngoingMatchAndUpdate {
	d.closeCurrentOngoingUpdate()
	return DefaultStatementWithUpdateBuilderCreate2(&d, UPDATE_TYPE_SET, expressions...)
}

func (d DefaultStatementBuilder) setWithNamed(variable Named, expression Expression) BuildableStatementAndOngoingMatchAndUpdate {
	return d.set(variable.getSymbolicName(), expression)
}

func (d DefaultStatementBuilder) setByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	return DefaultStatementWithUpdateBuilderCreate2(&d, UPDATE_TYPE_SET, set1(node, labels...))
}

func (d DefaultStatementBuilder) removeByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	return DefaultStatementWithUpdateBuilderCreate2(&d, UPDATE_TYPE_REMOVE, remove(node, labels...))
}

func (d DefaultStatementBuilder) remove(properties ...Property) BuildableStatementAndOngoingMatchAndUpdate {
	expressions := make([]Expression, len(properties))
	for i := range properties {
		expressions[i] = properties[i]
	}
	return DefaultStatementWithUpdateBuilderCreate2(&d, UPDATE_TYPE_REMOVE, expressions...)
}

func (d DefaultStatementBuilder) unwinds(expression ...Expression) OngoingUnwind {
	return d.unwind(ListOf(expression...))
}

func (d DefaultStatementBuilder) unwindByString(variable string) OngoingUnwind {
	return d.unwind(Name(variable))
}

func (d DefaultStatementBuilder) unwind(expression Expression) OngoingUnwind {
	d.closeCurrentOngoingMatch()
	return DefaultOngoingUnwindCreate(&d, expression)
}

func (d DefaultStatementBuilder) call(statement Statement) OngoingReadingWithoutWhere {
	d.closeCurrentOngoingCall()
	d.closeCurrentOngoingMatch()
	d.closeCurrentOngoingUpdate()
	singlePart, err := SubqueryCall(statement)
	if err != nil {
		d.err = err
		return d
	}
	d.currentSinglePartElements = append(d.currentSinglePartElements, singlePart)
	return d
}

func (d DefaultStatementBuilder) call1(namespaceAndProcedure ...string) OngoingInQueryCallWithoutArguments {
	d.closeCurrentOngoingMatch()
	d.closeCurrentOngoingCall()
	inQueryCallBuilder := InQueryCallBuilderCreate(&d, ProcedureNameCreate(namespaceAndProcedure...))
	d.currentOngoingCall = inQueryCallBuilder
	return inQueryCallBuilder
}

func (d DefaultStatementBuilder) asCondition() Condition {
	if !d.currentOngoingMatch.notNil || len(d.currentSinglePartElements) > 0 {
		panic("only simple MATCH statements can be used as existential subqueries")
	}
	return ExistentialSubqueryExists(d.currentOngoingMatch.buildMatch())
}

func (d DefaultStatementBuilder) and(condition Condition) OngoingReadingWithWhere {
	d.currentOngoingMatch.conditionBuilder.And(condition)
	return d
}

func (d DefaultStatementBuilder) andPattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	return d.and(RelationshipPatternConditionCreate(pattern))
}

func (d DefaultStatementBuilder) or(condition Condition) OngoingReadingWithWhere {
	d.currentOngoingMatch.conditionBuilder.Or(condition)
	return d
}

func (d DefaultStatementBuilder) orPattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	return d.or(RelationshipPatternConditionCreate(pattern))
}

func (d DefaultStatementBuilder) optionalMatch(pattern ...PatternElement) OngoingReadingWithoutWhere {
	return d.MatchDefault(true, pattern...)
}

func (d DefaultStatementBuilder) create(element ...PatternElement) OngoingUpdate {
	return d.update(UPDATE_TYPE_CREATE, PatternElementsToVisitables(element))
}

func (d DefaultStatementBuilder) onCreate() OngoingMergeAction {
	return d.ongoingOnAfterMerge(ON_CREATE)
}

func (d DefaultStatementBuilder) onMatch() OngoingMergeAction {
	return d.ongoingOnAfterMerge(ON_MATCH)
}

func (d *DefaultStatementBuilder) ongoingOnAfterMerge(mergeType MERGE_TYPE) OngoingMergeAction {
	if _, isSupports := d.currentOngoingUpdate.builder.(SupportsActionsOnTheUpdatingClause); isSupports ||
		!d.currentOngoingUpdate.notNil {
		return OngoingMergeActionInBuilderError(errors.New("merge must have been invoked before defining an event"))
	}
	return OngoingMergeActionInBuilderCreate(d, mergeType)
}

func (d DefaultStatementBuilder) match(pattern ...PatternElement) OngoingReadingWithoutWhere {
	return d.MatchDefault(false, pattern...)
}

func (d DefaultStatementBuilder) MatchDefault(optional bool, pattern ...PatternElement) OngoingReadingWithoutWhere {
	d.closeCurrentOngoingMatch()
	d.closeCurrentOngoingCall()
	d.currentOngoingMatch = MatchBuilderCreate(optional)
	d.currentOngoingMatch.patternList = append(d.currentOngoingMatch.patternList, pattern...)
	return d
}

func (d *DefaultStatementBuilder) closeCurrentOngoingMatch() {
	if !d.currentOngoingMatch.notNil {
		return
	}
	d.currentSinglePartElements = append(d.currentSinglePartElements, d.currentOngoingMatch.buildMatch())
	d.currentOngoingMatch = MatchBuilder{}
}

func (d *DefaultStatementBuilder) closeCurrentOngoingCall() {
	if d.currentOngoingCall == nil || !d.currentOngoingCall.isNotNil() {
		return
	}
	builtCall, err := d.currentOngoingCall.build()
	if err != nil {
		d.err = err
		return
	}
	d.currentSinglePartElements = append(d.currentSinglePartElements, builtCall)
	d.currentOngoingCall = InQueryCallBuilder{}
}

func (d *DefaultStatementBuilder) closeCurrentOngoingUpdate() {
	if !d.currentOngoingUpdate.notNil {
		return
	}
	d.currentSinglePartElements = append(d.currentSinglePartElements, d.currentOngoingUpdate.builder.build())
	d.currentOngoingUpdate = DefaultStatementWithUpdateBuilder{}
}

func (d DefaultStatementBuilder) returning(expression ...Expression) OngoingReadingAndReturn {
	return d.returningDefault(false, expression...)
}

func (d DefaultStatementBuilder) returningDistinct(expression ...Expression) OngoingReadingAndReturn {
	return d.returningDefault(true, expression...)
}

func (d DefaultStatementBuilder) returningDefault(distinct bool, expression ...Expression) OngoingReadingAndReturn {
	withReturnBuilder := DefaultStatementWithReturnBuilder{
		distinct:       distinct,
		defaultBuilder: &d,
	}
	withReturnBuilder.AddExpression(expression...)
	return withReturnBuilder
}

func (d DefaultStatementBuilder) build() (Statement, error) {
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
		builtCall, err := d.currentOngoingCall.build()
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
	d.closeCurrentOngoingMatch()
	d.currentSinglePartElements = append(d.currentSinglePartElements, clause)
	return *d
}

func (d *DefaultStatementBuilder) update(updateType UpdateType, pattern []Visitable) DefaultStatementBuilder {
	d.closeCurrentOngoingMatch()
	d.closeCurrentOngoingCall()
	d.closeCurrentOngoingUpdate()
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

func (o OngoingMergeActionInBuilder) getErr() error {
	return o.err
}

func (o OngoingMergeActionInBuilder) set(expressions ...Expression) OngoingMatchAndUpdateAndBuildableStatementAndExposesMergeAction {
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

func (o OngoingMergeActionInBuilder) setWithNamed(variable Named, expression Expression) OngoingMatchAndUpdateAndBuildableStatementAndExposesMergeAction {
	return o.set(variable.getSymbolicName(), expression)
}
