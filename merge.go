package cypher

type merge struct {
	pattern               Pattern
	onCreateOrMatchEvents []Visitable
	key                   string
	notNil                bool
	err                   error
}

var BLANK = RawStringLiteralCreate(" ")

func mergeCreate(pattern Pattern) merge {
	if pattern.GetError() != nil {
		return mergeError(pattern.GetError())
	}
	m := merge{
		pattern:               pattern,
		onCreateOrMatchEvents: make([]Visitable, 0),
		notNil:                true,
	}
	m.key = getAddress(&m)
	return m
}

func mergeCreate1(pattern Pattern, mergeActions []MergeAction) merge {
	if pattern.GetError() != nil {
		return mergeError(pattern.GetError())
	}
	for _, action := range mergeActions {
		if action.GetError() != nil {
			return mergeError(action.GetError())
		}
	}
	onCreateOrMatchEvents := make([]Visitable, 0)
	onCreateOrMatchEvents = append(onCreateOrMatchEvents, BLANK)
	for _, mergeAction := range mergeActions {
		onCreateOrMatchEvents = append(onCreateOrMatchEvents, mergeAction)
	}
	m := merge{
		pattern:               pattern,
		onCreateOrMatchEvents: onCreateOrMatchEvents,
		notNil:                true,
	}
	m.key = getAddress(&m)
	return m
}

func mergeError(err error) merge {
	return merge{err: err}
}

func (m merge) hasEvent() bool {
	return len(m.onCreateOrMatchEvents) > 0
}

func (m merge) GetError() error {
	return m.err
}

func (m merge) accept(visitor *CypherRenderer) {
	visitor.enter(m)
	m.pattern.accept(visitor)
	for _, event := range m.onCreateOrMatchEvents {
		event.accept(visitor)
	}
	visitor.leave(m)
}

func (m merge) enter(renderer *CypherRenderer) {
	renderer.append("MERGE ")
}

func (m merge) leave(renderer *CypherRenderer) {
	if !m.hasEvent() {
		renderer.append(" ")
	}
}

func (m merge) getKey() string {
	return m.key
}

func (m merge) isNotNil() bool {
	return m.notNil
}

func (m merge) isUpdatingClause() bool {
	panic("implement me")
}
