package cypher

type ExposesMatch interface {
	Match(pattern ...PatternElement) OngoingReadingWithoutWhere
	OptionalMatch(pattern ...PatternElement) OngoingReadingWithoutWhere
	MatchDefault(optional bool, pattern ...PatternElement) OngoingReadingWithoutWhere
}
