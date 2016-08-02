package main

import (
	"log"
	"os"
	"runss"
)

func init() {
	log.SetFlags(0)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	cmd := runss.NewCmd()

	prompt, err := runss.ParseFlag(cmd)

	if err != nil {
		panic(err)
	}

	if prompt {
		err := cmd.Prompt(os.Stdout)

		if err != nil {
			panic(err)
		}
	} else {
		err = cmd.Run()

		if err != nil {
			panic(err)
		}

		cmd.PrintResults(os.Stdout)
	}
}
