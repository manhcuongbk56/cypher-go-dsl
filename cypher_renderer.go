package cypher_go_dsl

import "strings"
import uuid "github.com/google/uuid"

type CypherRenderer struct {
	visitableEnterMap map[VisitableType]enter
	visitableLeaveMap map[VisitableType]leave
	visitableToAliased map[Visitable]AliasedExpression
	separatorOnLevel map[int]string
	resolvedSymbolicNames map[SymbolicName]string
	visitedNamed map[Named]int
	currentVisitedElements []Visitable
	currentAliasedElements []AliasedExpression
	currentLevel int
	skipNodeContent bool
	builder strings.Builder
}

var enterMap = map[VisitableType]enter {
	MatchVisitable: enterMatch,
}

var leaveMap = map[VisitableType]leave {
MatchVisitable: leaveMatch,
}

func NewRenderer() CypherRenderer {
	return CypherRenderer{
		visitableEnterMap: enterMap,
		visitableLeaveMap: leaveMap,
		visitableToAliased: make(map[Visitable]AliasedExpression),
		separatorOnLevel: make(map[int]string),
		resolvedSymbolicNames: make(map[SymbolicName]string),
		visitedNamed: make(map[Named]int),
		currentVisitedElements: make([]Visitable, 0),
		currentAliasedElements: make([]AliasedExpression, 0),
		skipNodeContent: false,
		builder: strings.Builder{},
	}
}


type enter func(visitor CypherRenderer, visitable Visitable)
type leave func(visitor CypherRenderer, visitable Visitable)

func (renderer CypherRenderer) Render(statement Statement) string {
	statement.Accept(renderer)
	return renderer.builder.String()
}

func (renderer CypherRenderer) enableSeparator(level int, on bool)  {
	if on {
		renderer.separatorOnLevel[level] = ""
	}else {
		delete(renderer.separatorOnLevel, level)
	}
}

func (renderer CypherRenderer) separatorOnCurrentLevel() (string, bool) {
	value, isExist := renderer.separatorOnLevel[renderer.currentLevel]
	return value, isExist
}

func (renderer CypherRenderer) resetSeparatorOnCurrentLevel() {
	renderer.separatorOnLevel[renderer.currentLevel] = ""
}

func (renderer CypherRenderer) resolve(name SymbolicName) string {
	if _, isExist := renderer.resolvedSymbolicNames[name]; !isExist {
		value := name.Value
		if len(value) > 0 {
			return value
		}
		return uuid.New().String()
	}
	return renderer.resolvedSymbolicNames[name]
}

func (renderer CypherRenderer) ExecuteEnter(visitable Visitable)  {
	renderer.visitableEnterMap[visitable.GetType()](renderer, visitable)
}

func (renderer CypherRenderer) ExecuteLeave(visitable Visitable)  {
	renderer.visitableLeaveMap[visitable.GetType()](renderer, visitable)
}

func (renderer CypherRenderer) Enter(visitable Visitable) {
	used := renderer.getAliasIfSeen(visitable)
	if renderer.PreEnter(visitable){
		renderer.currentVisitedElements = push(renderer.currentVisitedElements, used)
		renderer.ExecuteEnter(used)
	}
}

func (renderer CypherRenderer) Leave(visitable Visitable) {
	used := renderer.getAliasIfSeen(visitable)
	if peek(renderer.currentVisitedElements) == used {
		renderer.ExecuteLeave(used)
		renderer.PostLeave(used)
		renderer.currentVisitedElements = pop(renderer.currentVisitedElements)
	}
	aliasedExpression, isAliased := visitable.(AliasedExpression)
	if isAliased {
		renderer.visitableToAliased[aliasedExpression.delegate] = aliasedExpression
	}
}

func (renderer CypherRenderer) PreEnter(visitable Visitable) bool {
	lastAliased := peekAliased(renderer.currentAliasedElements)
	_, visited := renderer.visitableToAliased [lastAliased]
	if renderer.skipNodeContent || visited  {
		return false
	}
	aliasedExpression, isAliased := visitable.(AliasedExpression)
	if isAliased {
		renderer.currentAliasedElements = pushAliased(renderer.currentAliasedElements, aliasedExpression)
	}

	nextLevel := renderer.currentLevel + 2
	_, isSub := visitable.(SubVisitable)
	if isSub {
		renderer.enableSeparator(nextLevel, true)
	}
	 seperator, isSeperator := renderer.separatorOnCurrentLevel()
	if isSeperator {
		renderer.builder.WriteString(seperator)
		renderer.resetSeparatorOnCurrentLevel()
	}
	return !renderer.skipNodeContent
}



func (renderer CypherRenderer) PostLeave(visitable Visitable)  {
	_, isSeperator := renderer.separatorOnCurrentLevel()
	if isSeperator {
		renderer.separatorOnLevel[renderer.currentLevel] = ", "
	}

	_, isSub := visitable.(SubVisitable)
	if isSub {
		renderer.enableSeparator(renderer.currentLevel + 1, false)
	}

	if peekAliased(renderer.currentAliasedElements) == visitable {
		popAliased(renderer.currentAliasedElements)
	}
	renderer.currentLevel  = renderer.currentLevel - 1
}


func (renderer CypherRenderer) getAliasIfSeen(visitable Visitable) Visitable {
	if visitable.GetType() != AliasedExpressionVisitable {
		return visitable
	}
	aliasedExpression := visitable.(AliasedExpression)
	if expression, isExist := renderer.visitableToAliased [aliasedExpression]; isExist {
		return expression
	}
	return visitable
}

//Enter And Leave

func enterMatch(renderer CypherRenderer, visitable Visitable ) {
	match := visitable.(Match)
	if match.isOptional() {
		renderer.builder.WriteString("OPTIONAL ")
	}
	renderer.builder.WriteString("MATCH ")
}

func leaveMatch(renderer CypherRenderer, visitable Visitable ) {
	renderer.builder.WriteString(" ")
}

func enterWhere(renderer CypherRenderer, visitable Visitable ) {
	renderer.builder.WriteString(" WHERE ")
}

func leaveWhere(renderer CypherRenderer, visitable Visitable ) {
	renderer.builder.WriteString(" ")
}


func push(queue[] Visitable, element Visitable) []Visitable {
	queue = append(queue, element)
	return queue
}

func pop(queue[] Visitable) []Visitable {
	return queue[1:]
}

func peek(queue[] Visitable) Visitable {
	return queue[0]
}


func pushAliased(queue[] AliasedExpression, element AliasedExpression) []AliasedExpression {
	queue = append(queue, element)
	return queue
}

func popAliased(queue[] AliasedExpression) []AliasedExpression {
	return queue[1:]
}

func peekAliased(queue[] AliasedExpression) AliasedExpression {
	return queue[0]
}
