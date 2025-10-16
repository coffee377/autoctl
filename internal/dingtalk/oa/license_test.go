package oa

import (
	"testing"

	"github.com/coffee377/autoctl/internal/dingtalk/app"
)

func Test_Lic(t *testing.T) {
	oa, err := NewOA(app.New("ccl", "a57e9681-79cb-4242-96df-952be2dc3af7",
		app.WithClient("dingopfniakkw72klkjv", "6Il0DuPZPPIr-OG03uMrnqDNu_o03tpIkK03ScpuEPP6NAw7J52D0LWPvTjRf4BR"),
		app.WithRobot("dingopfniakkw72klkjv"),
		app.WithAgent("1038540627"),
	))
	if err != nil {
		panic(err)
	}
	oa.Demo()
}
