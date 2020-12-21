package cypher

type Literal interface {
	Expression
	GetContent() interface{}
	AsString() string
}
