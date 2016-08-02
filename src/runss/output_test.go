package runss

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintOutput(t *testing.T) {
	assert := assert.New(t)

	buf := &bytes.Buffer{}

	printOutput(buf, "abc\n&gt;&lt;\ndef\n")

	expected := "  Output: |\n    abc\n    ><\n    def\n    \n"

	assert.Equal(expected, buf.String())
}

func TestPrintResult(t *testing.T) {
	assert := assert.New(t)

	buf := &bytes.Buffer{}
	result := &Result{Status: "Success"}

	printResult(buf, "i-abcdef", result)

	expected := "- InstanceId: i-abcdef\n  Status: Success\n  Output: \n"

	assert.Equal(expected, buf.String())
}

func TestPrintResults(t *testing.T) {
	assert := assert.New(t)

	buf := &bytes.Buffer{}

	cmd := &Cmd{
		Results: map[string]*Result{
			"i-abcdef": &Result{Status: "Success"},
		},
	}

	cmd.PrintResults(buf)

	expected := "- InstanceId: i-abcdef\n  Status: Success\n  Output: \n"

	assert.Equal(expected, buf.String())
}
