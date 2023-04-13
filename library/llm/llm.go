package llm

import "context"

type LLM interface {
	Query(ctx context.Context, query string) (anser string, err error)
}
