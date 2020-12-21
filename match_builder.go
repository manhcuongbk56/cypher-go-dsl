package cypher

type MatchBuilder struct {
	patternList      []PatternElement
	conditionBuilder ConditionBuilder
	optional         bool
	notNil           bool
}

func MatchBuilderCreate(optional bool) MatchBuilder {
	return MatchBuilder{
		optional:    optional,
		patternList: make([]PatternElement, 0),
		notNil:      true,
	}
}

func (builder MatchBuilder) buildMatch() Match {
	pattern := Pattern{patternElements: builder.patternList}
	conditionBuilder := builder.conditionBuilder
	var optionalWhere Where = Where{}
	if conditionBuilder.condition != nil {
		builtCondition := conditionBuilder.buildCondition()
		optionalWhere = WhereCreate(builtCondition)
	}
	return MatchCreate(builder.optional, pattern, optionalWhere)
}
