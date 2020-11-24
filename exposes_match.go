package cypher_go_dsl

type ExposesMatch interface {
	Match(pattern ...PatternElement) OngoingReadingWithoutWhere

	OptionalMatch(pattern ...PatternElement) OngoingReadingWithoutWhere
}
