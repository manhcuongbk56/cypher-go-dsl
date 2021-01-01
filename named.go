package cypher

type Named interface {
	CanHasError
	getRequiredSymbolicName() SymbolicName
	getSymbolicName() SymbolicName
	isNotNil() bool
}
