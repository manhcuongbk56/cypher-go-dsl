package cypher

type Statement interface {
	Visitable
}

type RegularQuery interface {
	Statement
}

type SingleQuery interface {
	Statement
}

func StatementCall(namespaceAndProcedure ...string) OngoingStandaloneCallWithoutArguments {
	return StandaloneCallBuilderCreate(ProcedureNameCreate(namespaceAndProcedure...))
}
