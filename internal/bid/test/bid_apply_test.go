package test

import (
	"cds/bid/data"
	"cds/bid/ds"
	"cds/dingtalk/app"
	"cds/dingtalk/oa"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	kvs = []*data.KV{
		{Code: "00100", Name: "综合管理部"},
		{Code: "00200", Name: "项目实施中心"},
		{Code: "00300", Name: "研发中心"},
		{Code: "00301", Name: "卫生产品部"},
		{Code: "00302", Name: "社会管理产品部"},
		{Code: "00303", Name: "社会救助产品部"},
		{Code: "00304", Name: "技术支持部"},
		{Code: "00305", Name: "医疗保险产品部"},
		{Code: "00400", Name: "财务部"},
		{Code: "00500", Name: "营销中心"},
		{Code: "00501", Name: "市场一部"},
		{Code: "00502", Name: "市场二部"},
		{Code: "00503", Name: "市场三部"},
		{Code: "00600", Name: "行政部"},
		{Code: "00700", Name: "商务部"},
		{Code: "00900", Name: "售后服务中心"},
		{Code: "01000", Name: "贵州分公司"},
		{Code: "01100", Name: "人口健康平台事业部"},
		{Code: "01200", Name: "基层卫生事业部"},
		{Code: "01300", Name: "民政综合平台事业部"},
		{Code: "01400", Name: "社会救助事业部"},
		{Code: "01500", Name: "技术研究中心"},
		{Code: "01600", Name: "培训中心"},
		{Code: "01700", Name: "系统集成事业部"},
		{Code: "01800", Name: "智慧健康养老事业部"},
		{Code: "01900", Name: "人力资源部"},
		{Code: "02100", Name: "民政事业部"},
		{Code: "02200", Name: "卫生事业部"},
		{Code: "02300", Name: "医疗保障事业部"},
		{Code: "02400", Name: "智慧医疗子公司"},
		{Code: "02500", Name: "总裁办"},
		{Code: "02600", Name: "证券部"},
		{Code: "02700", Name: "知识产权部"},
		{Code: "02800", Name: "贵州梵晶公司"},
		{Code: "02900", Name: "成都分公司"},
		{Code: "03000", Name: "亨源合义"},
		{Code: "03100", Name: "北京晶奇和一子公司"},
		{Code: "03200", Name: "成都晶高孙公司"},
		{Code: "03300", Name: "晶讯健康子公司"},
		{Code: "03400", Name: "哈尔滨分公司"},
		{Code: "03500", Name: "青海分公司"},
		{Code: "03600", Name: "创新管理委员会"},
		{Code: "03601", Name: "委员会办公室"},
		{Code: "03602", Name: "创新业务发展部"},
		{Code: "03603", Name: "创新项目组"},
		{Code: "03700", Name: "西藏晶喜子公司"},
		{Code: "03800", Name: "四川晶奇泽盛子公司"},
		{Code: "03900", Name: "智慧医疗事业部"},
		{Code: "04000", Name: "智慧医院事业部"},
		{Code: "04100", Name: "数字政企事业部"},
		{Code: "04200", Name: "数字能源事业部"},
		{Code: "04300", Name: "安全及技术保障部"},
	}
	dHook data.DictHook = func(key string) (*data.KV, bool) {
		for i, kv := range kvs {
			if kv.Code == key || kv.Name == key {
				return kvs[i], true
			}
		}
		return nil, false
	}
)

func TestApplyBatch(t *testing.T) {
	approval, err := oa.New(app.New("a57e9681-79cb-4242-96df-952be2dc3af7", app.WithRedis()))
	assert.Nil(t, err)
	ids, err := approval.GetProcessInstanceIds(oa.BidApplyProcessCode, "2025-01-01", "2025-10-31", nil)
	assert.Nil(t, err)

	client, ok := ds.Mysql()
	defer ds.CloseMysql(client)
	assert.Equal(t, true, ok)

	ctx := context.Background()

	for i, id := range ids {
		t.Logf("%d: %s", i+1, id)
		res, err := approval.GetProcessInstance(id)
		assert.Nil(t, err)

		applyData, err := data.NewBidApply(id, res, data.WithUserHook(approval.GetUserHook()), data.WithDepartmentHook(dHook))
		assert.Nil(t, err)

		err = applyData.Save(ctx, client, false)
		assert.Nil(t, err)
	}
}

func TestApply(t *testing.T) {
	approval, err := oa.New(app.New("a57e9681-79cb-4242-96df-952be2dc3af7", app.WithRedis()))
	assert.Nil(t, err)

	client, ok := ds.Mysql()
	defer ds.CloseMysql(client)
	assert.Equal(t, true, ok)

	ctx := context.Background()
	id := "WuaVoI30TnmDXnEKdYLdNQ07201741157155"
	res, err := approval.GetProcessInstance(id)
	assert.Nil(t, err)

	url := fmt.Sprintf("https://applink.dingtalk.com/approval/detail?corpId=%s&instanceId=%s&from=%s",
		"dingd8b32bfb2b9da7b2", id, "dingopfniakkw72klkjv")
	t.Log(url)

	applyData, err := data.NewBidApply(id, res, data.WithUserHook(approval.GetUserHook()), data.WithDepartmentHook(dHook))
	assert.Nil(t, err)

	err = applyData.Save(ctx, client, false)
	assert.Nil(t, err)
}
