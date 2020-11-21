package cypher_go_dsl

import "cypher-go-dsl/expression"

type ExposesProperties interface {
	WithRawProperties(keysAndValues ...interface{}) (Node, error)

	WithProperties(newProperties expression.MapExpression) Node
}
