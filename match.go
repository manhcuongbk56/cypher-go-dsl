package cypher_go_dsl

type Match struct {
	optional      bool
	pattern       Pattern
	optionalWhere *Where
}

func (match Match) isOptional() bool  {
	return match.optional
}

func (match Match) Accept(visitor Visitor) {
	visitor.Enter(match)
	match.pattern.Accept(visitor)
	VisitIfNotNull(match.optionalWhere, visitor)
	visitor.Leave(match)
}

func (match Match) GetType() VisitableType {
	return MatchVisitable
}


