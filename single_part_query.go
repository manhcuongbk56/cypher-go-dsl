package cypher_go_dsl

import "errors"

type SinglePartQuery struct {
	precedingClauses []Visitable
	aReturn *Return
}

func (s SinglePartQuery) Accept(visitor Visitor) {
	for _, clause := range s.precedingClauses{
		clause.Accept(visitor)
		VisitIfNotNull(s.aReturn, visitor)
	}
}

func NewSinglePartQuery(clauses []Visitable, aReturn *Return) (*SinglePartQuery, error){
	if len(clauses) == 0  {
		_, isMatch := clauses[len(clauses)-1].(Match)
		if isMatch {
			if aReturn == nil {
				return nil, errors.New("Required return clause")
			}
		}
	}
	return &SinglePartQuery{
		precedingClauses: clauses,
		aReturn:          aReturn,
	}, nil
}

func (s SinglePartQuery) GetType() VisitableType {
	return SinglePartQueryVisitable
}


