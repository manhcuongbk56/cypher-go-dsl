package cypher

import (
	"errors"
)

type UnionQuery struct {
	all               bool
	firstQuery        SingleQuery
	additionalQueries []UnionPart
	key               string
	notNil            bool
	err               error
}

func UnionQueryCreate(all bool, queries []SingleQuery) UnionQuery {
	if queries == nil || len(queries) < 2 {
		return UnionQueryError(errors.New("at least two queries are needed"))
	}
	for _, query := range queries {
		if query != nil && query.getError() != nil {
			return UnionQueryError(query.getError())
		}
	}
	unionParts := make([]UnionPart, 0)
	for _, query := range queries[1:] {
		unionParts = append(unionParts, UnionPartCreate(all, query))
	}
	return unionQueryCreate1(all, queries[0], unionParts)
}

func unionQueryCreate1(all bool, firstQuery SingleQuery, additionalQueries []UnionPart) UnionQuery {

	union := UnionQuery{
		all:               all,
		firstQuery:        firstQuery,
		additionalQueries: additionalQueries,
		notNil:            true,
	}
	union.key = getAddress(&union)
	return union
}

func UnionQueryError(err error) UnionQuery {
	return UnionQuery{
		err: err,
	}
}

func (q UnionQuery) addAdditionalQueries(newAdditionalQueries []SingleQuery) UnionQuery {
	queries := make([]SingleQuery, 0)
	queries = append(queries, q.firstQuery)
	for _, unionPart := range q.additionalQueries {
		queries = append(queries, unionPart.query)
	}
	queries = append(queries, newAdditionalQueries...)
	return UnionQueryCreate(q.all, queries)
}

func (q UnionQuery) getError() error {
	return q.err
}

func (q UnionQuery) accept(visitor *CypherRenderer) {
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
