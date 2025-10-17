package oa

import (
	dingtalkworkflow10 "github.com/alibabacloud-go/dingtalk/workflow_1_0"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func (a *Approval) GetFormSchema(processCode string) (*dingtalkworkflow10.QuerySchemaByProcessCodeResponseBodyResult, error) {
	headers := &dingtalkworkflow10.QuerySchemaByProcessCodeHeaders{
		XAcsDingtalkAccessToken: tea.String(a.GetAccessToken()),
	}
	request := &dingtalkworkflow10.QuerySchemaByProcessCodeRequest{
		ProcessCode: tea.String(processCode),
	}
	res, err := a.cli.QuerySchemaByProcessCodeWithOptions(request, headers, &util.RuntimeOptions{})
	if err != nil {
		return nil, err
	}
	return res.Body.Result, nil
}
