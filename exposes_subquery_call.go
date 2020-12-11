package cypher_go_dsl

type ExposesSubqueryCall interface {
	Call(statement Statement) OngoingReadingWithoutWhere
}
