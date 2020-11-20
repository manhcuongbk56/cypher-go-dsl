package cypher_go_dsl

type Node struct {
	SymbolicName string
	Labels []string
	Properties string
}

func (node Node) Accept(visitor Visitor) {
}

func (node2 Node) RelationshipTo(node Node, types ...string) RelationshipPattern {
	panic("implement me")
}

func (node2 Node) RelationshipFrom(node Node, types ...string) RelationshipPattern {
	panic("implement me")
}

func (node2 Node) RelationshipBetween(node Node, types ...string) RelationshipPattern {
	panic("implement me")
}

func (node Node) WithProperties(keysAndValues ...interface{}) {
	panic("implement me")
}

func (node Node) Property(name string) {
	panic("implement me")
}



