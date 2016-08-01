package runss

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mockaws"
	"testing"
)

func TestSendCommand(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cmd := &Cmd{
		InstanceIds: []string{"i-abcdef"},
		Command:     "hostname",
	}

	mockssm := mockaws.NewMockSSMAPI(ctrl)

	mockssm.EXPECT().SendCommand(&ssm.SendCommandInput{
		DocumentName: aws.String(RunShellScriptDocumentName),
		InstanceIds:  []*string{aws.String("i-abcdef")},
		Parameters:   map[string][]*string{"commands": []*string{aws.String("hostname")}},
	}).Return(
		&ssm.SendCommandOutput{
			Command: &ssm.Command{
				CommandId: aws.String("<command id>"),
			},
		},
		nil,
	)

	cmd.sendCommand(mockssm)

	assert.Equal(aws.String("<command id>"), cmd.CommandId)
}

func TestListCommands(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cmd := &Cmd{
		InstanceIds: []string{"i-abcdef"},
		Command:     "hostname",
		CommandId:   aws.String("<command id>"),
	}

	mockssm := mockaws.NewMockSSMAPI(ctrl)

	mockssm.EXPECT().ListCommands(&ssm.ListCommandsInput{
		CommandId: aws.String("<command id>"),
	}).Return(
		&ssm.ListCommandsOutput{
			Commands: []*ssm.Command{&ssm.Command{
				Status: aws.String("Success"),
			}},
		},
		nil,
	)

	status, _ := cmd.listCommands(mockssm)

	assert.Equal("Success", status)
}
