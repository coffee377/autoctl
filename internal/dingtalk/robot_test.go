package dingtalk

import (
	"testing"

	"github.com/coffee377/autoctl/pkg/log"
)

const testCardTemplateId = "f49a15e0-0352-40e7-ac99-4205ac78a332.schema"

func TestSendMessage(t *testing.T) {
	app := &App{
		Id:           "118447d2-1c73-486f-8058-7daa046c9577",
		AgentId:      "194334207",
		ClientID:     "dingybihm3fg4sjh3dtx",
		ClientSecret: "smpvcY639CMUdAfmWOoyIImFCdD0woA09cMp7S5AsAQZGki6XFUUVrp0XCUCE-N2",
		RobotCode:    "dingybihm3fg4sjh3dtx",
	}
	robot, err := NewRobot(app)
	if err != nil {
		panic(err)
	}
	message, err := robot.SendCardMessage(SingleChat, testCardTemplateId)
	if err != nil {
		panic(err)
	}
	log.Info("ProcessQueryKey => %s", message)
}
