package cypher_go_dsl

type PropertyContainer interface {
	Named
	Property(name string)
}
