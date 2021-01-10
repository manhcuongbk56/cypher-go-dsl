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

func (builder MatchBuilder) buildMatch() MatchPhrase {
	pattern := Pattern{patternElements: builder.patternList}
	conditionBuilder := builder.conditionBuilder
	builtCondition := conditionBuilder.buildCondition()
	if builtCondition == nil || !builtCondition.isNotNil() {
		return MatchPhraseCreate(builder.optional, pattern, Where{})
	}
	return MatchPhraseCreate(builder.optional, pattern, WhereCreate(builtCondition))
}
