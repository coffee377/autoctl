package test

import (
	"cds/bid/data"
	"cds/bid/ent"
	"cds/bid/ent/bidapply"
	"cds/bid/ent/bidproject"
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
	//ids, err := approval.GetProcessInstanceIds(oa.BidApplyProcessCode, "2025-01-01", "", nil)
	//assert.Nil(t, err)

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

	//for i, id := range ids {
	id := "Rp9D_t0WQrqgpfxvxUZ_EQ07201760318742"
	res, err := approval.GetProcessInstance(id)
	assert.Nil(t, err)

	applyData, err := data.NewBidApply(id, res, data.WithUserHook(approval.GetUserHook()))
	assert.Nil(t, err)

	err = WithTx(ctx, client, func(tx *ent.Tx) error {
		// 项目信息
		project, err1 := saveProject(ctx, tx, applyData)
		if err1 != nil {
			return err1
		}
		// 申请信息
		apply, err2 := saveApply(ctx, tx, applyData, project)
		if err2 != nil {
			return err2
		}
		t.Log(apply)
		return nil
	})

	assert.Nil(t, err)
	//break
	//}
}

func saveApply(ctx context.Context, tx *ent.Tx, applyData *data.BidApplyForm, project *ent.BidProject) (*ent.BidApply, error) {
	count, err := tx.BidApply.Query().Where(bidapply.ID(applyData.ID)).Count(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		create := tx.BidApply.Create()
		create.SetID(applyData.ID)
		create.SetBusinessID(applyData.BusinessId)
		create.SetInstanceID(applyData.InstanceId)
		create.SetProjectID(project.ID)

		create.SetNillableOpeningDate(applyData.OpeningDate)
		create.SetNillableNoticeURL(applyData.NoticeUrl)
		create.SetBudgetAmount(applyData.BudgetAmount)
		create.SetNillableRemark(applyData.Remark)

		create.SetApprovalStatus(applyData.ApprovalStatus)

		apply, err := create.Save(ctx)
		if err != nil {
			return nil, err
		}
		return apply, nil
	}

	update := tx.BidApply.UpdateOneID(applyData.ID)
	update.SetBusinessID(applyData.BusinessId)
	update.SetInstanceID(applyData.InstanceId)
	update.SetProjectID(project.ID)

	update.SetNillableOpeningDate(applyData.OpeningDate)
	update.SetNillableNoticeURL(applyData.NoticeUrl)
	update.SetBudgetAmount(applyData.BudgetAmount)
	update.SetNillableRemark(applyData.Remark)

	update.SetApprovalStatus(applyData.ApprovalStatus)

	apply, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return apply, nil
}

func saveProject(ctx context.Context, tx *ent.Tx, applyData *data.BidApplyForm) (*ent.BidProject, error) {
	project, err := tx.BidProject.Query().Where(bidproject.IDEQ(applyData.ProjectID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	if project != nil {
		return project, nil
	}

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

	projectCreate.SetBizRepNo(applyData.CreateBy)
	projectCreate.SetBizRepName(applyData.CreatorName)

	project, err = projectCreate.Save(ctx)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
		}
	}()

	if err = fn(tx); err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rErr)
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
