package cypher

import (
	"strings"
)
import uuid "github.com/google/uuid"

type CypherRenderer struct {
	visitableToAliased     map[string]AliasedExpression
	separatorOnLevel       map[int]string
	resolvedSymbolicNames  map[SymbolicName]string
	visitedNamed           map[string]int
	currentVisitedElements []string
	currentAliasedElements []AliasedExpression
	currentLevel           int
	skipNodeContent        bool
	builder                strings.Builder
}

func (renderer *CypherRenderer) append(content string) *CypherRenderer {
	renderer.builder.WriteString(content)
	return renderer
}

func NewRenderer() *CypherRenderer {
	return &CypherRenderer{
		visitableToAliased:     make(map[string]AliasedExpression),
		separatorOnLevel:       make(map[int]string),
		resolvedSymbolicNames:  make(map[SymbolicName]string),
		visitedNamed:           make(map[string]int),
		currentVisitedElements: make([]string, 0),
		currentAliasedElements: make([]AliasedExpression, 0),
		skipNodeContent:        false,
		builder:                strings.Builder{},
	}
}

func (renderer CypherRenderer) Render(statement Statement) string {
	statement.accept(&renderer)
	return renderer.builder.String()
}

func (renderer *CypherRenderer) enableSeparator(level int, on bool) {
	if on {
		renderer.separatorOnLevel[level] = ""
	} else {
		delete(renderer.separatorOnLevel, level)
	}
}

func (renderer *CypherRenderer) separatorOnCurrentLevel() (string, bool) {
	value, isExist := renderer.separatorOnLevel[renderer.currentLevel]
	return value, isExist
}

func (renderer *CypherRenderer) resetSeparatorOnCurrentLevel() {
	renderer.separatorOnLevel[renderer.currentLevel] = ""
}

func (renderer *CypherRenderer) resolve(name SymbolicName) string {
	if _, isExist := renderer.resolvedSymbolicNames[name]; !isExist {
		value := name.value
		if len(value) > 0 {
			return value
		}
		return uuid.New().String()
	}
	return renderer.resolvedSymbolicNames[name]
}

func (renderer *CypherRenderer) enter(visitable Visitable) {
	used := renderer.getAliasedIfSeen(visitable)
	if renderer.PreEnter(&used) {
		renderer.currentVisitedElements = push(renderer.currentVisitedElements, visitable.getKey())
		visitable.enter(renderer)
	}
}

func (renderer *CypherRenderer) leave(visitable Visitable) {
	used := renderer.getAliasedIfSeen(visitable)
	current := peek(renderer.currentVisitedElements)
	if current == (used).getKey() {
		visitable.leave(renderer)
		renderer.postLeave(used)
		renderer.currentVisitedElements = pop(renderer.currentVisitedElements)
	}
	aliasedExpression, isAliased := (visitable).(AliasedExpression)
	if isAliased {
		renderer.visitableToAliased[aliasedExpression.delegate.getKey()] = aliasedExpression
	}
}

func (renderer *CypherRenderer) PreEnter(visitable *Visitable) bool {
	lastAliased := peekAliased(renderer.currentAliasedElements)
	visited := false
	if lastAliased != nil {
		_, visited = renderer.visitableToAliased[lastAliased.getKey()]
	}
	if renderer.skipNodeContent || visited {
		return false
	}
	aliasedExpression, isAliased := (*visitable).(AliasedExpression)
	if isAliased {
		renderer.currentAliasedElements = pushAliased(renderer.currentAliasedElements, aliasedExpression)
	}
	renderer.currentLevel = renderer.currentLevel + 1
	nextLevel := renderer.currentLevel + 1
	if _, isSub := (*visitable).(SubVisitable); isSub {
		renderer.enableSeparator(nextLevel, true)
	}
	if separator, isSeparator := renderer.separatorOnCurrentLevel(); isSeparator {
		renderer.append(separator)
		renderer.resetSeparatorOnCurrentLevel()
	}
	return !renderer.skipNodeContent
}

func (renderer *CypherRenderer) postLeave(visitable Visitable) {

	if _, isSeparator := renderer.separatorOnCurrentLevel(); isSeparator {
		renderer.separatorOnLevel[renderer.currentLevel] = ", "
	}
	if _, isSub := visitable.(SubVisitable); isSub {
		renderer.enableSeparator(renderer.currentLevel+1, false)
	}
	if peekAliased(renderer.currentAliasedElements) != nil &&
		peekAliased(renderer.currentAliasedElements).getKey() == visitable.getKey() {
		popAliased(renderer.currentAliasedElements)
	}
	renderer.currentLevel = renderer.currentLevel - 1
}

func (renderer *CypherRenderer) getAliasedIfSeen(visitable Visitable) Visitable {
	if expression, isExist := renderer.visitableToAliased[visitable.getKey()]; isExist {
		return expression
	}
	return visitable
}

const RelTypSeparator = "|"
const RelTypeStart = ":"
const NodeLabelStart = ":"

func enterRelationshipTypes(renderer *CypherRenderer, visitable Visitable) {
	relationshipTypes := visitable.(RelationshipTypes)
	typeWithPrefix := make([]string, 0)
	for _, typeRaw := range relationshipTypes.values {
		if typeRaw == "" {
			continue
		}
		typeWithPrefix = append(typeWithPrefix, RelTypeStart+typeRaw)
	}
	renderer.append(strings.Join(typeWithPrefix, RelTypSeparator))
}

func enterRelationshipDetail(renderer *CypherRenderer, visitable Visitable) {
	details := visitable.(RelationshipDetails)
	direction := details.direction
	renderer.append(direction.symbolLeft)
	if details.hasContent() {
		renderer.append("[")
	}
}

func leaveRelationshipDetail(renderer *CypherRenderer, visitable Visitable) {
	details := visitable.(RelationshipDetails)
	direction := details.direction
	if details.hasContent() {
		renderer.append("]")
	}
	renderer.append(direction.symbolRight)
}

func enterSkip(renderer *CypherRenderer, visitable Visitable) {
	renderer.append(" SKIP ")
}

func enterLimit(renderer *CypherRenderer, visitable Visitable) {
	renderer.append(" LIMIT ")
}

func enterLiteral(renderer *CypherRenderer, visitable Visitable) {
	literal := visitable.(Literal)
	renderer.append(literal.AsString())
}

func push(queue []string, element string) []string {
	queue = append(queue, element)
	return queue
}

func pop(queue []string) []string {
	if len(queue) > 0 {
		return queue[:len(queue)-1]
	}
	return queue
}

func peek(queue []string) string {
	if len(queue) > 0 {
		return queue[len(queue)-1]
	}
	return ""
}

func pushAliased(queue []AliasedExpression, element AliasedExpression) []AliasedExpression {
	queue = append(queue, element)
	return queue
}

func popAliased(queue []AliasedExpression) []AliasedExpression {
	if len(queue) > 0 {
		return queue[:len(queue)-1]
	}
	return queue
}

func peekAliased(queue []AliasedExpression) *AliasedExpression {
	if len(queue) > 0 {
		return &queue[len(queue)-1]
	}
	return nil
}
