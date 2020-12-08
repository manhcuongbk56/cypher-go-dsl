package cypher_go_dsl

import (
	"errors"
	"fmt"
)

type UnionQuery struct {
	all               bool
	firstQuery        SingleQuery
	additionalQueries []UnionPart
	key               string
	notNil            bool
}

func UnionQueryCreate(all bool, firstQuery SingleQuery, additionalQueries []UnionPart) UnionQuery {
	return UnionQuery{
		all:               all,
		firstQuery:        firstQuery,
		additionalQueries: additionalQueries,
		notNil:            true,
	}
}

func UnionQueryCreate1(all bool, queries []SingleQuery) (UnionQuery, error) {
	if queries == nil || len(queries) < 2 {
		return UnionQuery{}, errors.New("at least two queries are needed")
	}
	unionParts := make([]UnionPart, 0)
	for _, query := range queries[1:] {
		unionParts = append(unionParts, UnionPartCreate(all, query))
	}
	return UnionQueryCreate(all, queries[0], unionParts), nil
}

func (q UnionQuery) addAdditionalQueries(newAdditionalQueries []SingleQuery) (UnionQuery, error) {
	queries := make([]SingleQuery, 0)
	queries = append(queries, q.firstQuery)
	for _, unionPart := range q.additionalQueries {
		queries = append(queries, unionPart.query)
	}
	queries = append(queries, newAdditionalQueries...)
	return UnionQueryCreate1(q.all, queries)
}

func (q UnionQuery) accept(visitor *CypherRenderer) {
	q.key = fmt.Sprint(&q)
	visitor.enter(q)
	q.firstQuery.accept(visitor)
	for _, additionalQuery := range q.additionalQueries {
		additionalQuery.accept(visitor)
	}
	visitor.leave(q)
}

func (q UnionQuery) enter(renderer *CypherRenderer) {
}

func (q UnionQuery) leave(renderer *CypherRenderer) {
}

func (q UnionQuery) getKey() string {
	return q.key
}

func (q UnionQuery) isNotNil() bool {
	return q.notNil
}
