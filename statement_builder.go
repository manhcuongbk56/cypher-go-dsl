package cypher_go_dsl

type BuildableStatement interface {
	Build() Statement
}

type OngoingReadingWithoutWhere struct {
	ExposesReturningStruct
}

type OngoingReading struct {
	ExposesReturningStruct
}

type OngoingReadingAndReturn interface {
	BuildableStatement
}
