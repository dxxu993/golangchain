package agent

import (
	"context"
	"sync"
	"testing"

	"github.com/dxxu993/golangchain/library/conf"
	"github.com/dxxu993/golangchain/library/llm"
	"github.com/dxxu993/golangchain/library/prompt"
	"github.com/dxxu993/golangchain/library/resource"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	conf.RelRootPath = "../../"
	resource.Bootstrap(ctx)
	//  = logger.InitLogger()
	// resource.AppConf = conf.InitAppConf()
	// resource.Logger = logger.InitLogger()
	m.Run()
}

func TestLazyReplyAgent_Run(t *testing.T) {
	ctx := context.Background()
	l := llm.NewEchoLLM(1)
	prompptHelp := prompt.NewRawPromptHelper()
	a := NewLazyReplyAgent(l, prompptHelp)
	q := "你好"
	msgID := "111222"
	got, err := a.Run(ctx, q, msgID)
	if err != nil {
		t.Errorf("LazyReplyAgent.Run() error = %v", err)
		return
	}
	t.Logf("got: %+v", got)
}

func TestLazyReplyAgent_Run_Mul(t *testing.T) {
	ctx := context.Background()
	l := llm.NewEchoLLM(1)
	prompptHelp := prompt.NewRawPromptHelper()
	a := NewLazyReplyAgent(l, prompptHelp)
	q := "你好"
	msgID := "111222"
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		got, err := a.Run(ctx, q, msgID)
		if err != nil {
			t.Errorf("LazyReplyAgent.Run() error = %v", err)
			return
		}
		t.Logf("got: %+v", got)
	}()
	go func() {
		defer wg.Done()
		got, err := a.Run(ctx, q, msgID)
		if err != nil {
			t.Errorf("LazyReplyAgent.Run() error = %v", err)
			return
		}
		t.Logf("got: %+v", got)
	}()
	go func() {
		defer wg.Done()
		got, err := a.Run(ctx, q, msgID)
		if err != nil {
			t.Errorf("LazyReplyAgent.Run() error = %v", err)
			return
		}
		t.Logf("got: %+v", got)
	}()
	wg.Wait()
}
