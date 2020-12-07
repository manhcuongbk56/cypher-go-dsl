package cypher_go_dsl

import "fmt"

type NodeLabels struct {
	values []NodeLabel
	notNil bool
	key    string
}

func NodeLabelsCreate(values []NodeLabel) NodeLabels {
	return NodeLabels{
		values: values,
		notNil: true,
	}
}

func (n NodeLabels) accept(visitor *CypherRenderer) {
	n.key = fmt.Sprint(&n)
	visitor.enter(n)
	for _, value := range n.values {
		value.accept(visitor)
	}
	visitor.leave(n)
}

func (n NodeLabels) enter(renderer *CypherRenderer) {
}

func (n NodeLabels) leave(renderer *CypherRenderer) {
}

func (n NodeLabels) getKey() string {
	return n.key
}

func (n NodeLabels) isNotNil() bool {
	return n.notNil
}
