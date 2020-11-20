package cypher_go_dsl

type ExposesProperties interface {
	WithProperties(keysAndValues ...interface{})
}
