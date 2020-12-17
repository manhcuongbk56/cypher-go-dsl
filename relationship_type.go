package cypher_go_dsl

import "strings"

type RelationshipTypes struct {
	values []string
	key    string
	notNil bool
	err    error
}

func RelationshipTypesCreate(types []string) RelationshipTypes {
	r := RelationshipTypes{
		values: types,
		notNil: true,
	}
	r.key = getAddress(&r)
	return r
}

func RelationshipTypesError(err error) RelationshipTypes {
	return RelationshipTypes{
		err: err,
	}
}

func (r RelationshipTypes) getError() error {
	return r.err
}

func (r RelationshipTypes) isNotNil() bool {
	return r.notNil
}

func (r RelationshipTypes) getKey() string {
	return r.key
}

func (r RelationshipTypes) accept(visitor *CypherRenderer) {
	(*visitor).enter(r)
	(*visitor).leave(r)
}

func (r RelationshipTypes) enter(renderer *CypherRenderer) {
	typeWithPrefix := make([]string, 0)
	for _, typeRaw := range r.values {
		if typeRaw == "" {
			continue
		}
		typeWithPrefix = append(typeWithPrefix, RelTypeStart+escapeName(typeRaw))
	}
	renderer.builder.WriteString(strings.Join(typeWithPrefix, RelTypSeparator))
}

func (r RelationshipTypes) leave(renderer *CypherRenderer) {
}
