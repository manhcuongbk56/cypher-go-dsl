package cypher_go_dsl

type MatchBuilder struct {
	patternList []cypher_go_dsl.PatternElement
	condition   ConditionBuilder
	optional    bool
}

func (builder MatchBuilder) buildMatch() Match  {
	pattern := cypher_go_dsl.Pattern{PatternElements: builder.patternList}
	return
}
