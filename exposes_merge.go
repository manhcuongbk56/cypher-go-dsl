package cypher_go_dsl

type ExposesMerge interface {
	merge(pattern ...PatternElement) OngoingUpdate
}
