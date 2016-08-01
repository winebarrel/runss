package main

import (
	"fmt"
	"log"
	"regexp"
	"runss"
	"strings"
)

func init() {
	log.SetFlags(0)
}

func printOutput(output string) {
	fmt.Println("  Output: |")
	output = strings.TrimRight(output, "\n")
	r := regexp.MustCompile(`(?m)^`)
	output = r.ReplaceAllString(output, "    ")
	fmt.Println(output)
}

func printResult(instanceId string, result *runss.Result) {
	fmt.Printf("- InstanceId: %s\n", instanceId)
	fmt.Printf("  Status: %s\n", result.Status)
	printOutput(result.Output)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	cmd := &runss.Cmd{
		InstanceIds: []string{},
		Results:     map[string]*runss.Result{},
	}

	if err := runss.ParseFlag(cmd); err != nil {
		panic(err)
	}

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	for instanceId, result := range cmd.Results {
		printResult(instanceId, result)
	}
}
