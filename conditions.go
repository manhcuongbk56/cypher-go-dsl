package cypher

func ConditionsNot(element PatternElement) Condition {
	return ExcludedPatternCreate(element)
}
