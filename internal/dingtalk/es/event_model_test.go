package es

const (
	header   = ``
	instance = `{
  "processInstanceId" : "aIivQGsaToGCyZ_L4b9DEA07201766047418",
  "eventId" : "e0cf09bc7e194ebdae2188485cfd8a08",
  "resource" : "/v1.0/event/bpms_instance_change/processCode/PROC-BDBD627A-7E7F-4EBC-B80A-EC2A9E777D97/type/start",
  "createTime" : 1766047418000,
  "processCode" : "PROC-BDBD627A-7E7F-4EBC-B80A-EC2A9E777D97",
  "bizCategoryId" : "",
  "businessId" : "202512181643000385959",
  "title" : "吴玉杰提交的数据拉取测试申请",
  "type" : "start",
  "businessType" : "",
  "url" : "https://aflow.dingtalk.com/dingtalk/mobile/homepage.htm?corpid=dingd8b32bfb2b9da7b2&dd_share=false&showmenu=false&dd_progress=false&back=native&procInstId=aIivQGsaToGCyZ_L4b9DEA07201766047418&taskId=&swfrom=isv&dinghash=approval&dtaction=os&dd_from=#approval",
  "staffId" : "021404083621658683"
}`
	task = `{
  "processInstanceId" : "aIivQGsaToGCyZ_L4b9DEA07201766047418",
  "eventId" : "7a320e691954402baefa490a769c42a7",
  "finishTime" : 1766047844000,
  "resource" : "/v1.0/event/bpms_task_change/processCode/PROC-BDBD627A-7E7F-4EBC-B80A-EC2A9E777D97/type/finish",
  "businessId" : "202512181643000385959",
  "activityName" : "审核人",
  "actualActionerId" : "021404083621658683",
  "agentActionerIds" : [ ],
  "title" : "吴玉杰提交的数据拉取测试申请",
  "type" : "finish",
  "result" : "refuse",
  "activityId" : "1918_5cd3",
  "createTime" : 1766047419000,
  "processCode" : "PROC-BDBD627A-7E7F-4EBC-B80A-EC2A9E777D97",
  "bizCategoryId" : "",
  "businessType" : "",
  "staffId" : "021404083621658683",
  "taskId" : 98238556029
}`
)

//func TestProcessInstanceUnmarshalJSON(t *testing.T) {
//	p, err := CreateProcessInstanceFromJson([]byte(instance))
//	assert.Nil(t, err)
//	assert.NotNil(t, p)
//	fmt.Printf("%s,%s", p.CreateTime.String(), p.FinishTime.String())
//	t.Log(p)
//}
//
//func TestProcessTaskUnmarshalJSON(t *testing.T) {
//	ts, err := CreateProcessTaskFromJson([]byte(task))
//	assert.Nil(t, err)
//	assert.NotNil(t, ts)
//	t.Log(t)
//}
