package impl

import peekv1 "peek/src/gen/peek/v1"

type Handler interface {
	Exec() *peekv1.JobResult
	// TODO stream result
}