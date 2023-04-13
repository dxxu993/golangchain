package llm

import (
	"context"
	"fmt"

	"github.com/dxxu993/golangchain/library/resource"
	openai "github.com/sashabaranov/go-openai"
)

const (
	GptErrIncorrectAPIKey = "Incorrect API key provided"
)

type OpenaiGtpLLM struct {
	Client    *openai.Client
	MaxTokens int
}

func NewOpenaiGtpLLM(apiKey string) *OpenaiGtpLLM {
	client := openai.NewClient(apiKey)

	return &OpenaiGtpLLM{
		Client:    client,
		MaxTokens: 800,
	}
}

func (l *OpenaiGtpLLM) CheckSkInvalid(ctx context.Context) error {
	_, err := l.Client.ListModels(ctx)
	return err
}

func (l *OpenaiGtpLLM) Query(ctx context.Context, q string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "你是能够解答人生困惑的佛祖",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: q,
			},
		},
		MaxTokens: l.MaxTokens,
	}
	resp, err := l.Client.CreateChatCompletion(ctx, req)
	if err != nil {
		resource.Logger.Warnf("req openai llm fail, req: %+v , resp: %+v, err: %s", req, resp, err)
		return "", fmt.Errorf("completion fail: %s", err.Error())
	}
	resource.Logger.Debugf("req openai llm, req: %+v , resp: %+v", req, resp)
	if len(resp.Choices) == 0 {
		resource.Logger.Warnf("choices is empty, req: %+v , resp: %+v", req, resp)
		return "", fmt.Errorf("completion fail: %s", err.Error())
	}

	return resp.Choices[0].Message.Content, nil
}
