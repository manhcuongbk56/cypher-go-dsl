package cypher

import (
	"errors"
	"strings"
)

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

func NewNode(primaryLabel string) Node {
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

/**
 * Starts defining a named path by indicating a name.
 *
 * @param name The name of the new path
 * @return An ongoing definition of a named path
 * @since 1.1
 */
func CypherPathByString(name string) OngoingDefinitionWithName {
	return NamedPathBuilderWithNameByString(name)
}

/**
 * Starts defining a named path by indicating a name.
 *
 * @param name The name of the new path
 * @return An ongoing definition of a named path
 * @since 1.1
 */
func CypherPath(name SymbolicName) OngoingDefinitionWithName {
	return NamedPathBuilderWithName(name)
}

/**
 * Starts defining a named path defined by the {@code shortestPath} between a relationship by indicating a name.
 *
 * @param name The name of the new shortestPath path
 * @return An ongoing definition of a named path
 * @since 1.1.1
 */
func CypherShortestPathByString(name string) OngoingShortestPathDefinitionWithName {
	return NamedPathShortestPathWithNameByString(name, SHORTEST_PATH)
}

/**
 * Starts defining a named path defined by the {@code shortestPath} between a relationship by indicating a name.
 *
 * @param name The name of the new shortestPath path
 * @return An ongoing definition of a named path
 * @since 1.1.1
 */
func CypherShortestPath(name SymbolicName) OngoingShortestPathDefinitionWithName {
	return NamedPathShortestPathWithName(name, SHORTEST_PATH)
}

/**
 * Creates a new symbolic name.
 *
 * @param value The value of the symbolic name
 * @return A new symbolic name
 */
func CypherName(value string) SymbolicName {
	return SymbolicNameCreate(value)
}

/**
 * Creates a new parameter placeholder. Existing $-signs will be removed.
 *
 * @param name The name of the parameter, must not be null
 * @return The new parameter
 */
func CypherParameter(name string) Parameter {
	return ParameterCreate(name)
}

/**
 * Prepares an optional match statement.
 *
 * @param pattern The patterns to match
 * @return An ongoing match that is used to specify an optional where and a required return clause
 */
func CypherOptionalMatch(element ...PatternElement) OngoingReadingWithoutWhere {
	return DefaultStatementBuilderCreate().OptionalMatch(element...)
}

/**
 * Starts building a statement based on a match clause. Use {@link Cypher#node(String, String...)} and related to
 * retrieve a node or a relationship, which both are pattern elements.
 *
 * @param pattern The patterns to match
 * @return An ongoing match that is used to specify an optional where and a required return clause
 */
func MatchElements(element ...PatternElement) OngoingReadingWithoutWhere {
	return DefaultStatementBuilderCreate().Match(element...)
}

/**
 * Starts building a statement based on a match clause. Use {@link Cypher#node(String, String...)} and related to
 * retrieve a node or a relationship, which both are pattern elements.
 *
 * @param optional A flag whether the {@code MATCH} clause includes the {@code OPTIONAL} keyword.
 * @param pattern  The patterns to match
 * @return An ongoing match that is used to specify an optional where and a required return clause
 * @since 2020.1.3
 */
func CypherMatchDefault(optional bool, element ...PatternElement) OngoingReadingWithoutWhere {
	return DefaultStatementBuilderCreate().MatchDefault(optional, element...)
}

/**
 * Starts building a statement based on a {@code CREATE} clause.
 *
 * @param pattern The patterns to create
 * @param <T>     The type of the next step
 * @return An ongoing {@code CREATE} that can be used to specify {@code WITH} and {@code RETURNING} etc.
 */
func CypherCreate(patterns ...PatternElement) OngoingUpdateAndExposesSet {
	return DefaultStatementBuilderCreate().Create(patterns...)
}

/**
 * Starts a statement with a leading {@code WITH}. Those are useful for passing on lists of various type that
 * can be unwound later on etc. A leading {@code WITH} cannot be used with patterns obviously and needs its
 * arguments to have an alias.
 *
 * @param variables One ore more variables.
 * @return An ongoing with clause.
 * @since 2020.1.2
 */
func CypherWithByString(variables ...string) OrderableOngoingReadingAndWithWithoutWhere {
	return DefaultStatementBuilderCreate().WithByString(variables...)
}

/**
 * Starts a statement with a leading {@code WITH}. Those are useful for passing on lists of various type that
 * can be unwound later on etc. A leading {@code WITH} cannot be used with patterns obviously and needs its
 * arguments to have an alias.
 *
 * @param variables One ore more variables.
 * @return An ongoing with clause.
 * @since 2020.1.2
 */
func CypherWithByNamed(variables ...Named) OrderableOngoingReadingAndWithWithoutWhere {
	return DefaultStatementBuilderCreate().WithByNamed(variables...)
}

/**
 * Starts a statement with a leading {@code WITH}. Those are useful for passing on lists of various type that
 * can be unwound later on etc. A leading {@code WITH} cannot be used with patterns obviously and needs its
 * arguments to have an alias.
 *
 * @param expressions One ore more aliased expressions.
 * @return An ongoing with clause.
 */
func CypherWith(variables ...Expression) OrderableOngoingReadingAndWithWithoutWhere {
	return DefaultStatementBuilderCreate().With(variables...)
}

/**
 * Starts building a statement based on a {@code MERGE} clause.
 *
 * @param pattern The patterns to merge
 * @param <T>     The type of the next step
 * @return An ongoing {@code MERGE} that can be used to specify {@code WITH} and {@code RETURNING} etc.
 */
func CypherMerge(patterns ...PatternElement) OngoingUpdateAndExposesSet {
	return DefaultStatementBuilderCreate().Merge(patterns...)
}

/**
 * Starts building a statement starting with an {@code UNWIND} clause. The expression needs to be an expression
 * evaluating to a list, otherwise the query will fail.
 *
 * @param expression The expression to unwind
 * @return An ongoing {@code UNWIND}.
 */
func CypherUnwind(expression Expression) OngoingUnwind {
	return DefaultStatementBuilderCreate().Unwind(expression)
}

/**
 * Starts building a statement starting with an {@code UNWIND} clause. The expressions passed will be turned into a
 * list expression
 *
 * @param expressions expressions to unwind
 * @return a new instance of {@link StatementBuilder.OngoingUnwind}
 */
func CypherUnwindMulti(expressions ...Expression) OngoingUnwind {
	return DefaultStatementBuilderCreate().Unwind(CypherListOf(expressions...))
}

/**
 * Creates a new {@link SortItem} to be used as part of an {@link Order}.
 *
 * @param expression The expression by which things should be sorted
 * @return A sort item, providing means to specify ascending or descending order
 */
func CypherSort(expression Expression) SortItem {
	return SortItemCreate(expression, UNDEFINED)
}

/**
 * Creates a map of expression from a list of key/value pairs.
 *
 * @param keysAndValues A list of key and values. Must be an even number, with alternating {@link String} and {@link Expression}
 * @return A new map expression.
 */
func MapOf(objects ...interface{}) MapExpression {
	return NewMapExpression(objects...)
}

/**
 * Creates a {@link ListExpression list-expression} from several expressions.
 *
 * @param expressions expressions to get combined into a list
 * @return a new instance of {@link ListExpression}
 */
func CypherListOf(expressions ...Expression) ExpressionList {
	return ExpressionListCreate(expressions)
}

/**
 * @return The {@literal true} literal.
 */
func CypherLiteralTrue() BooleanLiteral {
	return TRUE
}

/**
 * @return The {@literal false} literal.
 */
func CypherLiteralFalse() BooleanLiteral {
	return FALSE
}

/**
 * Creates a {@code UNION} statement from several other statements. No checks are applied for matching return types.
 *
 * @param statements the statements to union.
 * @return A union statement.
 */
func CypherUnion(statements ...Statement) Statement {
	return unionImpl(false, statements...)
}

/**
 * Creates a {@code UNION ALL} statement from several other statements. No checks are applied for matching return types.
 *
 * @param statements the statements to union.
 * @return A union statement.
 */
func CypherUnionAll(statements ...Statement) Statement {
	return unionImpl(true, statements...)
}

/**
 * A {@literal RETURN} statement without a previous match.
 *
 * @param expressions The expressions to return
 * @return A buildable statement
 * @since 1.0.1
 */
func CypherReturning(expressions ...Expression) OngoingReadingAndReturn {
	return DefaultStatementBuilderCreate().Returning(expressions...)
}

/**
 * Creates a list comprehension starting with a {@link Relationship} or a {@link RelationshipChain chain of relationships}.
 *
 * @param relationshipPattern The relationship pattern on which the new list comprehension is based on.
 * @return An ongoing definition.
 * @since 2020.0.0
 */
func CypherListBasedOn(pattern RelationshipPattern) PatternComprehensionOngoingDefinitionWithPattern {
	return PatternComprehensionBasedOn(pattern)
}

/**
 * Creates a list comprehension starting with a {@link NamedPath named path}.
 *
 * @param namedPath The named path on which the new list comprehension is based on.
 * @return An ongoing definition.
 * @since 2020.1.1
 */
func CypherListBasedOnNamed(namedPath NamedPath) PatternComprehensionOngoingDefinitionWithPattern {
	return PatternComprehensionBasedOnNamePath(namedPath)
}

/**
 * Starts defining a {@link ListComprehension list comprehension}.
 *
 * @param variable The variable to which each element of the list is assigned.
 * @return An ongoing definition of a list comprehension
 * @since 1.0.1
 */
func CypherListWith(variable SymbolicName) OngoingDefinitionWithVariable {
	return ListComprehensionBuilderCreate(variable)
}

/**
 * Escapes and quotes the {@code unquotedString} for safe usage in Neo4j-Browser and Shell.
 *
 * @param unquotedString An unquoted string
 * @return A quoted string with special chars escaped.
 */
func CypherQuote(unquotedString string) string {
	return StringLiteralCreate(unquotedString).AsString()
}

/**
 * @return generic case expression start
 */
func CypherCaseExpression() Case {
	return GenericCaseCreate1()
}

/**
 * @param expression initial expression for the simple case statement
 * @return simple case expression start
 */
func CypherCaseExpression1(expression Expression) Case {
	return SimpleCaseCreate1(expression)
}

func CypherCallSimple(procedureName string) OngoingStandaloneCallWithoutArguments {
	if procedureName == "" {
		return StandaloneCallBuilderError(errors.New("the procedure name must not be nil or empty"))
	}
	return CypherCall(strings.Split(procedureName, "\\.")...)
}

/**
 * Starts defining a procedure call of the procedure with the given qualified name.
 *
 * @param namespaceAndProcedure The procedure name of the procedure to call.
 * @return An ongoing definition of a call
 */
func CypherCall(namespaceAndProcedure ...string) OngoingStandaloneCallWithoutArguments {
	return StatementCall(namespaceAndProcedure...)
}

/**
 * Starts building a statement based on one subquery.
 *
 * @param subquery The statement representing the subquery
 * @neo4j.version 4.0.0
 * @see ExposesSubqueryCall#call(Statement)
 * @since 2020.1.2
 * @return A new ongoing read without any further conditions or returns.
 */
func CypherCallByStatement(subquery Statement) OngoingReadingWithoutWhere {
	return DefaultStatementBuilderCreate().Call(subquery)
}

/**
 * Creates a closed range with given boundaries.
 *
 * @param targetExpression The target expression for the range
 * @param start            The inclusive start
 * @param end              The exclusive end
 * @return A range literal.
 * @since 2020.1.0
 */
func CypherSubList(targetExpression Expression, start int, end int) Expression {
	return SubList(targetExpression, NumberLiteralCreate1(start), NumberLiteralCreate1(end))
}

/**
 * Creates an open range starting at {@code start}.
 *
 * @param targetExpression The target expression for the range
 * @param start            The inclusive start
 * @return A range literal.
 * @since 2020.1.0
 */
func CypherSubListFrom(targetExpression Expression, start int) Expression {
	return SubListFrom(targetExpression, NumberLiteralCreate1(start))
}

/**
 * Creates an open range starting at {@code start}.
 *
 * @param targetExpression The target expression for the range
 * @param start            The inclusive start
 * @return A range literal.
 * @since 2020.1.0
 */
func CypherSubListFromByExpression(targetExpression Expression, start Expression) Expression {
	return SubListFrom(targetExpression, start)
}

/**
 * Creates an open range starting at {@code start}.
 *
 * @param targetExpression The target expression for the range
 * @param end              The exclusive end
 * @return A range literal.
 * @since 2020.1.0
 */
func CypherSubListUntil(targetExpression Expression, end int) Expression {
	return SubListUntil(targetExpression, NumberLiteralCreate1(end))
}

/**
 * Creates an open range starting at {@code start}.
 *
 * @param targetExpression The target expression for the range
 * @param end              The exclusive end
 * @return A range literal.
 * @since 2020.1.0
 */
func CypherSubListUntilByExpression(targetExpression Expression, end Expression) Expression {
	return SubListUntil(targetExpression, end)
}

/**
 * Creates a single valued range at {@code index}.
 *
 * @param targetExpression The target expression for the range
 * @param index            The index of the range
 * @return A range literal.
 * @since 2020.1.0
 */
func CypherSubListValueAt(targetExpression Expression, index int) Expression {
	return ValueAt(targetExpression, NumberLiteralCreate1(index))
}

/**
 * Creates a single valued range at {@code index}.
 *
 * @param targetExpression The target expression for the range
 * @param index            The index of the range
 * @return A range literal.
 * @since 2020.1.0
 */
func CypherSubListValueAtByExpression(targetExpression Expression, index Expression) Expression {
	return ValueAt(targetExpression, index)
}

func LiteralOf(object interface{}) Literal {
	if object == nil {
		return NIL_INSTANCE
	}
	if stringValue, isString := object.(string); isString {
		return StringLiteralCreate(stringValue)
	}
	if intValue, isInt := object.(int); isInt {
		return NumberLiteralCreate1(intValue)
	}
	if literalSlice, isLiteralSlice := object.([]Literal); isLiteralSlice {
		return ListLiteralCreate(literalSlice)
	}
	if booleanValue, isBoolean := object.(bool); isBoolean {
		return BooleanLiteralCreate(booleanValue)
	}
	return StringLiteralError(errors.New("cypher literal of: unsupported literal type"))
}

func escapeName(name string) string {
	return "`" + strings.ReplaceAll(name, "`", "``") + "`"
}

func escapeIfNecessary(name string) string {
	//TODO: maybe need to implement this
	return name
}

func unionImpl(unionAll bool, statements ...Statement) Statement {
	if statements == nil || len(statements) < 2 {
		return UnionQueryError(errors.New("at least 2 statements are required"))
	}
	i := 0
	existingUnionQuery := UnionQuery{}
	if unionQuery, isUnion := statements[0].(UnionQuery); isUnion {
		existingUnionQuery = unionQuery
		if existingUnionQuery.all != unionAll {
			return UnionQueryError(errors.New("cannot mix union and union all"))
		}
		i = 1
	}
	listOfQueries := make([]SingleQuery, 0)
	for _, query := range statements[i:] {
		if singleQuery, isSingle := query.(SingleQuery); isSingle {
			listOfQueries = append(listOfQueries, singleQuery)
			continue
		}
		return UnionQueryError(errors.New("can only union single queries"))
	}
	if !existingUnionQuery.notNil {
		return UnionQueryCreate(unionAll, listOfQueries)
	} else {
		return existingUnionQuery.addAdditionalQueries(listOfQueries)
	}
}
