package cypher

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

func (r RelationshipTypes) GetError() error {
	return r.err
}

func (r RelationshipTypes) isNotNil() bool {
	return r.notNil
}

func (r RelationshipTypes) getKey() string {
	return r.key
}

func (r RelationshipTypes) accept(visitor *CypherRenderer) {
	visitor.enter(r)
	visitor.leave(r)
}

func (r RelationshipTypes) enter(renderer *CypherRenderer) {
	typeWithPrefix := make([]string, 0)
	for _, typeRaw := range r.values {
		if typeRaw == "" {
			continue
		}
		typeWithPrefix = append(typeWithPrefix, escapeName(typeRaw))
	}
	if len(r.values) > 0 {
		renderer.append(RelTypeStart).append(strings.Join(typeWithPrefix, RelTypSeparator))
	}
}

func (r RelationshipTypes) leave(renderer *CypherRenderer) {
}
