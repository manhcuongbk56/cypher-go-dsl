package cypher_go_dsl

type Statement interface {
	Visitable
}

type RegularQuery interface {
	Statement
}

type SingleQuery interface {
	Statement
}
