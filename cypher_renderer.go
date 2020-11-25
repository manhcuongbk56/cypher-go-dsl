package cypher_go_dsl

import (
	"strings"
)
import uuid "github.com/google/uuid"

type CypherRenderer struct {
	visitableEnterMap map[VisitableType]enter
	visitableLeaveMap map[VisitableType]leave
	visitableToAliased map[interface{}]Visitable
	separatorOnLevel map[int]string
	resolvedSymbolicNames map[SymbolicName]string
	visitedNamed map[interface{}]int
	currentVisitedElements []*Visitable
	currentAliasedElements []AliasedExpression
	currentLevel int
	skipNodeContent bool
	builder strings.Builder
}

var enterMap = map[VisitableType]enter {
	MatchVisitable: enterMatch,
	NodeVisitable: enterNode,
	WhereVisitable: enterWhere,
	ReturnVisitable: enterReturn,
	PropertiesVisitable: enterProperties,
	SymbolicNameVisitable: enterSymbolic,
	RelationshipDetailsVisitable: enterRelationshipDetail,
	RelationshipTypesVisitable: enterRelationshipTypes,
	LiteralVisitable: enterLiteral,
	LimitVisitable: enterLimit,
	SkipVisitable: enterSkip,
}

var leaveMap = map[VisitableType]leave {
	MatchVisitable: leaveMatch,
	NodeVisitable: leaveNode,
	WhereVisitable: leaveWhere,
	ReturnVisitable: leaveReturn,
	PropertiesVisitable: leaveProperties,
	SymbolicNameVisitable: leaveSymbolic,
	RelationshipDetailsVisitable: leaveRelationshipDetail,

}

func NewRenderer() *CypherRenderer {
	return &CypherRenderer{
		visitableEnterMap: enterMap,
		visitableLeaveMap: leaveMap,
		visitableToAliased: make(map[interface{}]Visitable),
		separatorOnLevel: make(map[int]string),
		resolvedSymbolicNames: make(map[SymbolicName]string),
		visitedNamed: make(map[interface{}]int),
		currentVisitedElements: make([]*Visitable, 0),
		currentAliasedElements: make([]AliasedExpression, 0),
		skipNodeContent: false,
		builder: strings.Builder{},
	}
}


type enter func(visitor *CypherRenderer, visitable Visitable)
type leave func(visitor *CypherRenderer, visitable Visitable)

func (renderer CypherRenderer) Render(statement Statement) string {
	statement.Accept(&renderer)
	return renderer.builder.String()
}

func (renderer *CypherRenderer) enableSeparator(level int, on bool)  {
	if on {
		renderer.separatorOnLevel[level] = ""
	}else {
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
		value := name.Value
		if len(value) > 0 {
			return value
		}
		return uuid.New().String()
	}
	return renderer.resolvedSymbolicNames[name]
}

func (renderer *CypherRenderer) ExecuteEnter(visitable *Visitable)  {
	(*visitable).Enter(renderer)
	//enterFunc, isExist := renderer.visitableEnterMap[visitable.GetType()]
	//if isExist {
	//	enterFunc(&renderer, visitable)
	//	return
	//}
	//fmt.Printf("TYPE IS NOT IN ENTER MAP %d\n", visitable.GetType())
}

func (renderer *CypherRenderer) ExecuteLeave(visitable Visitable)  {
	visitable.Leave(renderer)
	//leaveFunc, isExist := renderer.visitableLeaveMap[visitable.GetType()]
	//if isExist {
	//	leaveFunc(&renderer, visitable)
	//	return
	//}
	//fmt.Printf("TYPE IS NOT IN LEAVE MAP %d\n", visitable.GetType())
}

func (renderer *CypherRenderer) Enter(visitable interface{}) {
	ok := visitable.(Visitable)
	used := renderer.getAliasIfSeen(&ok, nil)
	if renderer.PreEnter(&ok){
		renderer.currentVisitedElements = push(renderer.currentVisitedElements, used)
		renderer.ExecuteEnter(used)
	}
}

func (renderer *CypherRenderer) Leave(visitable interface{}) {
	visitableConverted := visitable.(*Visitable)
	used := renderer.getAliasIfSeen(visitableConverted, nil)
	current := peek(renderer.currentVisitedElements)
	if  current == used {
		renderer.ExecuteLeave(*used)
		renderer.PostLeave(*used)
		renderer.currentVisitedElements = pop(renderer.currentVisitedElements)
	}
	aliasedExpression, isAliased := (*visitableConverted).(AliasedExpression)
	if isAliased {
		renderer.visitableToAliased[aliasedExpression.delegate] = aliasedExpression
	}
}


func (renderer *CypherRenderer) EnterA(visitable Visitable, address interface{}) {
	visitableConverted := visitable.(Visitable)
	used := renderer.getAliasIfSeen(&visitableConverted, address)
	if renderer.PreEnter(&visitableConverted){
		renderer.currentVisitedElements = push(renderer.currentVisitedElements, used)
		renderer.ExecuteEnter(used)
	}
}

func (renderer *CypherRenderer) LeaveA(visitable interface{}, address interface{}) {
	visitableConverted := visitable.(*Visitable)
	used := renderer.getAliasIfSeen(visitableConverted, address)
	current := peek(renderer.currentVisitedElements)
	if  current == used {
		renderer.ExecuteLeave(*used)
		renderer.PostLeave(*used)
		renderer.currentVisitedElements = pop(renderer.currentVisitedElements)
	}
	aliasedExpression, isAliased := (*visitableConverted).(AliasedExpression)
	if isAliased {
		renderer.visitableToAliased[aliasedExpression.delegate] = aliasedExpression
	}
}

