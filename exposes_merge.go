package cypher_go_dsl

type ExposesMerge interface {
	Merge(pattern ...PatternElement) OngoingUpdate
}
