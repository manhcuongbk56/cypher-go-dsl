package cypher_go_dsl

type BuildableStatement interface {
	Build() Statement
}

type OngoingReadingWithoutWhere interface {
	OngoingReading
}

type OngoingReading interface {
	ExposesReturning
}

type OngoingReadingAndReturn interface {
	BuildableStatement
}
