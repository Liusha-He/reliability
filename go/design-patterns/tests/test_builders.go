package tests

import (
	"strings"
	"testing"

	builder "design-patterns/builders"
	assert "github.com/stretchr/testify/assert"
)

func TestStringBuilder(t *testing.T) {
	hello := "hello"

	builder := strings.Builder{}

	builder.WriteString("<p>")
	builder.WriteString(hello)
	builder.WriteString("</p>")

	assert.Equal(t, builder.String(), "<p>hello</p>",
		"string builder output incorrect.")
}

func TestFacets(t *testing.T) {
	// do something here...
}

func TestParameters(t *testing.T) {
	// do something here.
}

func TestFunctional(t *testing.T) {
	// do something here...
}
