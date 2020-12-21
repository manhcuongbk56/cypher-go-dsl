package cypher

type ExposesCreate interface {
	Create(element ...PatternElement) OngoingUpdateAndExposesSet
}
