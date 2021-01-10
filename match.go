package cypher

type MatchPhrase struct {
	optional      bool
	pattern       Pattern
	optionalWhere Where
	key           string
	notNil        bool
	err           error
}

func MatchPhraseCreate(optional bool, pattern Pattern, optionalWhere Where) MatchPhrase {
	if pattern.GetError() != nil {
		return matchError(pattern.GetError())
	}
	if optionalWhere.GetError() != nil {
		return matchError(optionalWhere.GetError())
	}
	m := MatchPhrase{
		optional:      optional,
		pattern:       pattern,
		optionalWhere: optionalWhere,
		notNil:        true,
	}
	m.key = getAddress(&m)
	return m
}

func matchError(err error) MatchPhrase {
	return MatchPhrase{err: err}
}

func (match MatchPhrase) GetError() error {
	return match.err
}

func (match MatchPhrase) isNotNil() bool {
	return match.notNil
}

func (match MatchPhrase) isOptional() bool {
	return match.optional
}

func (match MatchPhrase) accept(visitor *CypherRenderer) {
	visitor.enter(match)
	match.pattern.accept(visitor)
	VisitIfNotNull(match.optionalWhere, visitor)
	visitor.leave(match)
}

func (match MatchPhrase) getKey() string {
	return match.key
}

func (match MatchPhrase) enter(renderer *CypherRenderer) {
	if match.isOptional() {
		renderer.append("OPTIONAL ")
	}
	renderer.append("MATCH ")
}

func (match MatchPhrase) leave(renderer *CypherRenderer) {
	renderer.append(" ")
}
