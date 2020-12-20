package cypher_go_dsl

import (
	"fmt"
	"strconv"
	"testing"
)

func TestEscapeString(t *testing.T) {
	origin := "A\\B\\\\Ca'bc123\\"
	escaped := strconv.Quote(origin)
	fmt.Println(escaped)
}
