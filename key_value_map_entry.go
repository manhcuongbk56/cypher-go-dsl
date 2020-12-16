package cypher_go_dsl

type KeyValueMapEntry struct {
	entryKey string
	value    Expression
	key      string
	err      error
	notNil   bool
}

func KeyValueMapEntryCreate(key string, value Expression) KeyValueMapEntry {
	if value != nil && value.getError() != nil {
		return KeyValueMapEntryError(value.getError())
	}
	entry := KeyValueMapEntry{
		entryKey: key,
		value:    value,
	}
	entry.key = getAddress(&entry)
	return entry
}

func KeyValueMapEntryError(err error) KeyValueMapEntry {
	return KeyValueMapEntry{
		err: err,
	}
}

func (k KeyValueMapEntry) getError() error {
	return k.err
}

func (k KeyValueMapEntry) accept(visitor *CypherRenderer) {
	visitor.enter(k)
	k.value.accept(visitor)
	visitor.leave(k)
}

func (k KeyValueMapEntry) enter(renderer *CypherRenderer) {
	renderer.append(escapeName(k.entryKey)).append(": ")
}

func (k KeyValueMapEntry) leave(renderer *CypherRenderer) {
}

func (k KeyValueMapEntry) getKey() string {
	return k.key
}

func (k KeyValueMapEntry) isNotNil() bool {
	return k.notNil
}

func (k KeyValueMapEntry) GetExpressionType() ExpressionType {
	return "KeyValueMapEntry"
}
