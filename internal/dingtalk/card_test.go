package dingtalk

import (
	"testing"

	"github.com/coffee377/autoctl/internal/dingtalk/app"
	"github.com/coffee377/autoctl/pkg/log"
)

var (
	a = app.New("118447d2-1c73-486f-8058-7daa046c9577", app.WithRedis())
)

func TestCard_Create(t *testing.T) {
	card, err := NewCard(a)
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
	card, err := NewCard(a)
	if err != nil {
		panic(err)
	}
	result, err := card.CreateAndDeliver(testCardTemplateId)
	if err != nil {
		panic(err)
	}
	log.Info("result => %v", result)
}
