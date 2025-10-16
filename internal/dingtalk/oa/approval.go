package oa

import (
	"strconv"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dingtalkworkflow10 "github.com/alibabacloud-go/dingtalk/workflow_1_0"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/coffee377/autoctl/internal/dingtalk"
)

const (
	BidApplyProcessCode   = "PROC-958C3100-85BF-45D3-8583-6645DA922756" // 投标申请审批表单编码
	BidExpenseProcessCode = "PROC-D8453B77-B313-4BEB-BE42-C71EE81DA61A" // 投标项目转款表单编码
)

type Approval struct {
	*dingtalk.App
	cli *dingtalkworkflow10.Client
}

func New(app dingtalk.App) (*Approval, error) {
	config := &openapi.Config{}
	config.SetProtocol("https")
	config.SetRegionId("central")
	client, err := dingtalkworkflow10.NewClient(config)

	if err != nil {
		return nil, err
	}
	return &Approval{
		cli: client,
		App: &app,
	}, nil
}

// GetProcessInstanceIds 获取审批实例ID列表
// 时间格式 yyyy-MM-dd
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

// GetProcessInstance 获取审批实例详情
func (a *Approval) GetProcessInstance(instanceId string) (*dingtalkworkflow10.GetProcessInstanceResponseBody, error) {
	accessToken := a.GetAccessToken()
	processInstanceHeaders := &dingtalkworkflow10.GetProcessInstanceHeaders{
		XAcsDingtalkAccessToken: tea.String(accessToken),
	}
	processInstanceRequest := &dingtalkworkflow10.GetProcessInstanceRequest{
		ProcessInstanceId: tea.String(instanceId),
	}

	result, err := a.cli.GetProcessInstanceWithOptions(processInstanceRequest, processInstanceHeaders, &util.RuntimeOptions{})
	if err != nil {
		return nil, err
	}

	return result.Body, nil
}

func (a *Approval) parseTime(timeStr string) (*time.Time, error) {
	start, err := time.ParseInLocation(time.DateOnly, timeStr, time.Local)
	if err != nil {
		return nil, err
	}
	return &start, nil
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
