package cypher_go_dsl

type Match struct {
	optional      bool
	pattern       Pattern
	optionalWhere Where
	key           string
	notNil        bool
	err           error
}

func MatchCreate(optional bool, pattern Pattern, optionalWhere Where) Match {
	m := Match{
		optional:      optional,
		pattern:       pattern,
		optionalWhere: optionalWhere,
		notNil:        true,
	}
	m.key = getAddress(&m)
	return m
}

func (match Match) getError() error {
	return match.err
}

func (match Match) isNotNil() bool {
	return match.notNil
}

func (match Match) isOptional() bool {
	return match.optional
}

func (match Match) accept(visitor *CypherRenderer) {
	visitor.enter(match)
	match.pattern.accept(visitor)
	VisitIfNotNull(match.optionalWhere, visitor)
	visitor.leave(match)
}

func (match Match) getKey() string {
	return match.key
}

func (match Match) enter(renderer *CypherRenderer) {
	if match.isOptional() {
		renderer.builder.WriteString("OPTIONAL ")
	}
	renderer.builder.WriteString("MATCH ")
}

func (match Match) leave(renderer *CypherRenderer) {
	renderer.builder.WriteString(" ")
}
