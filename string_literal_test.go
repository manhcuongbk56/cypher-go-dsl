package cypher

import (
	"fmt"
	"strconv"
	"testing"
)

func TestEscapeString(t *testing.T) {
	origin := "A\\B\\\\Ca'bc123\\"
	escaped := strconv.Quote(origin)
	fmt.Println(escaped)
	literal := StringLiteral{content: origin}
	escaped = literal.AsString()
	// "'A\\\\B\\\\\\\\Ca\\'bc123\\\\'"
	fmt.Println(escaped)
}
