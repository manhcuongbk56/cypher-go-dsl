package cypher_go_dsl

type ExposesSubqueryCall interface {

	call(statement Statement) OngoingReadingWithoutWhere

}
