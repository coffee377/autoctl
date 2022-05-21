package git

import (
	"regexp"
	"strings"
)

// Header <type>[(scope)][!]: <description>
type Header struct {
	Type        string `json:"type"`        // 类型
	Scope       string `json:"scope"`       // 范围（可选）
	Broken      bool   `json:"broken"`      // 标记为破坏性变更（可选）
	Description string `json:"description"` // 简要描述
}

func HeaderFromTitle(title string) *Header {
	h := new(Header)
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

func NewHeader(msgType string, description string) *Header {
	h := &Header{
		Type:        msgType,
		Scope:       "",
		Broken:      false,
		Description: description,
	}
	return h
}

func (h *Header) setType(msgType string) *Header {
	h.Type = msgType
	return h
}

func (h *Header) setScope(scope string) *Header {
	h.Scope = scope
	return h
}

func (h *Header) setBroken(broken bool) *Header {
	h.Broken = broken
	return h
}

func (h *Header) setDescription(description string) *Header {
	h.Description = description
	return h
}
