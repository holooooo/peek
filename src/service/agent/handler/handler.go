package handler

import (
	peekv1 "peek/src/gen/peek/v1"
	"peek/src/service/agent/handler/impl"
)

type CreateHandler func(job *peekv1.Job) impl.Handler

var (
	factory = map[peekv1.JobCommand]CreateHandler{}
)

func init() {
	factory[peekv1.JobCommand_binary] = impl.NewBinaryHandler
}

func NewHandler(job *peekv1.Job) impl.Handler {
	return factory[job.Command](job)
}