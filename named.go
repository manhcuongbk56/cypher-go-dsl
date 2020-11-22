package cypher_go_dsl

type Named interface {
	getSymbolicName() SymbolicName
}
