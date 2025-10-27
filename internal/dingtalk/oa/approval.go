package oa

import (
	"cds/dingtalk/app"
	"strconv"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dingtalkworkflow10 "github.com/alibabacloud-go/dingtalk/workflow_1_0"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

const (
	BidApplyProcessCode      = "PROC-958C3100-85BF-45D3-8583-6645DA922756" // 投标申请审批表单编码
	BidExpenseProcessCode    = "PROC-D8453B77-B313-4BEB-BE42-C71EE81DA61A" // 投标项目转款表单编码
	BidExpenseProcessCodeNew = "PROC-4EDF46AF-CACD-47F2-A787-0C1EAC98C935"
)

type Approval struct {
	app.App
	cli *dingtalkworkflow10.Client
}

func New(app app.App) (*Approval, error) {
	config := &openapi.Config{}
	config.SetProtocol("https")
	config.SetRegionId("central")
	client, err := dingtalkworkflow10.NewClient(config)

	if err != nil {
		return nil, err
	}
	return &Approval{
		cli: client,
		App: app,
	}, nil
}

// GetProcessInstanceIds 获取审批实例ID列表,时间格式 yyyy-MM-dd
func (a *Approval) GetProcessInstanceIds(processCode string, startTime string, endTime string, nextToken *int64) ([]string, error) {
	timeRange, err := ProcessTimeRange(startTime, endTime)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, segment := range timeRange {
		ids, err := a.processInstanceIds(processCode, segment, nextToken)
		if err != nil {
			return nil, err
		}
		result = append(result, ids...)
	}
	return result, nil
}

func (a *Approval) GetProcessInstanceIdsByMonth(processCode string, year uint, month uint8, nextToken *int64) ([]string, error) {
	startTime, endTime, err := GetMonthStartAndEnd(year, month)
	if err != nil {
		return nil, err
	}
	return a.GetProcessInstanceIds(processCode, startTime, endTime, nextToken)
}

// GetProcessInstance 获取审批实例详情
func (a *Approval) GetProcessInstance(instanceId string) (*dingtalkworkflow10.GetProcessInstanceResponseBodyResult, error) {
	processInstanceHeaders := &dingtalkworkflow10.GetProcessInstanceHeaders{
		XAcsDingtalkAccessToken: tea.String(a.GetAccessToken()),
	}
	processInstanceRequest := &dingtalkworkflow10.GetProcessInstanceRequest{
		ProcessInstanceId: tea.String(instanceId),
	}

	result, err := a.cli.GetProcessInstanceWithOptions(processInstanceRequest, processInstanceHeaders, &util.RuntimeOptions{})
	if err != nil {
		return nil, err
	}

	return result.Body.Result, nil
}

func (a *Approval) GetAttachmentDownloadUri(instanceId string, fileId string) (*string, error) {
	// 注册信息文件
	headers := &dingtalkworkflow10.GrantProcessInstanceForDownloadFileHeaders{
		XAcsDingtalkAccessToken: tea.String(a.GetAccessToken()),
	}
	request := &dingtalkworkflow10.GrantProcessInstanceForDownloadFileRequest{
		ProcessInstanceId: tea.String(instanceId),
		FileId:            tea.String(fileId),
	}
	fileRes, err := a.cli.GrantProcessInstanceForDownloadFileWithOptions(request, headers, &util.RuntimeOptions{})
	if err != nil {
		return nil, err
	}
	return fileRes.Body.Result.DownloadUri, nil
}

func (a *Approval) processInstanceIds(processCode string, segment TimeSegment, token *int64) ([]string, error) {
	headers := &dingtalkworkflow10.ListProcessInstanceIdsHeaders{
		XAcsDingtalkAccessToken: tea.String(a.GetAccessToken()),
	}

	listProcessInstanceIdsRequest := &dingtalkworkflow10.ListProcessInstanceIdsRequest{
		ProcessCode: tea.String(processCode),
		StartTime:   tea.Int64(segment.Start.UnixMilli()), // 审批实例开始时间，Unix时间戳，单位毫秒
		EndTime:     tea.Int64(segment.End.UnixMilli()),   // 审批实例结束时间，Unix时间戳，单位毫秒
		NextToken:   tea.Int64(0),
		MaxResults:  tea.Int64(20),
		//Statuses:   tea.StringSlice([]string{"COMPLETED"}), // 流程实例状态： RUNNING：审批中 TERMINATED：已撤销 COMPLETED：审批完成
	}

	if token != nil {
		listProcessInstanceIdsRequest.NextToken = token
	}

	res, err := a.cli.ListProcessInstanceIdsWithOptions(listProcessInstanceIdsRequest, headers, &util.RuntimeOptions{})
	if err != nil {
		return nil, err
	}

	bodyResult := res.Body.Result
	var result []string

	for _, v := range bodyResult.List {
		if v != nil {
			result = append(result, *v)
		}
	}

	// 如果有更多数据继续获取
	if bodyResult.NextToken != nil {
		nextToken, _ := strconv.ParseInt(*bodyResult.NextToken, 10, 64)
		ids, e := a.processInstanceIds(processCode, segment, &nextToken)
		if e != nil {
			return nil, err
		}
		result = append(result, ids...)
	}

	return result, nil
}
