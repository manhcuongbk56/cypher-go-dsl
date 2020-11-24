package cypher_go_dsl

type Renderer interface {
	Renderer(statement Statement) string
}
