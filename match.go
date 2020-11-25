package cypher_go_dsl

type Match struct {
	optional      bool
	pattern       Pattern
	optionalWhere *Where
}

func (match Match) isOptional() bool  {
	return match.optional
}

func (match Match) Accept(visitor *CypherRenderer) {
	visitor.EnterA(match, &match)
	match.pattern.Accept(visitor)
	VisitIfNotNull(match.optionalWhere, visitor)
	visitor.LeaveA(match, &match)
}

func (match Match) GetType() VisitableType {
	return MatchVisitable
}

func (match Match) Enter(renderer *CypherRenderer) {
	if match.isOptional() {
		renderer.builder.WriteString("OPTIONAL ")
	}
	renderer.builder.WriteString("MATCH ")
}

func (match Match) Leave(renderer *CypherRenderer) {
	renderer.builder.WriteString(" ")
}



