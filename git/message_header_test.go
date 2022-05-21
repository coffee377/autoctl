package git

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeaderFromTitle(t *testing.T) {

	t.Run("仅含类型和描述", func(t *testing.T) {
		title := HeaderFromTitle(":bug: 仅含类型和描述")
		assert.Equal(t, ":bug:", title.Type)
		assert.Equal(t, "", title.Scope)
		assert.Equal(t, "仅含类型和描述", title.Description)
		assert.Equal(t, false, title.Broken)
	})

	t.Run("含类型、范围和描述", func(t *testing.T) {
		title := HeaderFromTitle(":bug:(test): 含类型、范围和描述")
		assert.Equal(t, ":bug:", title.Type)
		assert.Equal(t, "test", title.Scope)
		assert.Equal(t, "含类型、范围和描述", title.Description)
		assert.Equal(t, false, title.Broken)
	})

	t.Run("含类型、范围、描述和破坏性变更标识", func(t *testing.T) {
		title := HeaderFromTitle(":bug:(test)!: 含类型、范围、描述和破坏性变更标识")
		assert.Equal(t, ":bug:", title.Type)
		assert.Equal(t, "test", title.Scope)
		assert.Equal(t, "含类型、范围、描述和破坏性变更标识", title.Description)
		assert.Equal(t, true, title.Broken)
	})
}
