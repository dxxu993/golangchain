package prompt

// 生成prompt
type RawPromptHelper struct {
}

func NewRawPromptHelper() *RawPromptHelper {
	return &RawPromptHelper{}
}

func (p *RawPromptHelper) GetPrompt(q string) string {
	return q
}
