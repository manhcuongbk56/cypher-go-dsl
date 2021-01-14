package cypher

import (
	"testing"
)

func TestShouldEscapeContent(t *testing.T) {
	origin := "A\\B\\\\Ca'bc123\\"
	literal := StringLiteral{content: origin}
	escaped := literal.AsString()
	if escaped != "'A\\\\B\\\\\\\\Ca\\'bc123\\\\'" {
		t.Errorf("escaped is not match:\n %s", escaped)
	}
}


func TestShouldEscapeEmptyString(t *testing.T) {
	inputs := make([][]string, 0)
	inputs = append(inputs, []string{"", ""})
	inputs = append(inputs, []string{" \t ", " \t "})
	inputs = append(inputs, []string{"Nothing to escape", "Nothing to escape"})
	inputs = append(inputs, []string{"' \" '", "\\' \\\" \\'"})
	for _, input := range inputs {
		escaped := escapeStringLiteral(input[0])
		if escaped != input[1] {
			t.Errorf("escaped is not match:\n %s, expect is:\n %s", escaped, input[1])
		}
	}
}
