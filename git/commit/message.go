package git

import (
	"regexp"
	"strings"
)

// CommitMessage git 提交信息
type CommitMessage struct {
	Raw             []byte         `json:"raw"`    //
	Header          *MessageHeader `json:"header"` // 头部信息
	Body            *MessageBody   `json:"body"`   // 长描述
	Footer          *MessageFooter `json:"footer"` // 脚注信息
	title           string         ``              //
	longDescription []string       ``              // 长描述
	footers         []string       ``              // 脚注
}

var BlankLine = "[\r\n]\n"

// NewCommitMessage creates an instance of CommitMessage.
func NewCommitMessage(raw []byte) *CommitMessage {
	msg := &CommitMessage{}
	return msg.FromRaw(raw)
}

func (message *CommitMessage) FromRaw(raw []byte) *CommitMessage {
	message.Raw = raw
	// 按空行进行分割
	splits := regexp.MustCompile(BlankLine).Split(string(raw), -1)
	footerPos := len(splits) - 1
	// 获取标题
	title := splits[:1][0]
	message.title = title

	// 获取脚注
	footers := splits[footerPos:][0]
	message.footers = strings.Split(footers, "\n")

	// 获取长描述
	message.longDescription = splits[1:footerPos]

	message.dealHeader()
	message.dealBody()
	message.dealFooter()
	return message
}

func (message *CommitMessage) dealHeader() {
	header := CommitMessageHeaderFromTitle(message.title)
	message.Header = header
}

func (message *CommitMessage) dealBody() {
	body := CommitMessageBodyFromLongDescription(message.longDescription)
	message.Body = body
}

func (message *CommitMessage) dealFooter() {
	footer := CommitMessageFooterFromNotes(message.footers)
	message.Footer = footer
}
