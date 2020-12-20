package cypher_go_dsl

import "errors"

type PatternComprehension struct {
	pattern        PatternElement
	where          Where
	listDefinition Expression
	key            string
	notNil         bool
	err            error
}

func PatternComprehensionCreate(pattern PatternElement, where Where, listDefinition Expression) PatternComprehension {
	patternComprehension := PatternComprehension{
		pattern:        pattern,
		where:          where,
		listDefinition: listDefinition,
	}
	patternComprehension.key = getAddress(&patternComprehension)
	return patternComprehension
}

func PatternComprehensionError(err error) PatternComprehension {
	return PatternComprehension{err: err}
}

func PatternComprehensionBasedOn(pattern RelationshipPattern) PatternComprehensionOngoingDefinitionWithPattern {
	if !pattern.isNotNil() {
		return PatternComprehensionBuilder{err: errors.New("pattern comprehension builder: a pattern is required")}
	}
	return PatternComprehensionBuilderCreate(pattern)
}

func PatternComprehensionBasedOnNamePath(pattern NamedPath) PatternComprehensionOngoingDefinitionWithPattern {
	if !pattern.isNotNil() {
		return PatternComprehensionBuilder{err: errors.New("pattern comprehension builder: a pattern is required")}
	}
	return PatternComprehensionBuilderCreate(pattern)
}

func (p PatternComprehension) getError() error {
	return p.err
}

func (p PatternComprehension) accept(visitor *CypherRenderer) {
	visitor.enter(p)
	p.pattern.accept(visitor)
	VisitIfNotNull(p.where, visitor)
	PIPE.accept(visitor)
	p.listDefinition.accept(visitor)
	visitor.leave(p)
}

func (p PatternComprehension) enter(renderer *CypherRenderer) {
	renderer.append("[")
}

func (p PatternComprehension) leave(renderer *CypherRenderer) {
	renderer.append("]")
}

func (p PatternComprehension) getKey() string {
	return p.key
}

func (p PatternComprehension) isNotNil() bool {
	return p.notNil
}

func (p PatternComprehension) GetExpressionType() ExpressionType {
	return "PatternComprehension"
}

type PatternComprehensionOngoingDefinitionWithoutReturn interface {
	/**
	 * @param variables the elements to be returned from the pattern
	 * @return The final definition of the pattern comprehension
	 * @see #returning(Expression...)
	 */
	returningByNamed(variables ...Named) PatternComprehension
	/**
	 * @param listDefinition Defines the elements to be returned from the pattern
	 * @return The final definition of the pattern comprehension
	 */
	returning(listDefinitions ...Expression) PatternComprehension
}

type PatternComprehensionOngoingDefinitionWithPattern interface {
	PatternComprehensionOngoingDefinitionWithoutReturn
	where(condition Condition) PatternComprehensionOngoingDefinitionWithPatternAndWhere
	wherePattern(pathPattern RelationshipPattern) PatternComprehensionOngoingDefinitionWithPatternAndWhere
}

type PatternComprehensionOngoingDefinitionWithPatternAndWhere interface {
	PatternComprehensionOngoingDefinitionWithoutReturn
	And(condition Condition) PatternComprehensionOngoingDefinitionWithPatternAndWhere
	AndPattern(pattern RelationshipPattern) PatternComprehensionOngoingDefinitionWithPatternAndWhere
	Or(condition Condition) PatternComprehensionOngoingDefinitionWithPatternAndWhere
	OrPattern(pattern RelationshipPattern) PatternComprehensionOngoingDefinitionWithPatternAndWhere
	getError() error
}

type PatternComprehensionBuilder struct {
	pattern          PatternElement
	conditionBuilder ConditionBuilder
	err              error
}

func PatternComprehensionBuilderCreate(pattern PatternElement) PatternComprehensionBuilder {
	return PatternComprehensionBuilder{
		pattern:          pattern,
		conditionBuilder: ConditionBuilderCreate(),
	}
}

func (p PatternComprehensionBuilder) returningByNamed(variables ...Named) PatternComprehension {
	return p.returning(CreateSymbolicNameByNamed(variables...)...)
}

func (p PatternComprehensionBuilder) returning(listDefinitions ...Expression) PatternComprehension {
	where := Where{}
	condition := p.conditionBuilder.buildCondition()
	if condition != nil && condition.isNotNil() {
		where = WhereCreate(condition)
	}
	return PatternComprehensionCreate(p.pattern, where, ListOrSingleExpression(listDefinitions...))
}

func (p PatternComprehensionBuilder) where(condition Condition) PatternComprehensionOngoingDefinitionWithPatternAndWhere {
	p.conditionBuilder.Where(condition)
	return p
}

func (p PatternComprehensionBuilder) wherePattern(pathPattern RelationshipPattern) PatternComprehensionOngoingDefinitionWithPatternAndWhere {
	if !pathPattern.isNotNil() {
		return PatternComprehensionBuilder{err: errors.New("pattern comprehension builder: path pattern must not be nil")}
	}
	return p.where(RelationshipPatternConditionCreate(pathPattern))
}

func (p PatternComprehensionBuilder) And(condition Condition) PatternComprehensionOngoingDefinitionWithPatternAndWhere {
	p.conditionBuilder.And(condition)
	return p
}

func (p PatternComprehensionBuilder) AndPattern(pattern RelationshipPattern) PatternComprehensionOngoingDefinitionWithPatternAndWhere {
	return p.And(RelationshipPatternConditionCreate(pattern))
}

func (p PatternComprehensionBuilder) Or(condition Condition) PatternComprehensionOngoingDefinitionWithPatternAndWhere {
	p.conditionBuilder.Or(condition)
	return p
}

func (p PatternComprehensionBuilder) OrPattern(pattern RelationshipPattern) PatternComprehensionOngoingDefinitionWithPatternAndWhere {
	return p.Or(RelationshipPatternConditionCreate(pattern))
}

func (p PatternComprehensionBuilder) getError() error {
	return p.err
}
