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
	if parentContainer != nil && parentContainer.getError() != nil {
		return PropertyError(parentContainer.getError())
	}
	if parentContainer == nil {
		return PropertyError(errors.New("node or relationship is nil"))
	}
	requiredSymbolicName := parentContainer.getRequiredSymbolicName()
	if requiredSymbolicName.getError() != nil {
		return PropertyError(requiredSymbolicName.getError())
	}
	if name == "" {
		return PropertyError(errors.New("property name is required"))
	}
	return PropertyCreate1(requiredSymbolicName, PropertyLookupCreate(name))
}

func PropertyCreate2(container Expression, name string) Property {
	if container != nil && container.getError() != nil {
		return PropertyError(container.getError())
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
	if container != nil && container.getError() != nil {
		return PropertyError(container.getError())
	}
	if name.getError() != nil {
		return PropertyError(name.getError())
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
	return Property{err: err}
}

func (p Property) To(expression Expression) Operation {
	return OperationSet(p, expression)
}

func (p Property) getError() error {
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
