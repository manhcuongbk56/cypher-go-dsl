package cypher

type Aliased interface {
	GetAlias() string
	AsName() SymbolicName
}
