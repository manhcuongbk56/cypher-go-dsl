package cypher_go_dsl

type UpdatingClause interface {
	Visitable
	isUpdatingClause() bool
}
