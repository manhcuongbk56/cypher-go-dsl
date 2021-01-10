package cypher

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

func DefaultStatementBuilderCreate1(currentOngoingCall ProcedureCallBuilder) DefaultStatementBuilder {
	return DefaultStatementBuilder{
		currentOngoingCall:        currentOngoingCall,
		currentSinglePartElements: make([]Visitable, 0),
		multipartElements:         make([]MultiPartElement, 0),
	}
}

func DefaultStatementBuilderError(err error) DefaultStatementBuilder {
	return DefaultStatementBuilder{
		err: err,
	}
}

func (d DefaultStatementBuilder) Where(condition Condition) OngoingReadingWithWhere {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	if condition != nil && condition.GetError() != nil {
		return DefaultStatementBuilderError(condition.GetError())
	}
	d.currentOngoingMatch.conditionBuilder.Where(condition)
	return d
}

func (d DefaultStatementBuilder) WhereConditionContainer(container ConditionContainer) OngoingReadingWithWhere {
	return d.Where(container.Get())
}

func (d DefaultStatementBuilder) AddWith(with With) DefaultStatementBuilder {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	if with.GetError() != nil {
		return DefaultStatementBuilderError(with.GetError())
	}
	if with.isNotNil() {
		d.multipartElements = append(d.multipartElements, MultiPartElementCreate(d.BuildListOfVisitable(true), with))
	}
	return d
}

func (d DefaultStatementBuilder) WherePattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	if pattern != nil && pattern.GetError() != nil {
		return DefaultStatementBuilderError(pattern.GetError())
	}
	return d.Where(RelationshipPatternConditionCreate(pattern))
}

