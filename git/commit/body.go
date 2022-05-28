package git

// MessageBody 详细提交描述
type MessageBody struct {
	Description []string `json:"description,omitempty"`
}

func CommitMessageBodyFromLongDescription(descriptions []string) *MessageBody {
	b := new(MessageBody)
	b.Description = descriptions
	return b
}
