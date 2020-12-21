package cypher

type UpdatingClauseBuilder interface {
	build() UpdatingClause
}

//Implement

//Merge
type MergeBuilder struct {
	patternElements []PatternElement
	mergeActions    []MergeAction
}

func MergeBuilderCreate(patternElements []PatternElement) MergeBuilder {
	return MergeBuilder{
		patternElements: patternElements,
		mergeActions:    make([]MergeAction, 0),
	}
}

func (c MergeBuilder) on(mergeType MERGE_TYPE, expressions ...Expression) (SupportsActionsOnTheUpdatingClause, error) {
	expression, err := prepareSetExpression(expressions)
	if err != nil {
		return MergeBuilder{}, err
	}
	expressionList := ExpressionListCreate(expression)
	c.mergeActions = append(c.mergeActions, MergeActionCreate(mergeType, SetCreate(expressionList)))
	return c, nil
}

func (c MergeBuilder) build() UpdatingClause {
	return MergeCreate1(PatternCreate(c.patternElements), c.mergeActions)
}

//Create
type CreateBuilder struct {
	patternElements []PatternElement
}

func CreateBuilderCreate(patternElements []PatternElement) CreateBuilder {
	return CreateBuilder{patternElements: patternElements}
}

func (c CreateBuilder) build() UpdatingClause {
	return CreateCreate(PatternCreate(c.patternElements))
}

//Delete
type DeleteBuilder struct {
	delete Delete
}

func DeleteBuilderCreate(expressionList ExpressionList, detach bool) DeleteBuilder {
	return DeleteBuilder{delete: DeleteCreate(expressionList, detach)}
}

func (c DeleteBuilder) build() UpdatingClause {
	return c.delete
}

//Set
type SetBuilder struct {
	set Set
}

func SetBuilderCreate(expressionList ExpressionList) SetBuilder {
	return SetBuilder{set: SetCreate(expressionList)}
}

func (s SetBuilder) build() UpdatingClause {
	return s.set
}

//Remove
type RemoveBuilder struct {
	remove Remove
}

func RemoveBuilderCreate(expressionList ExpressionList) RemoveBuilder {
	return RemoveBuilder{remove: RemoveCreate(expressionList)}
}

func (r RemoveBuilder) build() UpdatingClause {
	return r.remove
}
