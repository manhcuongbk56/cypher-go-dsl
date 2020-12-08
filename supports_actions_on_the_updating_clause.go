package cypher_go_dsl

type SupportsActionsOnTheUpdatingClause interface {
	on(mergeType MERGE_TYPE, expressions ...Expression) (SupportsActionsOnTheUpdatingClause, error)
}
