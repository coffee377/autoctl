package git

import (
	"regexp"
	"strings"
)

// MessageHeader <type>[(scope)][!]: <description>
type MessageHeader struct {
	Type        string `json:"type,omitempty"`        // 类型
	Scope       string `json:"scope,omitempty"`       // 范围（可选）
	Broken      bool   `json:"broken,omitempty"`      // 标记为破坏性变更（可选）
	Description string `json:"description,omitempty"` // 简要描述
	Semver      string `json:"semver,omitempty"`      // 语义化版本
}

func CommitMessageHeaderFromTitle(title string) *MessageHeader {
	h := new(MessageHeader)
	reg := regexp.MustCompile("^(:?\\w+:?)(\\((\\w*)\\))?(!)?:? (.+)$")
	match := reg.FindStringSubmatch(strings.Trim(title, " "))

	h.setType(match[1])

	if len(match[3]) > 0 {
		h.setScope(match[3])
	}

	if match[4] == "!" {
		h.setBroken(true)
	}

	h.setDescription(match[5])

	return h
}

func NewCommitMessageHeader(msgType string, description string) *MessageHeader {
	h := &MessageHeader{
		Type:        msgType,
		Scope:       "",
		Broken:      false,
		Description: description,
	}
	return h
}

func (h *MessageHeader) setType(msgType string) *MessageHeader {
	h.Type = msgType
	return h
}

func (h *MessageHeader) setScope(scope string) *MessageHeader {
	h.Scope = scope
	return h
}

func (h *MessageHeader) setBroken(broken bool) *MessageHeader {
	h.Broken = broken
	return h
}

func (h *MessageHeader) setDescription(description string) *MessageHeader {
	h.Description = description
	return h
}
