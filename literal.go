package cypher_go_dsl

type Literal interface {
	Expression
	GetContent() interface{}
	AsString() string
}
