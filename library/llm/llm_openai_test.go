package llm

import (
	"context"
	"testing"

	"github.com/dxxu993/golangchain/library/conf"
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

func TestOpenaiGtpLLM_Query(t *testing.T) {
	l := &OpenaiGtpLLM{}
	ctx := context.Background()
	q := "你好"
	got, err := l.Query(ctx, q)
	if err != nil {
		t.Errorf("OpenaiGtpLLM.Query() error = %v", err)
		return
	}

	t.Logf("got : %+v", got)
}
