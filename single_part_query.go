package cypher

import (
	"errors"
)

type SinglePartQuery struct {
	precedingClauses []Visitable
	aReturn          Return
	key              string
	notNil           bool
	err              error
}

func SinglePartQueryCreate(clauses []Visitable, aReturn Return) SinglePartQuery {
	if len(clauses) == 0 && !aReturn.isNotNil() {
		return SinglePartQueryError(errors.New("required return clause"))
	}
	if len(clauses) > 0 {
		if _, isMatch := clauses[len(clauses)-1].(Match); isMatch && !aReturn.isNotNil() {
			return SinglePartQueryError(errors.New("required return clause"))
		}
	}
	for _, clause := range clauses {
		if clause != nil && clause.getError() != nil {
			return SinglePartQueryError(clause.getError())
		}
	}
	if aReturn.getError() != nil {
		return SinglePartQueryError(aReturn.getError())
	}
	singlePartQuery := SinglePartQuery{
		precedingClauses: clauses,
		aReturn:          aReturn,
		notNil:           true,
	}
	singlePartQuery.key = getAddress(&singlePartQuery)
	return singlePartQuery
}

func SinglePartQueryError(err error) SinglePartQuery {
	return SinglePartQuery{
		err: err,
	}
}

func (s SinglePartQuery) getError() error {
	return s.err
}

func (s SinglePartQuery) isNotNil() bool {
	return s.notNil
}

func (s SinglePartQuery) getKey() string {
	return s.key
}

func (s SinglePartQuery) accept(visitor *CypherRenderer) {
	for _, clause := range s.precedingClauses {
		clause.accept(visitor)
	}
	VisitIfNotNull(s.aReturn, visitor)
}

func (s SinglePartQuery) enter(renderer *CypherRenderer) {
}

func (s SinglePartQuery) leave(renderer *CypherRenderer) {
}

func (s SinglePartQuery) doesReturnElements() bool {
	return s.aReturn.isNotNil()
}
