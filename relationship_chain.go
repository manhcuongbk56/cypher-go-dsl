package cypher

import "errors"

type RelationshipChain struct {
	relationships []Relationship
	key           string
	notNil        bool
	err           error
}

func RelationshipChainCreate(relationship Relationship) RelationshipChain {
	relations := make([]Relationship, 1)
	relations[0] = relationship
	return RelationshipChain{relationships: relations, notNil: true}
}

func RelationshipChainError(msg string) RelationshipChain {
	return RelationshipChain{err: errors.New(msg)}
}

func (r RelationshipChain) GetError() error {
	return r.err
}

func (r RelationshipChain) isNotNil() bool {
	return r.notNil
}

func (r RelationshipChain) getKey() string {
	return r.key
}

func (r RelationshipChain) RelationshipTo(node Node, types ...string) RelationshipChain {
	newRelation := r.relationships[len(r.relationships)-1].right.RelationshipTo(node, types...)
	r.relationships = append(r.relationships, newRelation)
	return r
}

func (r RelationshipChain) RelationshipFrom(node Node, types ...string) RelationshipChain {
	newRelation := r.relationships[len(r.relationships)-1].right.RelationshipFrom(node, types...)
	r.relationships = append(r.relationships, newRelation)
	return r
}

func (r RelationshipChain) RelationshipBetween(node Node, types ...string) RelationshipChain {
	newRelation := r.relationships[len(r.relationships)-1].right.RelationshipBetween(node, types...)
	r.relationships = append(r.relationships, newRelation)
	return r
}

func (r RelationshipChain) Named(name string) Relationship {
	return RelationshipError(errors.New("can not use named for relationship chain"))
}

func (r RelationshipChain) NamedC(name string) RelationshipChain {
	namedRelation := r.relationships[len(r.relationships)-1].Named(name)
	r.relationships[len(r.relationships)-1] = namedRelation
	return r
}

func (r RelationshipChain) Unbounded() RelationshipChain {
	namedRelation := r.relationships[len(r.relationships)-1].Unbounded()
	r.relationships[len(r.relationships)-1] = namedRelation
	return r
}

func (r RelationshipChain) Min(minimum int) RelationshipChain {
	namedRelation := r.relationships[len(r.relationships)-1].Min(minimum)
	r.relationships[len(r.relationships)-1] = namedRelation
	return r
}

func (r RelationshipChain) Length(minimum int, maximum int) RelationshipChain {
	namedRelation := r.relationships[len(r.relationships)-1].Length(minimum, maximum)
	r.relationships[len(r.relationships)-1] = namedRelation
	return r
}

func (r RelationshipChain) Max(maximum int) RelationshipChain {
	namedRelation := r.relationships[len(r.relationships)-1].Max(maximum)
	r.relationships[len(r.relationships)-1] = namedRelation
	return r
}

func (r RelationshipChain) Properties(newProperties MapExpression) RelationshipChain {
	namedRelation := r.relationships[len(r.relationships)-1].WithProperties(newProperties)
	r.relationships[len(r.relationships)-1] = namedRelation
	return r
}

func (r RelationshipChain) PropertiesRaw(keysAndValues ...interface{}) RelationshipChain {
	namedRelation := r.relationships[len(r.relationships)-1].WithRawProperties(keysAndValues...)
	r.relationships[len(r.relationships)-1] = namedRelation
	return r
}

func (r RelationshipChain) accept(visitor *CypherRenderer) {
	visitor.enter(r)
	lastNode := Node{}
	for _, relationship := range r.relationships {
		relationship.left.accept(visitor)
		relationship.details.accept(visitor)
		lastNode = relationship.right
	}
	VisitIfNotNull(lastNode, visitor)
	visitor.leave(r)
}

func (r RelationshipChain) enter(renderer *CypherRenderer) {
}

func (r RelationshipChain) leave(renderer *CypherRenderer) {
}

func (r RelationshipChain) Add(relationship Relationship) RelationshipChain {
	r.relationships = append(r.relationships, relationship)
	return r
}

func (r RelationshipChain) IsPatternElement() bool {
	return true
}