func (renderer *CypherRenderer) PreEnter(visitable *Visitable) bool {
	lastAliased := peekAliased(renderer.currentAliasedElements)
	visited := false
	if lastAliased != nil{
		_, visited = renderer.visitableToAliased [*lastAliased]
	}
	if renderer.skipNodeContent || visited  {
		return false
	}
	aliasedExpression, isAliased := (*visitable).(AliasedExpression)
	if isAliased {
		renderer.currentAliasedElements = pushAliased(renderer.currentAliasedElements, aliasedExpression)
	}
	renderer.currentLevel = renderer.currentLevel + 1
	nextLevel := renderer.currentLevel + 1
	_, isSub := (*visitable).(SubVisitable)
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



func (renderer *CypherRenderer) PostLeave(visitable Visitable)  {
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


func (renderer *CypherRenderer) getAliasIfSeen(visitable *Visitable, address interface{}) *Visitable {
	if (*visitable).GetType() != AliasedExpressionVisitable {
		return visitable
	}
	if expression, isExist := renderer.visitableToAliased [address]; isExist {
		return &expression
	}
	return visitable
}

//Enter And Leave

func enterMatch(renderer *CypherRenderer, visitable Visitable ) {
	match := visitable.(Match)
	if match.isOptional() {
		renderer.builder.WriteString("OPTIONAL ")
	}
	renderer.builder.WriteString("MATCH ")
}

func leaveMatch(renderer *CypherRenderer, visitable Visitable ) {
	renderer.builder.WriteString(" ")
}

func enterWhere(renderer *CypherRenderer, visitable Visitable ) {
	renderer.builder.WriteString(" WHERE ")
}

func leaveWhere(renderer *CypherRenderer, visitable Visitable ) {
	renderer.builder.WriteString(" ")
}

func enterReturn(renderer *CypherRenderer, visitable Visitable ) {
	renderer.builder.WriteString("RETURN  ")
}

func leaveReturn(renderer *CypherRenderer, visitable Visitable ) {
}

func enterNode(renderer *CypherRenderer, visitable Visitable) {
	renderer.builder.WriteString("(")
	node := visitable.(Node)
	if !node.hasSymbolic(){
		return
	}
	named := visitable.(Named)
	_, renderer.skipNodeContent = renderer.visitedNamed[&named]
	if renderer.skipNodeContent {
		renderer.builder.WriteString(node.symbolicName.Value)
	}
}

func leaveNode(renderer *CypherRenderer, visitable Visitable)  {
	renderer.builder.WriteString(")")
	renderer.skipNodeContent = false
}

func enterProperties(renderer *CypherRenderer, visitable Visitable)  {
	renderer.builder.WriteString(" ")
}

func leaveProperties(renderer *CypherRenderer, visitable Visitable) {
}

func enterSymbolic(renderer *CypherRenderer, visitable Visitable ) {
	symbolicName := visitable.(SymbolicName)
	renderer.builder.WriteString(renderer.resolve(symbolicName))
}

func leaveSymbolic(renderer *CypherRenderer, visitable Visitable ) {

}

const REL_TYP_SEPARATOR =  "|"
const REL_TYPE_START = ":"
const NODE_LABEL_START = ":"

func enterRelationshipTypes(renderer *CypherRenderer, visitable Visitable ) {
	relationshipTypes := visitable.(RelationshipTypes)
	typeWithPrefix := make([]string, 0)
	for _, typeRaw := range relationshipTypes.values {
		if typeRaw == ""{
			continue
		}
		typeWithPrefix = append(typeWithPrefix, REL_TYPE_START + typeRaw)
	}
	renderer.builder.WriteString(strings.Join(typeWithPrefix, REL_TYP_SEPARATOR))
}



func enterRelationshipDetail(renderer *CypherRenderer, visitable Visitable){
	details := visitable.(RelationshipDetails)
	direction := details.direction
	renderer.builder.WriteString(direction.symbolLeft)
	if details.hasContent() {
		renderer.builder.WriteString("[")
	}
}

func leaveRelationshipDetail(renderer *CypherRenderer, visitable Visitable){
	details := visitable.(RelationshipDetails)
	direction := details.direction
	if details.hasContent() {
		renderer.builder.WriteString("]")
	}
	renderer.builder.WriteString(direction.symbolRight)
}

func enterSkip(renderer *CypherRenderer, visitable Visitable){
	renderer.builder.WriteString(" SKIP ")
}

func enterLimit(renderer *CypherRenderer, visitable Visitable){
	renderer.builder.WriteString(" LIMIT ")
}

func enterLiteral(renderer *CypherRenderer, visitable Visitable){
	literal := visitable.(Literal)
	renderer.builder.WriteString(literal.AsString())
}



func push(queue[] *Visitable, element *Visitable) []*Visitable {
	queue = append(queue, element)
	return queue
}

func pop(queue[] *Visitable) []*Visitable {
	if len(queue) > 0 {
		return queue[:len(queue)-1]
	}
	return queue
}

func peek(queue[] *Visitable) *Visitable {
	if len(queue) > 0 {
		return queue[len(queue)-1]
	}
	return nil
}


func pushAliased(queue[] AliasedExpression, element AliasedExpression) []AliasedExpression {
	queue = append(queue, element)
	return queue
}

func popAliased(queue[] AliasedExpression) []AliasedExpression {
	if len(queue) > 0 {
		return queue[:len(queue)-1]
	}
	return queue
}

func peekAliased(queue[] AliasedExpression) *AliasedExpression {
	if len(queue) > 0 {
		return &queue[len(queue)-1]
	}
	return nil
}
