package cypher

type Named interface {
	CanHasError
	GetRequiredSymbolicName() SymbolicName
	GetSymbolicName() SymbolicName
	isNotNil() bool
}
