package cypher

import "errors"

type Property struct {
	ExpressionContainer
	container Expression
	name      PropertyLookup
	key       string
	notNil    bool
	err       error
}

func PropertyCreate(parentContainer Named, name string) Property {
	if parentContainer != nil && parentContainer.GetError() != nil {
		return PropertyError(parentContainer.GetError())
	}
	if parentContainer == nil {
		return PropertyError(errors.New("node or relationship is nil"))
	}
	requiredSymbolicName := parentContainer.GetRequiredSymbolicName()
	if requiredSymbolicName.GetError() != nil {
		return PropertyError(requiredSymbolicName.GetError())
	}
	if name == "" {
		return PropertyError(errors.New("property name is required"))
	}
	return PropertyCreate1(requiredSymbolicName, PropertyLookupCreate(name))
}

func PropertyCreate2(container Expression, name string) Property {
	if container != nil && container.GetError() != nil {
		return PropertyError(container.GetError())
	}
	if container == nil {
		return PropertyError(errors.New("container is nil"))
	}
	if name == "" {
		return PropertyError(errors.New("property name is required"))
	}
	return PropertyCreate1(container, PropertyLookupCreate(name))
}

func PropertyCreate1(container Expression, name PropertyLookup) Property {
	if container != nil && container.GetError() != nil {
		return PropertyError(container.GetError())
	}
	if name.GetError() != nil {
		return PropertyError(name.GetError())
	}
	property := Property{
		container: container,
		name:      name,
		notNil:    true,
	}
	property.key = getAddress(&property)
	property.ExpressionContainer = ExpressionWrap(property)
	return property
}

func PropertyError(err error) Property {
	property := Property{err: err}
	property.ExpressionContainer = ExpressionWrap(property)
	return property
}

func (p Property) To(expression Expression) Operation {
	return OperationSet(p, expression)
}

func (p Property) GetError() error {
	return p.err
}

func (p Property) isNotNil() bool {
	return p.notNil
}

func (p Property) accept(visitor *CypherRenderer) {
	visitor.enter(p)
	p.container.accept(visitor)
	p.name.accept(visitor)
	visitor.leave(p)
}

func (p Property) enter(renderer *CypherRenderer) {
}

func (p Property) leave(renderer *CypherRenderer) {
}

func (p Property) getKey() string {
	return p.key
}

func (p Property) GetExpressionType() ExpressionType {
	return "Property"
}

func (p Property) getName() PropertyLookup {
	return p.name
}
