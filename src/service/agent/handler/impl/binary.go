package impl

import (
	"os/exec"
	peekv1 "peek/src/gen/peek/v1"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type binaryHandler struct {
	job    *peekv1.Job
	logger *logrus.Entry
}

func NewBinaryHandler(job *peekv1.Job) Handler {
	return &binaryHandler{job: job,
		logger: logrus.
			WithField("job", job.Name).
			WithField("cmd", strings.Join(job.Args, " "))}
}

func (h *binaryHandler) Exec() *peekv1.JobResult {
	h.job.Status = peekv1.JobStatus_running

	cmd := exec.Command(h.job.Args[0], h.job.Args[1:]...)
	err := cmd.Start()
	if err != nil {
		h.logger.Errorf("Failed to execute command: %v", err)
		h.job.Status = peekv1.JobStatus_failed
	}

	// execute the command
	if h.job.Config.Timeout > 0 {
		timeout := time.After(time.Duration(h.job.Config.Timeout) * time.Second)
		func() {
			for {
				select {
				case <-timeout:
					h.logger.Warnf("Command timeout: %v", h.job.Config.Timeout)
					h.job.Status = peekv1.JobStatus_completed
					cmd.Process.Kill()
					return
				default:
					if cmd.ProcessState != nil && cmd.ProcessState.Exited() {
						h.logger.Infof("Command exited with status: %v", cmd.ProcessState.ExitCode())
						h.job.Status = peekv1.JobStatus_completed
						return
					}
					time.Sleep(time.Second)
				}
			}
		}()
	} else {
		err = cmd.Wait()
	}
	if err != nil {
		h.logger.Errorf("Failed to execute command: %v", err)
	}

	// get the output
	exitCode := -1
	if cmd.ProcessState != nil {
		exitCode = cmd.ProcessState.ExitCode()
	}
	output, err := cmd.Output()
	if err != nil {
		h.logger.Errorf("Failed to read command stdout: %v", err)
	}
	return &peekv1.JobResult{
		Result:   &peekv1.JobResult_Stdout{Stdout: string(output)},
		Error:    err.Error(),
		ExitCode: int32(exitCode),
	}
}