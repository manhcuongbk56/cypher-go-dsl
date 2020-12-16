package cypher_go_dsl

type Named interface {
	CanHasError
	getRequiredSymbolicName() SymbolicName
	getSymbolicName() SymbolicName
}
