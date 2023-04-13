package agent

import (
	"context"

	"github.com/dxxu993/golangchain/library/llm"
	"github.com/dxxu993/golangchain/library/prompt"
)

type SimpleAgent struct {
	LLM          llm.LLM
	PromptHelper prompt.PromptHelper
}

func NewInitializeAgent(l llm.LLM, p prompt.PromptHelper) *SimpleAgent {
	agent := &SimpleAgent{
		LLM:          l,
		PromptHelper: p,
	}

	return agent
}

func (a *SimpleAgent) Run(ctx context.Context, userPrompt string, msgID string) (string, error) {
	refinePrompt := a.PromptHelper.GetPrompt(userPrompt)
	answer, err := a.LLM.Query(ctx, refinePrompt)
	return answer, err
}
