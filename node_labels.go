package cypher

type NodeLabels struct {
	values []NodeLabel
	notNil bool
	key    string
	err    error
}

func NodeLabelsCreate(values []NodeLabel) NodeLabels {
	for _, value := range values {
		if value.GetError() != nil {
			return NodeLabelsError(value.GetError())
		}
	}
	n := NodeLabels{
		values: values,
		notNil: true,
	}
	n.key = getAddress(&n)
	return n
}

func NodeLabelsError(err error) NodeLabels {
	return NodeLabels{err: err}
}

func (n NodeLabels) GetError() error {
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
