package cypher

type ExposesSubqueryCall interface {
	Call(statement Statement) OngoingReadingWithoutWhere
}
