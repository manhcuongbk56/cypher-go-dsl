package cypher

import "strings"

type ListLiteral struct {
	ExpressionContainer
	content []Literal
	key     string
	err     error
	notNil  bool
}

func ListLiteralCreate(contents []Literal) ListLiteral {
	for _, content := range contents {
		if content != nil && content.GetError() != nil {
			return ListLiteralError(content.GetError())
		}
	}
	list := ListLiteral{
		content: contents,
	}
	list.key = getAddress(&list)
	list.ExpressionContainer = ExpressionWrap(list)
	list.notNil = true
	return list
}

func ListLiteralError(err error) ListLiteral {
	return ListLiteral{
		err: err,
	}
}

func (l ListLiteral) GetError() error {
	return l.err
}

func (l ListLiteral) accept(visitor *CypherRenderer) {
	visitor.enter(l)
	for i, literal := range l.content {
		if(i > 0) {
			visitor.append(", ")
		}
		literal.accept(visitor)
	}
	visitor.leave(l)
}

func (l ListLiteral) enter(renderer *CypherRenderer) {
	renderer.append("[")
}

func (l ListLiteral) leave(renderer *CypherRenderer) {
	renderer.append("]")
}

func (l ListLiteral) getKey() string {
	return l.key
}

func (l ListLiteral) isNotNil() bool {
	return l.notNil
}

func (l ListLiteral) GetExpressionType() ExpressionType {
	return "ListLiteral"
}

func (l ListLiteral) GetContent() interface{} {
	return l.content
}

func (l ListLiteral) AsString() string {
	builder := strings.Builder{}
	builder.WriteString("[")
	for i, literal := range l.content {
		if i != 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(literal.AsString())
	}
	builder.WriteString("]")
	return builder.String()
}
