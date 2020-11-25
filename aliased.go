package cypher_go_dsl

type Aliased interface {
	GetAlias() string

	AsName() SymbolicName
}



