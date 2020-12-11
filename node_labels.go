package cypher_go_dsl

type NodeLabels struct {
	values []NodeLabel
	notNil bool
	key    string
	err    error
}

func NodeLabelsCreate(values []NodeLabel) NodeLabels {
	n := NodeLabels{
		values: values,
		notNil: true,
	}
	n.key = getAddress(&n)
	return n
}

func (n NodeLabels) getError() error {
	return n.err
}

func (n NodeLabels) accept(visitor *CypherRenderer) {
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
