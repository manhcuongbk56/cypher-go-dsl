package cypher

type MergeAction struct {
	mergeType MERGE_TYPE
	set       Set
	key       string
	notNil    bool
	err       error
}

func MergeActionCreate(mergeType MERGE_TYPE, set Set) MergeAction {
	m := MergeAction{
		mergeType: mergeType,
		set:       set,
		notNil:    true,
	}
	m.key = getAddress(&m)
	return m
}

func (m MergeAction) getError() error {
	return m.err
}

func (m MergeAction) accept(visitor *CypherRenderer) {
	visitor.enter(m)
	m.set.accept(visitor)
	visitor.leave(m)
}

func (m MergeAction) enter(renderer *CypherRenderer) {
	switch m.mergeType {
	case ON_CREATE:
		renderer.append("ON CREATE")
	case ON_MATCH:
		renderer.append("ON MATCH")
	}
	renderer.append(" ")
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
