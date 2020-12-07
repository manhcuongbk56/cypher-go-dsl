package cypher_go_dsl

type ExposesMatch interface {
	match(pattern ...PatternElement) OngoingReadingWithoutWhere

	optionalMatch(pattern ...PatternElement) OngoingReadingWithoutWhere

	MatchDefault(optional bool, pattern ...PatternElement) OngoingReadingWithoutWhere
}
