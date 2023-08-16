package authority

import "testing"

type Action struct {
	Name        string
	Position    uint
	Description string
}

func (a Action) GetName() string {
	return a.Name
}

func (a Action) GetPosition() uint {
	return a.Position
}

func (a Action) GetDescription() string {
	return a.Description
}

var (
	QueryAction       = Action{"QUERY", 1, "查询"}
	AddAction         = Action{"ADD", 2, "增加"}
	DeleteAction      = Action{"DELETE", 3, "删除"}
	BatchDeleteAction = Action{"BATCH_DELETE", 4, "批量删除"}
	UpdateAction      = Action{"UPDATE", 5, "更新"}
	ImportAction      = Action{"IMPORT", 6, "导入"}
	ExportAction      = Action{"EXPORT", 7, "导出"}
)

func TestWithAuthority(t *testing.T) {
	authority := WithAuthority(ExportAction)

	authority.Add(QueryAction, AddAction, DeleteAction)
	printLog(authority, t)

	authority.Remove(DeleteAction)
	authority.Remove(QueryAction)
	printLog(authority, t)

	//authority.Reset()
	authority.Add(UpdateAction, ImportAction, ExportAction)

	t.Log(authority.Has(DeleteAction))
	t.Log(authority.HasAll(QueryAction, AddAction, DeleteAction, BatchDeleteAction))
	t.Log(authority.HasAny(QueryAction, AddAction, DeleteAction, BatchDeleteAction))
}

func printLog(authority Functions, t *testing.T) {
	t.Logf("source:%s => data:%s\t|\t%s", authority.GetSource().String(), ""+authority.Get().Text(10), authority.GetBaseFunctions(2))
}
