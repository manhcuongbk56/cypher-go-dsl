package cypher

type Renderer interface {
	Renderer(statement Statement) string
}
