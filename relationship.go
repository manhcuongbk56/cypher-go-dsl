package cypher_go_dsl

type Relationship struct {
	left    *Node
	right   *Node
	details *RelationshipDetails
}

func (r Relationship) RelationshipTo(node Node, types ...string) RelationshipPattern {
	var relationshipI interface{} = (*(r.right)).RelationshipTo(node, types...)
	relationship, _ := relationshipI.(Relationship)
	return FirstElement(r).Add(relationship)
}

func (r Relationship) RelationshipFrom(node Node, types ...string) RelationshipPattern {
	panic("implement me")
}

func (r Relationship) RelationshipBetween(node Node, types ...string) RelationshipPattern {
	panic("implement me")
}

func (r Relationship) IsPatternElement() bool {
	return true
}

func (r Relationship) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(r)
	r.left.Accept(visitor)
	r.details.Accept(visitor)
	r.right.Accept(visitor)
	(*visitor).Leave(r)
}

func (r Relationship) GetType() VisitableType {
	return RelationshipVisitable
}

func (r Relationship) Enter(renderer *CypherRenderer) {
}

func (r Relationship) Leave(renderer *CypherRenderer) {
}

