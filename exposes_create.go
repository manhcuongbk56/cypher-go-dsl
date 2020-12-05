package cypher_go_dsl

type ExposesCreate interface {
	create(element ...PatternElement) OngoingUpdate
}
