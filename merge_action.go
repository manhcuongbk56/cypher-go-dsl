package cypher_go_dsl

import "fmt"

type MergeAction struct {
	mergeType MERGE_TYPE
	set       Set
	key       string
	notNil    bool
}

func (m MergeAction) accept(visitor *CypherRenderer) {
	m.key = fmt.Sprint(&m)
	visitor.enter(m)
	m.set.accept(visitor)
	visitor.leave(m)
}

func (m MergeAction) enter(renderer *CypherRenderer) {
	switch m.mergeType {
	case ON_CREATE:
		renderer.builder.WriteString("ON CREATE")
	case ON_MATCH:
		renderer.builder.WriteString("ON MATCH")
	}
	renderer.builder.WriteString(" ")
}

func (m MergeAction) leave(renderer *CypherRenderer) {
}

func (m MergeAction) getKey() string {
	return m.key
}

func (m MergeAction) isNotNil() bool {
	return m.notNil
}

type MERGE_TYPE string

const (
	ON_CREATE MERGE_TYPE = "onCreate"
	ON_MATCH             = "onMatch"
)
