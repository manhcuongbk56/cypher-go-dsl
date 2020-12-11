package cypher_go_dsl

type ExposesCreate interface {
	Create(element ...PatternElement) OngoingUpdate
}
