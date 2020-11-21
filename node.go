package cypher_go_dsl

import (
	"cypher-go-dsl/expression"
	v "cypher-go-dsl/visitable"
)

type Node struct {
	symbolicName string
	labels       []string
	properties   Properties
}

func (node Node) Accept(visitor v.Visitor) {
}

func (node Node) RelationshipTo(nodeDest Node, types ...string) RelationshipPattern {
	panic("implement me")
}

func (node Node) RelationshipFrom(nodeDest Node, types ...string) RelationshipPattern {
	panic("implement me")
}

func (node Node) RelationshipBetween(nodeDest Node, types ...string) RelationshipPattern {
	panic("implement me")
}

func (node Node) WithRawProperties(keysAndValues ...interface{}) (Node, error){
	var properties = &expression.MapExpression{}
	if keysAndValues != nil && len(keysAndValues) != 0 {
		var err error
		*properties, err = expression.NewMapExpression(keysAndValues)
		if err != nil {
			return Node{}, err
		}
	}
	return node.WithProperties(*properties), nil
}

func (node Node) WithProperties(newProperties expression.MapExpression) Node {
	return Node{symbolicName: node.symbolicName, labels: node.labels, properties: Properties{newProperties}}
}

func (node Node) Property(name string) {
	panic("implement me")
}



