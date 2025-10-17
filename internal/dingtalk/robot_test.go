package dingtalk

import (
	"testing"

	"github.com/coffee377/autoctl/internal/dingtalk/app"
	"github.com/coffee377/autoctl/pkg/log"
	"github.com/stretchr/testify/assert"
)

const testCardTemplateId = "f49a15e0-0352-40e7-ac99-4205ac78a332.schema"

func TestSendMessage(t *testing.T) {
	robot, err := NewRobot(app.New("118447d2-1c73-486f-8058-7daa046c9577", app.WithRedis()))
	if err != nil {
		panic(err)
	}
	message, err := robot.SendCardMessage(SingleChat, testCardTemplateId)
	assert.Nil(t, err)
	log.Info("ProcessQueryKey => %s", message)
}
