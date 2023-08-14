package git

import (
	"strings"
)

// Footnote 脚注备注项
type Footnote struct {
	Token string `json:"token"`
	Value string `json:"value"`
}

func NewFootnote(raw string) Footnote {
	match := strings.Split(raw, ": ")
	footnote := Footnote{
		Token: match[0],
		Value: match[1],
	}
	return footnote
}

// MessageFooter <token>: <value>
// https://www.conventionalcommits.org/
// 每行脚注都必须包含一个令牌（token），后面紧跟 :<space> 或 <space># 作为分隔符，后面再紧跟令牌的值
// 脚注的令牌必须使用 - 作为连字符，比如 Acked-by (这样有助于 区分脚注和多行正文)
// 有一种例外情况就是 BREAKING CHANGE，它可以被认为是一个令牌
// 脚注的值可以包含空格和换行，值的解析过程必须直到下一个脚注的令牌/分隔符出现为止
// A footer’s token MUST use - in place of whitespace characters, e.g., Acked-by (this helps differentiate the footer section from a multi-paragraph body). An exception is made for BREAKING CHANGE, which MAY also be used as a token.
// Reviewed-by: Z
// Refs: #123
// Closes #123,234,456
// BREAKING CHANGE: use JavaScript features not available in Node 6.
type MessageFooter struct {
	Items          []Footnote `json:"items,omitempty"`          // 其他相关的脚注信息
	Closes         []string   `json:"closes,omitempty"`         // 关闭相关 issues 或 bugs
	Refs           []string   `json:"refs,omitempty"`           // 相关需求引用
	BreakingChange string     `json:"breakingChange,omitempty"` // 破坏性变更信息
}

func CommitMessageFooterFromNotes(notes []string) *MessageFooter {
	footer := new(MessageFooter)
	for _, n := range notes {
		footer.AddFooterItem(n)
	}
	return footer
}

func (f *MessageFooter) AddFooterItem(rawItem string) *MessageFooter {
	footerItem := NewFootnote(rawItem)
	if footerItem.Token == "Closes" || footerItem.Token == "Refs" {
		f.Closes = strings.Split(footerItem.Value, ",")
	} else if footerItem.Token == "BREAKING CHANGE" {
		f.BreakingChange = footerItem.Value
	} else {
		items := append(f.Items, footerItem)
		f.Items = items
	}
	return f
}
