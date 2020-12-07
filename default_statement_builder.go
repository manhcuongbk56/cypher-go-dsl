package cypher_go_dsl

type DefaultStatementBuilder struct {
	invalidReason             string
	currentSinglePartElements []Visitable
	multipartElements         []MultiPartElement
	currentOngoingMatch       MatchBuilder
	currentOngoingUpdate      DefaultStatementWithUpdateBuilder
	currentOngoingCall        DefaultStatementWithUpdateBuilder
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
	panic("implement me")
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
	panic("implement me")
}

func (d DefaultStatementBuilder) withByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) with(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) withDistinctByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) withDistinctByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) withDistinct(expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) withDefault(distinct bool, expressions ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) deleteByString(variables ...string) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) deleteByNamed(variables ...Named) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) delete(expressions ...Expression) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) detachDeleteByString(variables ...string) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) detachDeleteByNamed(variables ...Named) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) detachDelete(expressions ...Expression) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) merge(pattern ...PatternElement) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) set(expressions ...Expression) BuildableStatementAndOngoingMatchAndUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) setWithNamed(variable Named, expression Expression) BuildableStatementAndOngoingMatchAndUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) setByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) removeByNode(node Node, labels ...string) BuildableStatementAndOngoingMatchAndUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) remove(properties ...Property) BuildableStatementAndOngoingMatchAndUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) unwinds(expression ...Expression) OngoingUnwind {
	panic("implement me")
}

func (d DefaultStatementBuilder) unwindByString(variable string) OngoingUnwind {
	panic("implement me")
}

func (d DefaultStatementBuilder) unwind(expression Expression) OngoingUnwind {
	panic("implement me")
}

func (d DefaultStatementBuilder) call(statement Statement) OngoingReadingWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) call1(namespaceAndProcedure ...string) OngoingInQueryCallWithoutArguments {
	panic("implement me")
}

func (d DefaultStatementBuilder) asCondition() Expression {
	panic("implement me")
}

func (d DefaultStatementBuilder) and(condition Condition) OngoingReadingWithWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) andPattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) or(condition Condition) OngoingReadingWithWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) orPattern(pattern RelationshipPattern) OngoingReadingWithWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) optionalMatch(pattern ...PatternElement) OngoingReadingWithoutWhere {
	panic("implement me")
}

func (d DefaultStatementBuilder) create(element ...PatternElement) OngoingUpdate {
	panic("implement me")
}

func (d DefaultStatementBuilder) onCreate() OngoingMergeAction {
	panic("implement me")
}

func (d DefaultStatementBuilder) onMatch() OngoingMergeAction {
	panic("implement me")
}

func NewDefaultBuilder() DefaultStatementBuilder {
	return DefaultStatementBuilder{
		currentSinglePartElements: make([]Visitable, 0),
	}
}

func (d DefaultStatementBuilder) match(pattern ...PatternElement) OngoingReadingWithoutWhere {
	if pattern == nil || len(pattern) == 0 {
		return DefaultStatementBuilder{invalidReason: "patterns to match is required"}
	}
	if d.currentOngoingMatch.notNil {
		d.currentSinglePartElements = append(d.currentSinglePartElements, d.currentOngoingMatch.buildMatch())
	}
	d.currentOngoingMatch = MatchBuilder{
		optional: false,
		notNil:   true,
	}
	d.currentOngoingMatch.patternList = append(d.currentOngoingMatch.patternList, pattern...)
	return d
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
	d.currentOngoingCall = DefaultStatementWithUpdateBuilder{}
}

func (d *DefaultStatementBuilder) closeCurrentOngoingCall() {
	if !d.currentOngoingCall.notNil {
		return
	}
	d.currentSinglePartElements = append(d.currentSinglePartElements, d.currentOngoingCall.build())
	d.currentOngoingCall = DefaultStatementWithUpdateBuilder{}
}

func (d *DefaultStatementBuilder) closeCurrentOngoingUpdate() {
	if !d.currentOngoingUpdate.notNil {
		return
	}
	d.currentSinglePartElements = append(d.currentSinglePartElements, d.currentOngoingUpdate.builder.build())
	d.currentOngoingCall = DefaultStatementWithUpdateBuilder{}
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

func (d DefaultStatementBuilder) build() Statement {
	return d.BuildImpl(false, Return{})
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
	if d.currentOngoingCall.notNil {
		visitables = append(visitables, d.currentOngoingCall.build())
	}
	if clearAfter {
		d.currentOngoingMatch = MatchBuilder{}
		d.currentOngoingUpdate = DefaultStatementWithUpdateBuilder{}
		d.currentOngoingCall = DefaultStatementWithUpdateBuilder{}
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
	d.currentOngoingUpdate = DefaultStatementWithUpdateBuilderCreate1(d, updateType, pattern)
	return *d
}

func getUpdatingClauseBuilder(updateType UpdateType, patternOrExpression ...Visitable) UpdatingClauseBuilder {
	mergeOrCreate := updateType == UPDATE_TYPE_MERGE || updateType == UPDATE_TYPE_CREATE
	if mergeOrCreate {
		patternElements := make([]PatternElement, 0)
		for _, visitable := range patternOrExpression {
			patternElements := append(patternElements, visitable.(PatternElement))
			if updateType == UPDATE_TYPE_CREATE {
				return CreateBuilderCreate(patternElements)
			} else {
				return MergeBuilderCreate(patternElements)
			}
		}
	} else {
		expressions := make([]Expression, 0)
		for _, visitable := range patternOrExpression {
			expressions = append(expressions, visitable.(Expression))
		}
		var expressionList ExpressionList
		if updateType == UPDATE_TYPE_SET {
			expressionList = ExpressionListCreate(prepareSetExpression(expressions))
		} else {
			expressionList = ExpressionListCreate(expressions)
		}
		switch updateType {
		case UPDATE_TYPE_DETACH_DELETE:
			return DeleteBuilderCreate(expressionList, true)
		case UPDATE_TYPE_DELETE:
			return DeleteBuilderCreate(expressionList, false)
		case UPDATE_TYPE_SET:
			return SetBuilderCreate(expressionList)
		case UPDATE_TYPE_REMOVE:
			return RemoveBuilderCreate(expressionList)
		default:
			panic("unsupported update type")
		}
	}
	//Just return to make compiler happy
	return RemoveBuilder{}
}

func prepareSetExpression(possibleSetOperations []Expression) []Expression {
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
		panic("the list of expression to set must be even")
	}
	for i := 0; i < len(listOfExpressions); i += 2 {
		propertyOperations = append(propertyOperations, set(listOfExpressions[i], listOfExpressions[i+1]))
	}
	return propertyOperations
}
