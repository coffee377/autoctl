package test

import (
	"cds/bid/data"
	"cds/bid/ent"
	"cds/dingtalk/app"
	"cds/dingtalk/oa"
	"context"
	"fmt"
	"log"
	"testing"

	"entgo.io/ent/dialect"
	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	approval, err := oa.New(app.New("a57e9681-79cb-4242-96df-952be2dc3af7", app.WithRedis()))
	assert.Nil(t, err)
	ids, err := approval.GetProcessInstanceIds(oa.BidApplyProcessCode, "2025-01-01", "", nil)
	assert.Nil(t, err)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", "root", "root!@@&", "localhost", "3306", "cds_infra",
		"charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai")
	client, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalf("failed closing connection to mysql: %v", err)
		}
	}(client)

	ctx := context.Background()

	for i, id := range ids {
		t.Log(i, id)
		res, err := approval.GetProcessInstance(id)
		assert.Nil(t, err)

		applyData, err := data.GetApplyData(id, res)
		assert.Nil(t, err)
		t.Log(applyData)

		tx, err := client.Tx(ctx)
		assert.Nil(t, err)

		// 项目信息
		projectCreate := tx.BidProject.Create()
		projectCreate.SetID(applyData.ProjectID)

		if applyData.ProjectCode != nil {
			projectCreate.SetCode(*applyData.ProjectCode)
		}
		projectCreate.SetName(applyData.ProjectName)

		if applyData.DepartmentCode != nil {
			projectCreate.SetDepartmentCode(*applyData.DepartmentCode)
		}
		projectCreate.SetDepartmentName(applyData.DepartmentName)

		//projectCreate.SetBizRepNo("")
		//projectCreate.SetBizRepName("")

		project, err := projectCreate.Save(ctx)
		assert.Nil(t, err)
		t.Log(project)

		// 申请信息
		applyCreate := tx.BidApply.Create()
		applyCreate.SetID(applyData.ID)
		applyCreate.SetBusinessID(applyData.ApprovalNumber)
		applyCreate.SetInstanceID(applyData.InstanceId)
		applyCreate.SetProjectID(project.ID)

		applyCreate.SetNillableOpeningDate(applyData.OpeningDate)
		applyCreate.SetNillableNoticeURL(applyData.NoticeUrl)
		applyCreate.SetBudgetAmount(applyData.BudgetAmount)
		applyCreate.SetNillableRemark(applyData.Remark)

		applyCreate.SetApprovalStatus(applyData.Handler) // todo 审批状态

		apply, err := applyCreate.Save(ctx)
		assert.Nil(t, err)
		t.Log(apply)

		err = tx.Commit()
		assert.Nil(t, err)
		break
	}
}
