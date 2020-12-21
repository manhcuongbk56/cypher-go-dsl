package cypher

type Predicate struct {
	implementationName string
}

func (p Predicate) getImplementationName() string {
	return p.implementationName
}

func (p Predicate) isAggregate() bool {
	return false
}

var ALL = Predicate{implementationName: "all"}
var ANY = Predicate{implementationName: "any"}
var EXISTS = Predicate{implementationName: "exists"}
var NONE = Predicate{implementationName: "none"}
var SINGLE = Predicate{implementationName: "single"}

type Scalars struct {
	implementationName string
}

func (s Scalars) getImplementationName() string {
	return s.implementationName
}

func (s Scalars) isAggregate() bool {
	return false
}

var COALESCE = Scalars{implementationName: "coalesce"}
var END_NODE = Scalars{implementationName: "endNode"}
var HEAD = Scalars{implementationName: "head"}
var ID = Scalars{implementationName: "id"}
var LAST = Scalars{implementationName: "last"}
var PROPERTIES = Scalars{implementationName: "properties"}
var SHORTEST_PATH = Scalars{implementationName: "shortestPath"}
var SIZE = Scalars{implementationName: "FunctionSize"}
var START_NODE = Scalars{implementationName: "startNode"}
var TYPE = Scalars{implementationName: "type"}

type Strings struct {
	implementationName string
}

func (s Strings) getImplementationName() string {
	return s.implementationName
}

func (s Strings) isAggregate() bool {
	return false
}

var TO_LOWER = Strings{implementationName: "toLower"}

type Spatials struct {
	implementationName string
}

func (s Spatials) getImplementationName() string {
	return s.implementationName
}

func (s Spatials) isAggregate() bool {
	return false
}

var POINT = Spatials{implementationName: "point"}
var DISTANCE = Spatials{implementationName: "distance"}

type Aggregates struct {
	implementationName string
}

func (a Aggregates) getImplementationName() string {
	return a.implementationName
}

func (a Aggregates) isAggregate() bool {
	return true
}

var AVG = Aggregates{"avg"}
var COLLECT = Aggregates{"collect"}
var COUNT = Aggregates{"FunctionCount"}
var MAX = Aggregates{"max"}
var MIN = Aggregates{"min"}
var PERCENTILE_CONT = Aggregates{"percentileCont"}
var PERCENTILE_DISC = Aggregates{"percentileDisc"}
var ST_DEV = Aggregates{"stDev"}
var ST_DEV_P = Aggregates{"stDevP"}
var SUM = Aggregates{"sum"}

type Lists struct {
	implementationName string
}

func (l Lists) getImplementationName() string {
	return l.implementationName
}

func (l Lists) isAggregate() bool {
	return false
}

var LABELS = Lists{implementationName: "labels"}
var NODES = Lists{implementationName: "nodes"}
var RANGE = Lists{implementationName: "range"}
var RELATIONSHIPS = Lists{implementationName: "relationships"}

type Temporals struct {
	implementationName string
}

func (t Temporals) getImplementationName() string {
	return t.implementationName
}

func (t Temporals) isAggregate() bool {
	return false
}

var DATE = Temporals{implementationName: "date"}
var DATETIME = Temporals{implementationName: "datetime"}
var LOCALDATETIME = Temporals{implementationName: "localdatetime"}
var LOCALTIME = Temporals{implementationName: "localtime"}
var TIME = Temporals{implementationName: "time"}
var DURATION = Temporals{implementationName: "duration"}
