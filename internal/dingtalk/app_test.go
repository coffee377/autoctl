package dingtalk

import (
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	app := App{
		Id:           "118447d2-1c73-486f-8058-7daa046c9577",
		AgentId:      "194334207",
		ClientKey:    clientId,
		ClientSecret: clientSecret,
	}
	accessToken := app.GetAccessToken()
	t.Log(accessToken)

	// 获取审批实例ID列表
	//config := &openapi.Config{}
	//config.SetProtocol("https")
	//config.SetRegionId("central")
	//client, _ := dingtalkworkflow10.NewClient(config)
	//
	//headers := &dingtalkworkflow10.ListProcessInstanceIdsHeaders{
	//	XAcsDingtalkAccessToken: &accessToken,
	//}
	//listProcessInstanceIdsRequest := &dingtalkworkflow10.ListProcessInstanceIdsRequest{
	//	ProcessCode: tea.String("PROC-CFC2784B-CD66-43C3-91A5-B3CD7D3ABBC6"),
	//	StartTime:   tea.Int64(1688140800000),
	//	//EndTime:     tea.Int64(time.Now().UnixMilli()),
	//	NextToken:  tea.Int64(0),
	//	MaxResults: tea.Int64(20),
	//	Statuses:   tea.StringSlice([]string{"COMPLETED"}),
	//}
	//result, err := client.ListProcessInstanceIdsWithOptions(listProcessInstanceIdsRequest, headers, &service.RuntimeOptions{})
	//if err != nil {
	//	t.Log(result)
	//}
}
