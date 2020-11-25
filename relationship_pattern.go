package cypher_go_dsl

import "strings"

type RelationshipPattern interface {
	ExposesRelationship
	PatternElement
}



/**
RelationshipTypes
 */
type RelationshipTypes struct {
	values []string
}

func (r RelationshipTypes) Accept(visitor *CypherRenderer) {
	(*visitor).Enter(r)
	(*visitor).Leave(r)
}

func (r RelationshipTypes) GetType() VisitableType {
	return RelationshipTypesVisitable
}

func (r RelationshipTypes) Enter(renderer *CypherRenderer) {
	typeWithPrefix := make([]string, 0)
	for _, typeRaw := range r.values {
		if typeRaw == ""{
			continue
		}
		typeWithPrefix = append(typeWithPrefix, REL_TYPE_START + typeRaw)
	}
	renderer.builder.WriteString(strings.Join(typeWithPrefix, REL_TYP_SEPARATOR))
}

func (r RelationshipTypes) Leave(renderer *CypherRenderer) {
	panic("implement me")
}

/**
Relationship length
 */

type RelationshipLength struct {
	minimum *int8
	maximum *int8
	unbounded bool
}

func (relationshipLength RelationshipLength)Unbounded() RelationshipLength {
	return RelationshipLength{nil, nil, true}
}




type Direction struct {
	symbolLeft string
	symbolRight string
}

func LTR() Direction {
	return Direction{symbolLeft: "-", symbolRight: "->"}
}

func RTL() Direction {
	return Direction{symbolLeft: "<-", symbolRight: "-"}
}

func UNI() Direction {
	return Direction{symbolLeft: "-", symbolRight: "-"}
}


func CreateLTR(left Node, direction Direction, right Node, types ...string) Relationship {
	typeSlice := make([]string, 0)
	typeSlice = append(typeSlice, types...)
	relationshipTypes := RelationshipTypes{values: typeSlice}
	details := RelationshipDetails{
		direction:    &direction,
		symbolicName: nil,
		types:        &relationshipTypes,
	}
	return Relationship{
		left:    &left,
		right:   &right,
		details: &details,
	}
}



