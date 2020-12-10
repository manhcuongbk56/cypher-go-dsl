package cypher_go_dsl

import "fmt"

type Delete struct {
	deleteItems ExpressionList
	detach      bool
	key         string
	notNil      bool
	err         error
}

func DeleteCreate(deleteItems ExpressionList, detach bool) Delete {
	d := Delete{
		deleteItems: deleteItems,
		detach:      detach,
		notNil:      true,
	}
	d.key = getAddress(&d)
	return d
}

func (d Delete) isDetach() bool {
	return d.detach
}

func (d Delete) getError() error {
	return d.err
}

func (d Delete) accept(visitor *CypherRenderer) {
	visitor.enter(d)
	d.deleteItems.accept(visitor)
	visitor.leave(d)
}

func (d Delete) enter(renderer *CypherRenderer) {
	if d.detach {
		renderer.builder.WriteString("DETACH ")
	}
	renderer.builder.WriteString("DELETE ")
}

func (d Delete) leave(renderer *CypherRenderer) {
	renderer.builder.WriteString(" ")
}

func (d Delete) getKey() string {
	return d.key
}

func (d Delete) isNotNil() bool {
	return d.notNil
}

func (d Delete) isUpdatingClause() bool {
	panic("implement me")
}
