package cypher_go_dsl

type ExposesProperties interface {
	WithRawProperties(keysAndValues ...interface{}) (Node, error)

	WithProperties(newProperties MapExpression) Node
}
