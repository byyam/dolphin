package processor

import (
	"context"
	hello "github.com/byyam/dolphin/proto"
)

type HelloProcessor struct {
}

func NewHelloProcessor() *HelloProcessor {
	return &HelloProcessor{}
}

func (p HelloProcessor) Greeting(ctx context.Context, req *hello.GreetingRequest) (*hello.GreetingResponse, error) {
	answer := "hello "
	if req != nil {
		answer = answer + req.GetQuestion()
	}
	rsp := &hello.GreetingResponse{
		Answer: answer,
	}
	return rsp, nil
}
