package cypher_go_dsl

type Literal interface {
	IsExpression
	GetContent() interface{}
	AsString() string
}
