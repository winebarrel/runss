package runss

import (
	"flag"
	"fmt"
	"strings"
)

func ParseFlag(cmd *Cmd) (prompt bool, err error) {
	idList := ""

	flag.StringVar(&cmd.Command, "command", "", "shell script command")
	flag.StringVar(&idList, "instance-ids", "", "comma separated instance ids")
	flag.BoolVar(&prompt, "prompt", false, "show prompt")
	flag.Parse()

	instanceIds := strings.Split(idList, ",")

	for _, id := range instanceIds {
		id = strings.TrimSpace(id)

		if len(id) > 0 {
			cmd.InstanceIds = append(cmd.InstanceIds, id)
		}
	}

	if !prompt && cmd.Command == "" {
		err = fmt.Errorf("'command' is required")
		return
	}

	if len(cmd.InstanceIds) < 1 {
		err = fmt.Errorf("'instance-ids' is required")
		return
	}

	return
}
