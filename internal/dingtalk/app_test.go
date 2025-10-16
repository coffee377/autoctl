package dingtalk

import (
	"testing"

	"github.com/coffee377/autoctl/internal/dingtalk/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAccessToken(t *testing.T) {
	app1 := App{
		Id:           "118447d2-1c73-486f-8058-7daa046c9577",
		Name:         utils.ToPtr("代码工匠实验室-监控平台"),
		AgentId:      "194334207",
		ClientID:     clientId,
		ClientSecret: clientSecret,
	}
	accessToken := app1.GetAccessToken()
	assert.NotNil(t, accessToken)

	app2 := App{
		Id:           "a57e9681-79cb-4242-96df-952be2dc3af7",
		Name:         utils.ToPtr("安徽晶奇-统一认证"),
		AgentId:      "1038540627",
		ClientID:     "dingopfniakkw72klkjv",
		ClientSecret: "6Il0DuPZPPIr-OG03uMrnqDNu_o03tpIkK03ScpuEPP6NAw7J52D0LWPvTjRf4BR",
	}

	assert.NotNil(t, app2.GetAccessToken())
}
