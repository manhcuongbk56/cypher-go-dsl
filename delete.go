package cypher_go_dsl

type Delete struct {
	deleteItems ExpressionList
	detach      bool
	key         string
	notNil      bool
	err         error
}

func DeleteCreate(deleteItems ExpressionList, detach bool) Delete {
	if deleteItems.getError() != nil {
		return DeleteError(deleteItems.getError())
	}
	d := Delete{
		deleteItems: deleteItems,
		detach:      detach,
		notNil:      true,
	}
	d.key = getAddress(&d)
	return d
}

func DeleteError(err error) Delete {
	return Delete{
		err: err,
	}
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
		renderer.append("DETACH ")
	}
	renderer.append("DELETE ")
}

func (d Delete) leave(renderer *CypherRenderer) {
	renderer.append(" ")
}

func (d Delete) getKey() string {
	return d.key
}

func (d Delete) isNotNil() bool {
	return d.notNil
}

func (d Delete) isUpdatingClause() bool {
	return true
}
