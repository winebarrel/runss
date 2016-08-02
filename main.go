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

	if err := runss.ParseFlag(cmd); err != nil {
		panic(err)
	}

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	cmd.PrintResults(os.Stdout)
}
