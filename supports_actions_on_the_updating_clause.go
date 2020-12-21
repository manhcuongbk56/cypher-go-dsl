package cypher

type SupportsActionsOnTheUpdatingClause interface {
	on(mergeType MERGE_TYPE, expressions ...Expression) (SupportsActionsOnTheUpdatingClause, error)
}
