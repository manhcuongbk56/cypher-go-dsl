package cypher_go_dsl

type RelationshipChain struct {
	relationships []Relationship
}

func (r RelationshipChain) RelationshipTo(node Node, types ...string) RelationshipPattern {
	newRelation := (*(r.relationships[len(r.relationships) - 1].right)).RelationshipTo(node, types...)
	r.relationships = append(r.relationships, newRelation)
	return r
}

func (r RelationshipChain) RelationshipFrom(node Node, types ...string) RelationshipPattern {
	panic("implement me")
}

func (r RelationshipChain) RelationshipBetween(node Node, types ...string) RelationshipPattern {
	panic("implement me")
}

func (r RelationshipChain) Accept(visitor Visitor) {
	panic("implement me")
}

func FirstElement(relationship Relationship) RelationshipChain {
	relations := make([]Relationship, 1)
	relations = append(relations, relationship)
	return RelationshipChain{relationships: relations}
}

func (r RelationshipChain) Add(relationship Relationship) RelationshipChain {
	r.relationships = append(r.relationships, relationship)
	return r
}

func (r RelationshipChain) IsPatternElement() bool {
	return true
}
