package runss

import (
	"fmt"
	"gopkg.in/readline.v1"
	"io"
)

func (cmd *Cmd) Prompt(writer io.Writer) (err error) {
	rl, err := readline.New("> ")

	if err != nil {
		return
	}

	defer rl.Close()

	for {
		line, rlerr := rl.Readline()

		if rlerr != nil {
			break
		}

		cmd.Command = line
		cmderr := cmd.Run()

		if cmderr != nil {
			fmt.Println(cmderr)
		} else {
			cmd.PrintResults(writer)
		}
	}

	return
}
