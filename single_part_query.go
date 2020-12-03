package cypher_go_dsl

import (
	"errors"
	"fmt"
)

type SinglePartQuery struct {
	precedingClauses []Visitable
	aReturn          Return
	key              string
	notNil           bool
}

func (s SinglePartQuery) isNotNil() bool {
	return s.notNil
}

func (s SinglePartQuery) getKey() string {
	return s.key
}

func (s SinglePartQuery) accept(visitor *CypherRenderer) {
	s.key = fmt.Sprint(&s)
	for _, clause := range s.precedingClauses {
		clause.accept(visitor)
		VisitIfNotNull(s.aReturn, visitor)
	}
}

func SinglePartQueryCreate(clauses []Visitable, aReturn Return) (SinglePartQuery, error) {
	if len(clauses) > 0 {
		_, isMatch := clauses[len(clauses)-1].(Match)
		if isMatch {
			if !aReturn.isNotNil() {
				return SinglePartQuery{}, errors.New("Required return clause")
			}
		}
	}
	return SinglePartQuery{
		precedingClauses: clauses,
		aReturn:          aReturn,
	}, nil
}

func (s SinglePartQuery) enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (s SinglePartQuery) leave(renderer *CypherRenderer) {
	panic("implement me")
}
