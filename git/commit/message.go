package git

import (
	"regexp"
	"strings"
)

// CommitMessage git 提交信息
type CommitMessage struct {
	Raw    []byte         `json:"raw,omitempty"`    //
	Header *MessageHeader `json:"header,omitempty"` // 头部信息
	Body   *MessageBody   `json:"body,omitempty"`   // 长描述
	Footer *MessageFooter `json:"footer,omitempty"` // 脚注信息
	OutRaw bool           `json:"-"`                // json 序列化是否输出原始数据
}

var BlankLine = "[\r\n]\n"

// NewCommitMessage creates an instance of CommitMessage.
func NewCommitMessage(raw []byte) *CommitMessage {
	msg := &CommitMessage{}
	// raw,omitempty
	return msg.FromRaw(raw)
}

func (message *CommitMessage) FromRaw(raw []byte) *CommitMessage {
	message.Raw = raw
	var (
		title           string
		longDescription []string
		footers         []string
	)

	//reflect.ValueOf(message).Field(0).String()
	//reflect.ValueOf(message).Elem().

	// 按空行进行分割
	splits := regexp.MustCompile(BlankLine).Split(string(raw), -1)
	pos := len(splits) - 1
	// 获取标题
	title = splits[:1][0]

	if len(splits) > 1 {
		// todo 要验证最后一项是否真的是脚注，可能是长描述
		// 获取脚注
		footers = strings.Split(splits[pos:][0], "\n")

		// 获取长描述
		longDescription = splits[1:pos]
	}

	message.dealHeader(title)
	message.dealBody(longDescription)
	message.dealFooter(footers)
	return message
}

func (message *CommitMessage) dealHeader(title string) {
	header := CommitMessageHeaderFromTitle(title)
	message.Header = header
}

func (message *CommitMessage) dealBody(longDescription []string) {
	body := CommitMessageBodyFromLongDescription(longDescription)
	message.Body = body
}

func (message *CommitMessage) dealFooter(footers []string) {
	footer := CommitMessageFooterFromNotes(footers)
	message.Footer = footer
}
