package process

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstanceMessageUnmarshalJSON(t *testing.T) {
	var i = InstanceMessage{}
	err := i.UnmarshalJSON([]byte(`{
  "processInstanceId" : "JOjIG0zDSsWeFoNB_VnXNg07201762843919",
  "eventId" : "b69db327a8f94316abef1aee2d4ae2c6",
  "finishTime" : 1762848996000,
  "resource" : "/v1.0/event/bpms_instance_change/processCode/PROC-BDBD627A-7E7F-4EBC-B80A-EC2A9E777D97/type/finish",
  "businessId" : "202511111451000598927",
  "title" : "吴玉杰提交的数据拉取测试申请",
  "type" : "finish",
  "url" : "https://aflow.dingtalk.com/dingtalk/mobile/homepage.htm?corpid=dingd8b32bfb2b9da7b2&dd_share=false&showmenu=false&dd_progress=false&back=native&procInstId=JOjIG0zDSsWeFoNB_VnXNg07201762843919&taskId=&swfrom=isv&dinghash=approval&dtaction=os&dd_from=#approval",
  "result" : "refuse",
  "createTime" : 1762843919000,
  "processCode" : "PROC-BDBD627A-7E7F-4EBC-B80A-EC2A9E777D97",
  "bizCategoryId" : "",
  "businessType" : "",
  "staffId" : "021404083621658683"
}`))
	assert.Nil(t, err)
	t.Log(i)
}
