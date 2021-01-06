package cypher

import "golang.org/x/xerrors"

func ConditionsNotByPattern(element PatternElement) Condition {
	return ExcludedPatternCreate(element)
}

func ConditionsNot(condition Condition) Condition {
	if condition == nil || !condition.isNotNil() {
		return CompoundConditionError(xerrors.New("condition not: condition to negate must not be nil"))
	}
	return condition.Not().Get()
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
