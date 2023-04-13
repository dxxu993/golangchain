package agent

import (
	"github.com/dxxu993/golangchain/library/llm"
	"github.com/dxxu993/golangchain/library/prompt"
)

type BatchsReplyAgent struct {
	LLM          llm.LLM
	PromptHelper prompt.PromptHelper
	BatchSize    int // 每批次的字节数
}

// BatchsReplyAgent todo 支持从大语言模型的流式token积攒到指定个数后成为一个批次返回，直到流返回结束，因此整个过程可能会产生多个批次
func NewBatchsReturnAgent(l llm.LLM, p prompt.PromptHelper, batchSize int) *BatchsReplyAgent {
	agent := &BatchsReplyAgent{
		LLM:          l,
		PromptHelper: p,
		BatchSize:    batchSize,
	}

	return agent
}
