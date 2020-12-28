package cypher

func ConditionsNot(element PatternElement) Condition {
	return ExcludedPatternCreate(element)
}

func ConditionsNoCondition() Condition {
	return EMPTY_CONDITION
}

func ConditionsIsTrue() Condition {
	return TRUE_CONDITION
}

func ConditionsIsFalse() Condition {
	return FALSE_CONDITION
}