func (d DefaultStatementBuilder) ReturningByString(variables ...string) OngoingReadingAndReturn {
	return d.Returning(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) ReturningByNamed(variables ...Named) OngoingReadingAndReturn {
	if d.err != nil {
		return DefaultStatementWithReturnBuilderError(d.err)
	}
	for _, variable := range variables {
		if variable != nil && variable.GetError() != nil {
			return DefaultStatementWithReturnBuilderError(variable.GetError())
		}
	}
	return d.Returning(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) ReturningDistinctByString(variables ...string) OngoingReadingAndReturn {
	if d.err != nil {
		return DefaultStatementWithReturnBuilderError(d.err)
	}
	return d.ReturningDistinct(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) ReturningDistinctByNamed(variables ...Named) OngoingReadingAndReturn {
	if d.err != nil {
		return DefaultStatementWithReturnBuilderError(d.err)
	}
	for _, variable := range variables {
		if variable != nil && variable.GetError() != nil {
			return DefaultStatementWithReturnBuilderError(variable.GetError())
		}
	}
	return d.ReturningDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) WithByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	if d.err != nil {
		return DefaultStatementWithWithBuilderError(d.err)
	}
	return d.With(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) WithByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	if d.err != nil {
		return DefaultStatementWithWithBuilderError(d.err)
	}
	for _, variable := range variables {
		if variable != nil && variable.GetError() != nil {
			return DefaultStatementWithWithBuilderError(variable.GetError())
		}
	}
	return d.With(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) With(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	if d.err != nil {
		return DefaultStatementWithWithBuilderError(d.err)
	}
	for _, expression := range expressions {
		if expression != nil && expression.GetError() != nil {
			return DefaultStatementWithWithBuilderError(expression.GetError())
		}
	}
	return d.WithDefault(false, expressions...)
}

func (d DefaultStatementBuilder) WithDistinctByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	if d.err != nil {
		return DefaultStatementWithWithBuilderError(d.err)
	}
	return d.WithDistinct(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) WithDistinctByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	if d.err != nil {
		return DefaultStatementWithWithBuilderError(d.err)
	}
	for _, variable := range variables {
		if variable != nil && variable.GetError() != nil {
			return DefaultStatementWithWithBuilderError(variable.GetError())
		}
	}
	return d.WithDistinct(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) WithDistinct(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	if d.err != nil {
		return DefaultStatementWithWithBuilderError(d.err)
	}
	for _, expression := range expressions {
		if expression != nil && expression.GetError() != nil {
			return DefaultStatementWithWithBuilderError(expression.GetError())
		}
	}
	return d.WithDefault(true, expressions...)
}

func (d DefaultStatementBuilder) WithDefault(distinct bool, expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	if d.err != nil {
		return DefaultStatementWithWithBuilderError(d.err)
	}
	for _, expression := range expressions {
		if expression != nil && expression.GetError() != nil {
			return DefaultStatementWithWithBuilderError(expression.GetError())
		}
	}
	ongoingMatchAndWith := DefaultStatementWithWithBuilderCreate(&d, distinct)
	ongoingMatchAndWith.addExpression(expressions...)
	return ongoingMatchAndWith
}

func (d DefaultStatementBuilder) DeleteByString(variables ...string) OngoingUpdate {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	return d.Delete(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) DeleteByNamed(variables ...Named) OngoingUpdate {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	for _, variable := range variables {
		if variable != nil && variable.GetError() != nil {
			return DefaultStatementBuilderError(variable.GetError())
		}
	}
	return d.Delete(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) Delete(expressions ...Expression) OngoingUpdate {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	for _, expression := range expressions {
		if expression != nil && expression.GetError() != nil {
			return DefaultStatementBuilderError(expression.GetError())
		}
	}
	return d.Update(UPDATE_TYPE_DELETE, ExpressionsToVisitables(expressions))
}

func (d DefaultStatementBuilder) DetachDeleteByString(variables ...string) OngoingUpdate {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	return d.DetachDelete(CreateSymbolicNameByString(variables...)...)
}

func (d DefaultStatementBuilder) DetachDeleteByNamed(variables ...Named) OngoingUpdate {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	for _, variable := range variables {
		if variable != nil && variable.GetError() != nil {
			return DefaultStatementBuilderError(variable.GetError())
		}
	}
	return d.DetachDelete(CreateSymbolicNameByNamed(variables...)...)
}

func (d DefaultStatementBuilder) DetachDelete(expressions ...Expression) OngoingUpdate {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	for _, expression := range expressions {
		if expression != nil && expression.GetError() != nil {
			return DefaultStatementBuilderError(expression.GetError())
		}
	}
	return d.Update(UPDATE_TYPE_DETACH_DELETE, ExpressionsToVisitables(expressions))
}

func (d DefaultStatementBuilder) Merge(patterns ...PatternElement) OngoingUpdateAndExposesSet {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	for _, pattern := range patterns {
		if pattern != nil && pattern.GetError() != nil {
			return DefaultStatementBuilderError(pattern.GetError())
		}
	}
	return d.Update(UPDATE_TYPE_MERGE, PatternElementsToVisitables(patterns))
}

func (d DefaultStatementBuilder) Set(expressions ...Expression) BuildableStatementAndOngoingMatchAndUpdate {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	for _, expression := range expressions {
		if expression != nil && expression.GetError() != nil {
			return DefaultStatementBuilderError(expression.GetError())
		}
	}
	d.CloseCurrentOngoingUpdate()
	return DefaultStatementWithUpdateBuilderCreate2(&d, UPDATE_TYPE_SET, expressions...)
}

func (d DefaultStatementBuilder) SetWithNamed(variable Named, expression Expression) BuildableStatementAndOngoingMatchAndUpdate {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	if variable != nil && variable.GetError() != nil {
		return DefaultStatementBuilderError(variable.GetError())
	}
	if expression != nil && expression.GetError() != nil {
		return DefaultStatementBuilderError(expression.GetError())
	}
	return d.Set(variable.GetSymbolicName(), expression)
}

func (d DefaultStatementBuilder) SetByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	if node.GetError() != nil {
		return DefaultStatementBuilderError(node.GetError())
	}
	return DefaultStatementWithUpdateBuilderCreate2(&d, UPDATE_TYPE_SET, OperationSetLabel(node, labels...))
}

func (d DefaultStatementBuilder) RemoveByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	if node.GetError() != nil {
		return DefaultStatementBuilderError(node.GetError())
	}
	return DefaultStatementWithUpdateBuilderCreate2(&d, UPDATE_TYPE_REMOVE, OperationRemove(node, labels...))
}

func (d DefaultStatementBuilder) Remove(properties ...Property) BuildableStatementAndOngoingMatchAndUpdate {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	for _, property := range properties {
		if property.GetError() != nil {
			return DefaultStatementBuilderError(property.GetError())
		}
	}
	expressions := make([]Expression, len(properties))
	for i := range properties {
		expressions[i] = properties[i]
	}
	return DefaultStatementWithUpdateBuilderCreate2(&d, UPDATE_TYPE_REMOVE, expressions...)
}

func (d DefaultStatementBuilder) Unwinds(expressions ...Expression) OngoingUnwind {
	if d.err != nil {
		return DefaultOngoingUnwindError(d.err)
	}
	for _, expression := range expressions {
		if expression != nil && expression.GetError() != nil {
			return DefaultOngoingUnwindError(expression.GetError())
		}
	}
	return d.Unwind(ListOf(expressions...))
}

func (d DefaultStatementBuilder) UnwindByString(variable string) OngoingUnwind {
	if d.err != nil {
		return DefaultOngoingUnwindError(d.err)
	}
	return d.Unwind(ASymbolic(variable))
}

func (d DefaultStatementBuilder) Unwind(expression Expression) OngoingUnwind {
	if d.err != nil {
		return DefaultOngoingUnwindError(d.err)
	}
	if expression != nil && expression.GetError() != nil {
		return DefaultOngoingUnwindError(expression.GetError())
	}
	d.CloseCurrentOngoingMatch()
	return DefaultOngoingUnwindCreate(&d, expression)
}

func (d DefaultStatementBuilder) Call(statement Statement) OngoingReadingWithoutWhere {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	if statement != nil && statement.GetError() != nil {
		return DefaultStatementBuilderError(statement.GetError())
	}
	if d.err != nil {
		return d
	}
	d.CloseCurrentOngoingCall()
	d.CloseCurrentOngoingMatch()
	d.CloseCurrentOngoingUpdate()
	singlePart := SubqueryCall(statement)
	if singlePart.GetError() != nil {
		d.err = singlePart.GetError()
		return d
	}
	d.currentSinglePartElements = append(d.currentSinglePartElements, singlePart)
	return d
}

func (d DefaultStatementBuilder) Call1(namespaceAndProcedure ...string) OngoingInQueryCallWithoutArguments {
	if d.err != nil {
		return InQueryCallBuilder{err: d.err}
	}
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
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	d.currentOngoingMatch.conditionBuilder.And(condition)
	return d
}

func (d DefaultStatementBuilder) AndPattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	if pattern != nil && pattern.GetError() != nil {
		return DefaultStatementBuilderError(pattern.GetError())
	}
	return d.And(RelationshipPatternConditionCreate(pattern))
}

func (d DefaultStatementBuilder) Or(condition Condition) OngoingReadingWithWhere {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	if condition != nil && condition.GetError() != nil {
		return DefaultStatementBuilderError(condition.GetError())
	}
	d.currentOngoingMatch.conditionBuilder.Or(condition)
	return d
}

func (d DefaultStatementBuilder) OrPattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	if pattern != nil && pattern.GetError() != nil {
		return DefaultStatementBuilderError(pattern.GetError())
	}
	return d.Or(RelationshipPatternConditionCreate(pattern))
}

func (d DefaultStatementBuilder) OptionalMatch(patterns ...PatternElement) OngoingReadingWithoutWhere {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	for _, pattern := range patterns {
		if pattern != nil && pattern.GetError() != nil {
			return DefaultStatementBuilderError(pattern.GetError())
		}
	}
	return d.MatchDefault(true, patterns...)
}

func (d DefaultStatementBuilder) Create(elements ...PatternElement) OngoingUpdateAndExposesSet {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	for _, element := range elements {
		if element != nil && element.GetError() != nil {
			return DefaultStatementBuilderError(element.GetError())
		}
	}
	return d.Update(UPDATE_TYPE_CREATE, PatternElementsToVisitables(elements))
}

func (d DefaultStatementBuilder) OnCreate() OngoingMergeAction {
	return d.OngoingOnAfterMerge(ON_CREATE)
}

func (d DefaultStatementBuilder) OnMatch() OngoingMergeAction {
	return d.OngoingOnAfterMerge(ON_MATCH)
}

func (d *DefaultStatementBuilder) OngoingOnAfterMerge(mergeType MERGE_TYPE) OngoingMergeAction {
	if d.err != nil {
		return OngoingMergeActionInBuilderError(d.err)
	}
	if _, isSupports := d.currentOngoingUpdate.builder.(SupportsActionsOnTheUpdatingClause); !isSupports ||
		!d.currentOngoingUpdate.notNil {
		return OngoingMergeActionInBuilderError(errors.New("merge must have been invoked before defining an event"))
	}
	return OngoingMergeActionInBuilderCreate(d, mergeType)
}

func (d DefaultStatementBuilder) Match(patterns ...PatternElement) OngoingReadingWithoutWhere {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	for _, pattern := range patterns {
		if pattern != nil && pattern.GetError() != nil {
			return DefaultStatementBuilderError(pattern.GetError())
		}
	}
	return d.MatchDefault(false, patterns...)
}

func (d DefaultStatementBuilder) MatchDefault(optional bool, patterns ...PatternElement) OngoingReadingWithoutWhere {
	if d.err != nil {
		return DefaultStatementBuilderError(d.err)
	}
	for _, pattern := range patterns {
		if pattern != nil && pattern.GetError() != nil {
			return DefaultStatementBuilderError(pattern.GetError())
		}
	}
	d.CloseCurrentOngoingMatch()
	d.CloseCurrentOngoingCall()
	d.currentOngoingMatch = MatchBuilderCreate(optional)
	d.currentOngoingMatch.patternList = append(d.currentOngoingMatch.patternList, patterns...)
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

func (d DefaultStatementBuilder) Returning(expressions ...Expression) OngoingReadingAndReturn {
	return d.ReturningDefault(false, expressions...)
}

func (d DefaultStatementBuilder) ReturningDistinct(expression ...Expression) OngoingReadingAndReturn {
	return d.ReturningDefault(true, expression...)
}

func (d DefaultStatementBuilder) ReturningDefault(distinct bool, expressions ...Expression) OngoingReadingAndReturn {
	if d.err != nil {
		return DefaultStatementWithReturnBuilderError(d.err)
	}
	for _, expression := range expressions {
		if expression != nil && expression.GetError() != nil {
			return DefaultStatementWithReturnBuilderError(expression.GetError())
		}
	}
	withReturnBuilder := DefaultStatementWithReturnBuilder{
		distinct:       distinct,
		defaultBuilder: &d,
	}
	withReturnBuilder.AddExpression(expressions...)
	return withReturnBuilder
}

func (d DefaultStatementBuilder) Build() (Statement, error) {
	if d.err != nil {
		return nil, d.err
	}
	return d.BuildImpl(false, Return{}), nil
}

func (d DefaultStatementBuilder) BuildImpl(clearCurrentBuildSteps bool, returning Return) Statement {
	singlePartQuery := SinglePartQueryCreate(d.BuildListOfVisitable(clearCurrentBuildSteps), returning)
	if len(d.multipartElements) == 0 {
		return singlePartQuery
	}
	return MultiPartQueryCreate(d.multipartElements, singlePartQuery)
}

func (d *DefaultStatementBuilder) BuildListOfVisitable(clearAfter bool) []Visitable {
	visitables := make([]Visitable, 0)
	visitables = append(visitables, d.currentSinglePartElements...)
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
		return nil, xerrors.New("the list of expression to OperationSet must be even")
	}
	for i := 0; i < len(listOfExpressions); i += 2 {
		propertyOperations = append(propertyOperations, OperationSet(listOfExpressions[i], listOfExpressions[i+1]))
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
	if defaultBuilder.err != nil {
		return OngoingMergeActionInBuilderError(defaultBuilder.err)
	}
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
	if o.err != nil {
		return DefaultStatementBuilderError(o.err)
	}
	for _, expression := range expressions {
		if expression != nil && expression.GetError() != nil {
			return DefaultStatementBuilderError(expression.GetError())
		}
	}
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
	if o.err != nil {
		return DefaultStatementBuilderError(o.err)
	}
	if variable != nil && variable.GetError() != nil {
		return DefaultStatementBuilderError(variable.GetError())
	}
	if expression != nil && expression.GetError() != nil {
		return DefaultStatementBuilderError(expression.GetError())
	}
	return o.Set(variable.GetSymbolicName(), expression)
}
