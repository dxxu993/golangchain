package agent

import (
	"context"
)

type Agent interface {
	Run(ctx context.Context, prompt string, msgID string) (string, error)
}
