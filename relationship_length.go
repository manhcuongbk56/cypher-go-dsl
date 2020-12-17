package cypher_go_dsl

import "strconv"

type RelationshipLength struct {
	minimum   *int
	maximum   *int
	unbounded bool
	key       string
	notNil    bool
	err       error
}

func RelationshipLengthCreate(unbounded bool) RelationshipLength {
	r := RelationshipLength{
		unbounded: unbounded,
		notNil:    true,
	}
	r.key = getAddress(&r)
	return r
}

func RelationshipLengthCreate1(minimum *int, maximum *int) RelationshipLength {
	r := RelationshipLength{
		minimum:   minimum,
		maximum:   maximum,
		unbounded: false,
		notNil:    true,
	}
	r.key = getAddress(&r)
	return r
}

func (relationshipLength RelationshipLength) getError() error {
	return relationshipLength.err
}

func (relationshipLength RelationshipLength) accept(visitor *CypherRenderer) {
	visitor.enter(relationshipLength)
	visitor.leave(relationshipLength)
}

func (relationshipLength RelationshipLength) enter(renderer *CypherRenderer) {
	minimum := relationshipLength.minimum
	maximum := relationshipLength.maximum
	if relationshipLength.unbounded {
		renderer.builder.WriteString("*")
		return
	}
	if minimum == nil && maximum == nil {
		return
	}
	renderer.builder.WriteString("*")
	if minimum != nil {
		renderer.builder.WriteString(strconv.Itoa(*minimum))
	}
	renderer.builder.WriteString("..")
	if maximum != nil {
		renderer.builder.WriteString(strconv.Itoa(*maximum))
	}
}

func (relationshipLength RelationshipLength) leave(renderer *CypherRenderer) {
}

func (relationshipLength RelationshipLength) getKey() string {
	return relationshipLength.key
}

func (relationshipLength RelationshipLength) isNotNil() bool {
	return relationshipLength.notNil
}

func (relationshipLength RelationshipLength) Unbounded() RelationshipLength {
	return RelationshipLengthCreate(true)
}
