package runss

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"time"
)

const (
	RunShellScriptDocumentName = "AWS-RunShellScript"
)

type Result struct {
	Output string
	Status string
}

type Cmd struct {
	Command     string
	InstanceIds []string
	CommandId   *string
	Results     map[string]*Result
}

func (cmd *Cmd) sendCommand(svc *ssm.SSM) (err error) {
	instanceIds := []*string{}

	for _, id := range cmd.InstanceIds {
		instanceIds = append(instanceIds, aws.String(id))
	}

	commands := []*string{aws.String(cmd.Command)}

	params := &ssm.SendCommandInput{
		DocumentName: aws.String(RunShellScriptDocumentName),
		InstanceIds:  instanceIds,
		Parameters:   map[string][]*string{"commands": commands},
	}

	resp, err := svc.SendCommand(params)

	if err != nil {
		return
	}

	cmd.CommandId = resp.Command.CommandId

	return
}

func (cmd *Cmd) listCommands(svc *ssm.SSM) (status string, err error) {
	params := &ssm.ListCommandsInput{
		CommandId: cmd.CommandId,
	}

	resp, err := svc.ListCommands(params)

	if err != nil {
		return
	}

	if len(resp.Commands) < 1 {
		err = fmt.Errorf("Command not found: %s", *cmd.CommandId)
		return
	}

	status = *resp.Commands[0].Status

	return
}

func (cmd *Cmd) listCommandInvocations(svc *ssm.SSM) (err error) {
	params := &ssm.ListCommandInvocationsInput{
		CommandId: cmd.CommandId,
		Details:   aws.Bool(true),
	}

	outputs := []*ssm.CommandInvocation{}

	err = svc.ListCommandInvocationsPages(params, func(page *ssm.ListCommandInvocationsOutput, lastPage bool) bool {
		for _, ci := range page.CommandInvocations {
			outputs = append(outputs, ci)
		}

		return !lastPage
	})

	if err != nil {
		return
	}

	if len(outputs) < 1 {
		err = fmt.Errorf("CommandInvocation not found: %s", *cmd.CommandId)
		return
	}

	for _, output := range outputs {
		commandPlugin := output.CommandPlugins[0]

		cmd.Results[*output.InstanceId] = &Result{
			Output: *commandPlugin.Output,
			Status: *commandPlugin.Status,
		}
	}

	return
}

func (cmd *Cmd) waitCommand(svc *ssm.SSM) (err error) {
	status := ""

	for {
		status, err = cmd.listCommands(svc)

		if err != nil {
			return
		}

		if status != "Pending" && status != "InProgress" {
			break
		}

		time.Sleep(1 * time.Second)
	}

	if status != "Success" && status != "Failed" {
		err = fmt.Errorf("Ccommand faile: %s", status)
		return
	}

	return
}

func (cmd *Cmd) Run() (err error) {
	svc := ssm.New(session.New())

	err = cmd.sendCommand(svc)

	if err != nil {
		return
	}

	err = cmd.waitCommand(svc)

	if err != nil {
		return
	}

	err = cmd.listCommandInvocations(svc)

	if err != nil {
		return
	}

	return
}
