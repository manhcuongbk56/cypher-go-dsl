package cypher

type PropertyContainer interface {
	Named
	Property(name string) Property
}
