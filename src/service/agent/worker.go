package agent

import (
	"os"
	peekv1 "peek/src/gen/peek/v1"
	"peek/src/service/agent/handler"
	"strconv"
)

var (
	asyncWorkerNum = os.Getenv("AGENT_ASYNC_WORKER_NUM")
)

type Worker struct {
	asyncJobCh chan *peekv1.Job
	syncJobCh  chan *peekv1.Job
	stop       chan struct{}

	asyncWorkerNum int
	jobMap         map[string]*peekv1.Job
	jobs           []*peekv1.Job
}

func NewWorker(stop chan struct{}) *Worker {
	asyncWorkerNum, err := strconv.ParseInt(asyncWorkerNum, 10, 64)
	if err != nil {
		logger.Fatalf("parse async worker num error: %v", err)
	}

	w := &Worker{
		asyncJobCh:     make(chan *peekv1.Job, 1024),
		syncJobCh:      make(chan *peekv1.Job, 0),
		stop:           stop,
		asyncWorkerNum: int(asyncWorkerNum),

		jobMap: make(map[string]*peekv1.Job),
		jobs:   make([]*peekv1.Job, 0),
	}
	return w
}

func (w *Worker) Run() {
	for i := 0; i < w.asyncWorkerNum; i++ {
		go w.asyncRun()
	}
	w.syncRun()
}

func (w *Worker) asyncRun() {
	for job := range w.asyncJobCh {
		go func(job *peekv1.Job) {
			job.Result = handler.NewHandler(job).Exec()
		}(job)

		select {
		case <-w.stop:
			return
		default:
		}
	}
}

func (w *Worker) syncRun() {
	for job := range w.syncJobCh {
		job.Result = handler.NewHandler(job).Exec()

		select {
		case <-w.stop:
			return
		default:
		}
	}
}

func (w *Worker) AddJob(job *peekv1.Job) {
	w.asyncJobCh <- job
	w.jobMap[job.Name] = job
	w.jobs = append(w.jobs, job)
}

func (w *Worker) GetJobs() []*peekv1.Job {
	return w.jobs
}

func (w *Worker) GetJob(name string) *peekv1.Job {
	return w.jobMap[name]
}