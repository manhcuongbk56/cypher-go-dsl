package cypher_go_dsl

type Asterisk struct {
	content string
	key     string
	notNil  bool
	err     error
}

func AsteriskCreate() Asterisk {
	asterisk := Asterisk{
		content: "*",
		notNil:  true,
	}
	asterisk.key = getAddress(&asterisk)
	return asterisk
}

func (s Asterisk) getError() error {
	return s.err
}

func (s Asterisk) isNotNil() bool {
	return s.notNil
}

func escapeAsterisk(value string) string {
	return "'" + "*" + "'"
}

func (s Asterisk) getKey() string {
	return s.key
}

func (s Asterisk) GetExpressionType() ExpressionType {
	return "Asterisk"
}

func (s Asterisk) GetContent() interface{} {
	return s.content
}

func (s Asterisk) AsString() string {
	return s.content
}

func (s Asterisk) accept(visitor *CypherRenderer) {
	visitor.enter(s)
	visitor.leave(s)
}

func (s Asterisk) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(escapeAsterisk(s.AsString()))
}

func (s Asterisk) leave(renderer *CypherRenderer) {
}
