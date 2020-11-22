package cypher_go_dsl

type ExposesMatch interface {
	Match(pattern ...c.PatternElement) OngoingReadingWithoutWhere

	OptionalMatch(pattern ...c.PatternElement)
}
