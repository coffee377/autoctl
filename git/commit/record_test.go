package git

import (
	"testing"
)

var rawCommit = `e564a3396849178d68e12f938732372ba6254fde;e564a33;WuYujie;coffee377@dingtalk.com;1653057941;2022/05/20 22:45:41;:fix:(log)! prevent racing of requests

Introduce a request id and a reference to the latest request. Dismiss
incoming responses other than from the latest request.

Remove timeouts which were used to mitigate the racing issue but are
obsolete now.

Reviewed-by: Z
Refs: #123
BREAKING CHANGE: use JavaScript features not available in Node 6.`

func TestNewCommitRecord(t *testing.T) {

	t.Run("解析原始提交信息", func(t *testing.T) {
		//message := NewCommitRecord(rawCommit)
		//var data, _ = json.Marshal(message)
		//
		//fmt.Println("s%", data)

	})
}
