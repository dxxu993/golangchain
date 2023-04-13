package prompt

type PromptHelper interface {
	// 获取适配了指定模板的prompt
	GetPrompt(q string) string
}
