package dingtalk

import (
	"github.com/coffee377/autoctl/pkg/log"
	"testing"
)

func TestCard_Create(t *testing.T) {
	app := &App{
		Id:           "118447d2-1c73-486f-8058-7daa046c9577",
		AgentId:      "194334207",
		ClientKey:    "dingybihm3fg4sjh3dtx",
		ClientSecret: "smpvcY639CMUdAfmWOoyIImFCdD0woA09cMp7S5AsAQZGki6XFUUVrp0XCUCE-N2",
		RobotCode:    "dingybihm3fg4sjh3dtx",
	}
	card, err := NewCard(app)
	if err != nil {
		panic(err)
	}
	outTrackId, err := card.Create("d6f799c5-27d5-484e-a725-c7e176424baf.schema")
	if err != nil {
		panic(err)
	}
	res, err := card.Deliver(outTrackId)
	if err != nil {
		panic(err)
	}
	log.Info("result => %v", res)
}

func TestCard_CreateAndDeliver(t *testing.T) {
	app := &App{
		Id:           "118447d2-1c73-486f-8058-7daa046c9577",
		AgentId:      "194334207",
		ClientKey:    "dingybihm3fg4sjh3dtx",
		ClientSecret: "smpvcY639CMUdAfmWOoyIImFCdD0woA09cMp7S5AsAQZGki6XFUUVrp0XCUCE-N2",
		RobotCode:    "dingybihm3fg4sjh3dtx",
	}
	card, err := NewCard(app)
	if err != nil {
		panic(err)
	}
	result, err := card.CreateAndDeliver("5c9ffe50-2720-4d99-8ace-b2045807f6b5.schema")
	//result, err := card.CreateAndDeliver("d6f799c5-27d5-484e-a725-c7e176424baf.schema")
	if err != nil {
		panic(err)
	}
	log.Info("result => %v", result)
}
