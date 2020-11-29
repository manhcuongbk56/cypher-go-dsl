package cypher_go_dsl

type Expression interface {
	Visitable
	IsExpression() bool
}

