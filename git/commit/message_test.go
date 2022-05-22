package git

import (
	"fmt"
	"testing"
)

var rawMsg = []byte(`:fix:(test) prevent racing of requests

Introduce a request id and a reference to latest request. Dismiss
incoming responses other than from latest request.

Remove timeouts which were used to mitigate the racing issue but are
obsolete now.

Reviewed-by: Z
Refs: #123
BREAKING CHANGE: use JavaScript features not available in Node 6.`)

func TestCommitMessage(t *testing.T) {

	t.Run("解析原始提交信息", func(t *testing.T) {
		message := NewCommitMessage(rawMsg)
		fmt.Println("v%", message)
	})
}
