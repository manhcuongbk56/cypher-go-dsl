package cypher_go_dsl

type PatternComprehension struct {
}

type PatternComprehensionOngoingDefinitionWithoutReturn interface {
	/**
	 * @param variables the elements to be returned from the pattern
	 * @return The final definition of the pattern comprehension
	 * @see #returning(Expression...)
	 */
	returningDefault(variables ...Named) PatternComprehension
	/**
	 * @param listDefinition Defines the elements to be returned from the pattern
	 * @return The final definition of the pattern comprehension
	 */
	returning(listDefinitions ...Expression) PatternComprehension
}

type PatternComprehensionOngoingDefinitionWithPattern interface {
	PatternComprehensionOngoingDefinitionWithoutReturn
	where(condition Condition) PatternComprehensionOngoingDefinitionWithPatternAndWhere
}

type PatternComprehensionOngoingDefinitionWithPatternAndWhere interface {
	PatternComprehensionOngoingDefinitionWithoutReturn
}
