package cypher_go_dsl

type Propertied interface {
	ExposesProperties
	PropertyContainer
}
