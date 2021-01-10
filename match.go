package cypher

type match struct {
	optional      bool
	pattern       Pattern
	optionalWhere Where
	key           string
	notNil        bool
	err           error
}

func matchCreate(optional bool, pattern Pattern, optionalWhere Where) match {
	if pattern.GetError() != nil {
		return matchError(pattern.GetError())
	}
	if optionalWhere.GetError() != nil {
		return matchError(optionalWhere.GetError())
	}
	m := match{
		optional:      optional,
		pattern:       pattern,
		optionalWhere: optionalWhere,
		notNil:        true,
	}
	m.key = getAddress(&m)
	return m
}

func matchError(err error) match {
	return match{err: err}
}

func (match match) GetError() error {
	return match.err
}

func (match match) isNotNil() bool {
	return match.notNil
}

func (match match) isOptional() bool {
	return match.optional
}

func (match match) accept(visitor *CypherRenderer) {
	visitor.enter(match)
	match.pattern.accept(visitor)
	VisitIfNotNull(match.optionalWhere, visitor)
	visitor.leave(match)
}

func (match match) getKey() string {
	return match.key
}

func (match match) enter(renderer *CypherRenderer) {
	if match.isOptional() {
		renderer.append("OPTIONAL ")
	}
	renderer.append("MATCH ")
}

func (match match) leave(renderer *CypherRenderer) {
	renderer.append(" ")
}
