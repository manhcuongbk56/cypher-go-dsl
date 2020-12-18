package cypher_go_dsl

import "strings"

/**
 * Create a new Node representation with at least one label, the "primary" label. This is required. All other labels
 * are optional.
 *
 * @param primaryLabel     The primary label this node is identified by.
 * @param additionalLabels Additional labels
 * @return A new node representation
 */
func CypherNewNodeWithLabels(primaryLabel string, additionalLabel ...string) Node {
	var labels = make([]NodeLabel, 0)
	labels = append(labels, NodeLabel{value: primaryLabel})
	for _, label := range additionalLabel {
		labels = append(labels, NodeLabel{value: label})
	}
	return Node{
		labels: labels,
	}
}

/**
 * Create a new Node representation with at least one label, the "primary" label. This is required. All other labels
 * are optional. This method also takes a map of properties. This allows the returned node object to be used in a
 * {@code MATCH} or {@code MERGE} statement.
 *
 * @param primaryLabel     The primary label this node is identified by.
 * @param properties       The properties expected to exist on the node.
 * @param additionalLabels Additional labels
 * @return A new node representation
 */
func CypherNewNodeWithProperties(primaryLabel string, properties MapExpression, additionalLabel ...string) Node {
	return NodeCreate5(primaryLabel, properties, additionalLabel...)
}

func CypherNewNode(primaryLabel string) Node {
	return NodeCreate2(primaryLabel)
}

/**
 * @return A node matching any node.
 */
func CypherAnyNode() Node {
	return NodeCreate()
}

/**
 * @return The {@code *} wildcard literal.
 */
func CypherAsterisk() Asterisk {
	return ASTERISK
}

/**
 * @param symbolicName The new symbolic name
 * @return A node matching any node with the symbolic the given {@code symbolicName}.
 */
func CypherAnyNode1(symbolicName string) Node {
	return NodeCreate().NamedByString(symbolicName)
}

/**
 * @param symbolicName The new symbolic name
 * @return A node matching any node with the symbolic the given {@code symbolicName}.
 */
func CypherAnyNode2(symbolicName SymbolicName) Node {
	return NodeCreate().Named(symbolicName)
}

/**
 * Dereferences a property for a symbolic name, most likely pointing to a property container like a node or a relationship.
 *
 * @param containerName The symbolic name of a property container
 * @param name          The name of the property to dereference
 * @return A new property
 */
func CypherProperty(containerName string, name string) Property {
	return CypherPropertyByExpression(CypherName(containerName), name)
}

/**
 * Dereferences a property on a arbitrary expression.
 *
 * @param expression The expression that describes some sort of accessible map
 * @param name       The name of the property to dereference
 * @return A new property.
 */
func CypherPropertyByExpression(expression Expression, name string) Property {
	return PropertyCreate2(expression, name)
}

func Matchs(element ...PatternElement) OngoingReadingWithoutWhere {
	return DefaultStatementBuilderCreate().Match(element...)
}

func MapOf(objects ...interface{}) MapExpression {
	return NewMapExpression(objects...)
}

func Sort(expression Expression) SortItem {
	return SortItemCreate(expression, UNDEFINED)
}

func escapeName(name string) string {
	return "`" + strings.ReplaceAll(name, "`", "``") + "`"
}

func CypherName(value string) SymbolicName {
	return SymbolicNameCreate(value)
}

func ListOf(expressions ...Expression) ExpressionList {
	return ExpressionListCreate(expressions)
}

func Name(value string) SymbolicName {
	return SymbolicNameCreate(value)
}
