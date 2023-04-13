package agent

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/dxxu993/golangchain/library/llm"
	"github.com/dxxu993/golangchain/library/prompt"
	"github.com/dxxu993/golangchain/library/resource"
)

var waitReplyMsgTaskPool *WaitReplyMsgTaskPool

const (
	MsgTaskStatusRunning = 1
	MsgTaskStatusFinish  = 2
)

type LazyReplyAgent struct {
	LLM          llm.LLM
	PromptHelper prompt.PromptHelper
	MsgTaskPool  *WaitReplyMsgTaskPool
}

type WaitReplyMsgTaskPool struct {
	Mu       *sync.Mutex
	TaskPool map[string]*WaitReplyMsgTask
}

type WaitReplyMsgTask struct {
	MsgID            string
	FinishNotifyChan chan struct{} // 用于通知监听者处理结果，从llm接收完返回后，通知所有接听者
	ReplayContent    string        // 回复内容
	ReplayError      error         // llm返回错误信息
	Counter          atomic.Int32  // 记录当前任务被添加了几次
	StartTime        time.Time     // 任务第一次被添加的时间
}

func init() {
	waitReplyMsgTaskPool = &WaitReplyMsgTaskPool{
		Mu:       &sync.Mutex{},
		TaskPool: make(map[string]*WaitReplyMsgTask),
	}
}

// addTask 添加一个任务，相同msgID的任务同时只会有一个task在执行，执行完会通知给所有阻塞等待的addTask
func (a *LazyReplyAgent) addTask(ctx context.Context, msgID, q string) (string, error) {
	a.MsgTaskPool.Mu.Lock()
	task, ok := a.MsgTaskPool.TaskPool[msgID]
	a.MsgTaskPool.Mu.Unlock()

	if !ok {
		a.MsgTaskPool.Mu.Lock()
		task = &WaitReplyMsgTask{
			MsgID:            msgID,
			FinishNotifyChan: make(chan struct{}),
			ReplayContent:    "",
			Counter:          atomic.Int32{},
			StartTime:        time.Now(),
		}
		a.MsgTaskPool.TaskPool[msgID] = task
		a.MsgTaskPool.Mu.Unlock()
		go func() {
			a.runTask(ctx, msgID, q, task)
			a.MsgTaskPool.Mu.Lock()
			delete(a.MsgTaskPool.TaskPool, msgID)
			a.MsgTaskPool.Mu.Unlock()
		}()
	}
	task.Counter.Add(1)
	select {
	case <-task.FinishNotifyChan:
		resource.Logger.Infof("recv FinishNotifyChan sig, task: %+v", *task)
		return task.ReplayContent, task.ReplayError
	case <-ctx.Done():
		resource.Logger.Warnf("wait task is canceld: %s", ctx.Err())
		return "", ctx.Err()
	}
}

func (a *LazyReplyAgent) runTask(ctx context.Context, msgID, q string, task *WaitReplyMsgTask) {
	refinePrompt := a.PromptHelper.GetPrompt(q)
	answer, err := a.LLM.Query(ctx, refinePrompt)
	task.ReplayError = err
	task.ReplayContent = answer
	close(task.FinishNotifyChan)
}

func NewLazyReplyAgent(l llm.LLM, p prompt.PromptHelper) *LazyReplyAgent {
	agent := &LazyReplyAgent{
		LLM:          l,
		PromptHelper: p,
		MsgTaskPool:  waitReplyMsgTaskPool,
	}
	return agent
}

func (a *LazyReplyAgent) Run(ctx context.Context, userPrompt string, msgID string) (string, error) {
	return a.addTask(ctx, msgID, userPrompt)
}
