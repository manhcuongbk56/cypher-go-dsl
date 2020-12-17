package cypher_go_dsl

func ConditionsNot(element PatternElement) Condition {
	return ExcludedPatternCreate(element)
}
