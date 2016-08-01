package runss

import (
	"fmt"
	"html"
	"io"
	"regexp"
)

func printOutput(writer io.Writer, output string) {
	fmt.Fprintln(writer, "  Output: |")
	output = html.UnescapeString(output)
	r := regexp.MustCompile(`(?m)^`)
	output = r.ReplaceAllString(output, "    ")
	fmt.Fprintln(writer, output)
}

func printResult(writer io.Writer, instanceId string, result *Result) {
	fmt.Fprintf(writer, "- InstanceId: %s\n", instanceId)
	fmt.Fprintf(writer, "  Status: %s\n", result.Status)
	printOutput(writer, result.Output)
}

func (cmd *Cmd) PrintResults(writer io.Writer) {
	for instanceId, result := range cmd.Results {
		printResult(writer, instanceId, result)
	}
}
