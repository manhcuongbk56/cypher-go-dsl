package cypher

type UpdatingClause interface {
	Visitable
	isUpdatingClause() bool
}
