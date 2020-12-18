package cypher_go_dsl

import "errors"

type NamePath struct {
	name    SymbolicName
	pattern Visitable
	key     string
	notNil  bool
	err     error
}

func NamePathCreate(name SymbolicName, pattern RelationshipPattern) NamePath {
	if name.getError() != nil {
		return NamePathError(name.getError())
	}
	if pattern != nil && pattern.getError() != nil {
		return NamePathError(pattern.getError())
	}
	n := NamePath{name: name, pattern: pattern}
	n.key = getAddress(&n)
	return n
}

func NamePathCreate1(name SymbolicName, algorithm FunctionInvocation) NamePath {
	if name.getError() != nil {
		return NamePathError(name.getError())
	}
	if algorithm.getError() != nil {
		return NamePathError(algorithm.getError())
	}
	n := NamePath{name: name, pattern: algorithm}
	n.key = getAddress(&n)
	return n
}

func NamePathError(err error) NamePath {
	return NamePath{err: err}
}

func NamePathBuilderWithNameByString(name string) OngoingDefinitionWithName {
	if name == "" {
		return NamePathBuilder{
			err: errors.New("a name is required"),
		}
	}
	return NamePathBuilderWithName(SymbolicNameCreate(name))
}

func NamePathBuilderWithName(name SymbolicName) OngoingDefinitionWithName {
	if !name.isNotNil() {
		return NamePathBuilder{
			err: errors.New("a name is required"),
		}
	}
	return NamePathBuilder{name: name}
}

func (n NamePath) getError() error {
	return n.err
}

func (n NamePath) getRequiredSymbolicName() SymbolicName {
	if n.name.isNotNil() {
		return n.name
	}
	return SymbolicNameError(errors.New("no name present"))
}

func (n NamePath) getSymbolicName() SymbolicName {
	return n.name
}

func (n NamePath) IsPatternElement() bool {
	return true
}

func (n NamePath) accept(visitor *CypherRenderer) {
	visitor.enter(n)
	n.name.accept(visitor)
	ASSIGMENT.accept(visitor)
	n.pattern.accept(visitor)
	visitor.leave(n)
}

func (n NamePath) enter(renderer *CypherRenderer) {
}

func (n NamePath) leave(renderer *CypherRenderer) {
}

func (n NamePath) getKey() string {
	return n.key
}

func (n NamePath) isNotNil() bool {
	return n.notNil
}

//Interface
type OngoingDefinitionWithName interface {
	definedByRelationPattern(pattern RelationshipPattern) NamePath
	getError() error
}

type OngoingShortestPathDefinitionWithName interface {
	definedBy(relationship Relationship) NamePath
}

type NamePathBuilder struct {
	name SymbolicName
	err  error
}

func NamePathBuilderCreate(name SymbolicName) NamePathBuilder {
	return NamePathBuilder{name: name}
}

func (n NamePathBuilder) definedByRelationPattern(pattern RelationshipPattern) NamePath {
	return NamePathCreate(n.name, pattern)
}

func (n NamePathBuilder) getError() error {
	return n.err
}

type ShortestPathBuilder struct {
	name      SymbolicName
	algorithm FunctionDefinition
}

func ShortestPathBuilderCreate(name SymbolicName, algorithm FunctionDefinition) ShortestPathBuilder {
	return ShortestPathBuilder{
		name:      name,
		algorithm: algorithm,
	}
}

func (s ShortestPathBuilder) definedBy(relationship Relationship) NamePath {
	return NamePathCreate1(s.name, FunctionInvocationCreateWithPatternElement(s.algorithm, relationship))
}
