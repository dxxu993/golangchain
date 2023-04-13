package llm

import (
	"context"
	"time"
)

type EchoLLM struct {
	waitSec int
}

// NewEchoLLM 用于调试的一个虚拟llm
func NewEchoLLM(waitSec int) *EchoLLM {
	return &EchoLLM{waitSec: waitSec}
}

func (l *EchoLLM) Query(ctx context.Context, q string) (string, error) {
	time.Sleep(time.Duration(l.waitSec) * time.Second)
	return q, nil
}
