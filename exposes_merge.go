package cypher

type ExposesMerge interface {
	Merge(pattern ...PatternElement) OngoingUpdateAndExposesSet
}
