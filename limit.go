package cypher

import "errors"

type Limit struct {
	limitAmount NumberLiteral
	key         string
	notNil      bool
	err         error
}

func LimitCreate(number int) Limit {
	if number == 0 {
		return LimitError(errors.New("limit can not be zero"))
	}
	literal := NumberLiteralCreate1(number)
	l := Limit{limitAmount: literal, notNil: true}
	l.key = getAddress(&l)
	return l
}

func LimitError(err error) Limit {
	return Limit{
		err: err,
	}
}

func (l Limit) GetError() error {
	return l.err
}

func (l Limit) isNotNil() bool {
	return l.notNil
}

func (l Limit) getKey() string {
	return l.key
}

func (l Limit) accept(visitor *CypherRenderer) {
	visitor.enter(l)
	l.limitAmount.accept(visitor)
	visitor.leave(l)
}

func (l Limit) enter(renderer *CypherRenderer) {
	renderer.append(" LIMIT ")
}

func (l Limit) leave(renderer *CypherRenderer) {
}
