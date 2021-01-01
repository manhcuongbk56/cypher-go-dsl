package cypher

import "errors"

type NamedPath struct {
	name    SymbolicName
	pattern Visitable
	key     string
	notNil  bool
	err     error
}

func NamedPathCreate(name SymbolicName, pattern RelationshipPattern) NamedPath {
	if name.getError() != nil {
		return NamedPathError(name.getError())
	}
	if pattern != nil && pattern.getError() != nil {
		return NamedPathError(pattern.getError())
	}
	n := NamedPath{name: name, pattern: pattern}
	n.key = getAddress(&n)
	return n
}

func NamedPathCreate1(name SymbolicName, algorithm FunctionInvocation) NamedPath {
	if name.getError() != nil {
		return NamedPathError(name.getError())
	}
	if algorithm.getError() != nil {
		return NamedPathError(algorithm.getError())
	}
	n := NamedPath{name: name, pattern: algorithm}
	n.key = getAddress(&n)
	return n
}

func NamedPathError(err error) NamedPath {
	return NamedPath{err: err}
}

func NamedPathBuilderWithNameByString(name string) OngoingDefinitionWithName {
	if name == "" {
		return NamedPathBuilder{
			err: errors.New("a name is required"),
		}
	}
	return NamedPathBuilderWithName(SymbolicNameCreate(name))
}

func NamedPathBuilderWithName(name SymbolicName) OngoingDefinitionWithName {
	if !name.isNotNil() {
		return NamedPathBuilder{
			err: errors.New("a name is required"),
		}
	}
	return NamedPathBuilder{name: name}
}

func NamedPathShortestPathWithNameByString(name string, algorithm FunctionDefinition) OngoingShortestPathDefinitionWithName {
	if name == "" {
		return ShortestPathBuilder{
			err: errors.New("a name is required"),
		}
	}
	return NamedPathShortestPathWithName(SymbolicNameCreate(name), algorithm)
}

func NamedPathShortestPathWithName(name SymbolicName, algorithm FunctionDefinition) OngoingShortestPathDefinitionWithName {
	if !name.isNotNil() {
		return ShortestPathBuilder{
			err: errors.New("a name is required"),
		}
	}
	return ShortestPathBuilder{
		name:      name,
		algorithm: algorithm,
	}
}

func (n NamedPath) getError() error {
	return n.err
}

func (n NamedPath) getRequiredSymbolicName() SymbolicName {
	if n.name.isNotNil() {
		return n.name
	}
	return SymbolicNameError(errors.New("namepath get symbolic name: no name present"))
}

func (n NamedPath) getSymbolicName() SymbolicName {
	return n.name
}

func (n NamedPath) IsPatternElement() bool {
	return true
}

func (n NamedPath) accept(visitor *CypherRenderer) {
	visitor.enter(n)
	n.name.accept(visitor)
	ASSIGMENT.accept(visitor)
	n.pattern.accept(visitor)
	visitor.leave(n)
}

func (n NamedPath) enter(renderer *CypherRenderer) {
}

func (n NamedPath) leave(renderer *CypherRenderer) {
}

func (n NamedPath) getKey() string {
	return n.key
}

func (n NamedPath) isNotNil() bool {
	return n.notNil
}

//Interface
type OngoingDefinitionWithName interface {
	definedByRelationPattern(pattern RelationshipPattern) NamedPath
	getError() error
}

type OngoingShortestPathDefinitionWithName interface {
	DefinedBy(relationship Relationship) NamedPath
	getError() error
}

type NamedPathBuilder struct {
	name SymbolicName
	err  error
}

func NamedPathBuilderCreate(name SymbolicName) NamedPathBuilder {
	return NamedPathBuilder{name: name}
}

func (n NamedPathBuilder) definedByRelationPattern(pattern RelationshipPattern) NamedPath {
	return NamedPathCreate(n.name, pattern)
}

func (n NamedPathBuilder) getError() error {
	return n.err
}

type ShortestPathBuilder struct {
	name      SymbolicName
	algorithm FunctionDefinition
	err       error
}

func ShortestPathBuilderCreate(name SymbolicName, algorithm FunctionDefinition) ShortestPathBuilder {
	return ShortestPathBuilder{
		name:      name,
		algorithm: algorithm,
	}
}

func (s ShortestPathBuilder) DefinedBy(relationship Relationship) NamedPath {
	if s.err != nil {
		return NamedPathError(s.err)
	}
	if relationship.err != nil {
		return NamedPathError(relationship.err)
	}
	return NamedPathCreate1(s.name, FunctionInvocationCreateWithPatternElement(s.algorithm, relationship))
}

func (s ShortestPathBuilder) getError() error {
	return s.err
}
