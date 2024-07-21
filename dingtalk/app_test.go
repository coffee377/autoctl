package dingtalk

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dingtalkworkflow10 "github.com/alibabacloud-go/dingtalk/workflow_1_0"
	"github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	app := App{}
	app.AgentId = "1038540627"
	app.AppKey = "dingopfniakkw72klkjv"
	app.AppSecret = "6Il0DuPZPPIr-OG03uMrnqDNu_o03tpIkK03ScpuEPP6NAw7J52D0LWPvTjRf4BR"
	accessToken := app.GetAccessToken()
	t.Log(accessToken)

	// 获取审批实例ID列表
	config := &openapi.Config{}
	config.SetProtocol("https")
	config.SetRegionId("central")
	client, _ := dingtalkworkflow10.NewClient(config)

	headers := &dingtalkworkflow10.ListProcessInstanceIdsHeaders{
		XAcsDingtalkAccessToken: &accessToken,
	}
	listProcessInstanceIdsRequest := &dingtalkworkflow10.ListProcessInstanceIdsRequest{
		ProcessCode: tea.String("PROC-CFC2784B-CD66-43C3-91A5-B3CD7D3ABBC6"),
		StartTime:   tea.Int64(1688140800000),
		//EndTime:     tea.Int64(time.Now().UnixMilli()),
		NextToken:  tea.Int64(0),
		MaxResults: tea.Int64(20),
		Statuses:   tea.StringSlice([]string{"COMPLETED"}),
	}
	result, err := client.ListProcessInstanceIdsWithOptions(listProcessInstanceIdsRequest, headers, &service.RuntimeOptions{})
	if err != nil {
		t.Log(result)
	}
}

/**
 * 使用 Token 初始化账号Client
 * @return Client
 * @throws Exception
 */
func CreateClient() (_result *dingtalkworkflow10.Client, _err error) {
	config := &openapi.Config{}
	config.SetProtocol("https")
	config.SetRegionId("central")
	//_result = &dingtalkworkflow_1_0.Client{}
	_result, _err = dingtalkworkflow10.NewClient(config)
	return _result, _err
}

func _main(args []*string) (_err error) {
	client, _err := CreateClient()
	if _err != nil {
		return _err
	}

	listProcessInstanceIdsHeaders := &dingtalkworkflow10.ListProcessInstanceIdsHeaders{}
	listProcessInstanceIdsHeaders.XAcsDingtalkAccessToken = tea.String("<your access token>")
	listProcessInstanceIdsRequest := &dingtalkworkflow10.ListProcessInstanceIdsRequest{
		StartTime:   tea.Int64(1680278400000),
		ProcessCode: tea.String("PROC-CFC2784B-CD66-43C3-91A5-B3CD7D3ABBC6"),
		NextToken:   tea.Int64(0),
		MaxResults:  tea.Int64(20),
	}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		_, _err = client.ListProcessInstanceIdsWithOptions(listProcessInstanceIdsRequest, listProcessInstanceIdsHeaders, &service.RuntimeOptions{})
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var err = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			err = _t
		} else {
			err.Message = tea.String(tryErr.Error())
		}
		if !tea.BoolValue(service.Empty(err.Code)) && !tea.BoolValue(service.Empty(err.Message)) {
			// err 中含有 code 和 message 属性，可帮助开发定位问题
		}

	}
	return _err
}
