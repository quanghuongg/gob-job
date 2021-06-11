package worker

import (
	"demo-job/job"
	"fmt"
)

type Worker struct {
	WorkerId   int
	Done       chan bool
	JobRunning chan job.Job
}

func NewWorker(workerId int, jobChan chan job.Job) *Worker {
	return &Worker{
		WorkerId:   workerId,
		Done:       make(chan bool),
		JobRunning: jobChan,
	}
}

func (w *Worker) Run () {
	fmt.Println("Run workerId id ", w.WorkerId)
	go func() {
		for {
			select {
			case job := <-w.JobRunning:
				fmt.Println("job running ", w.WorkerId)
				job.Process()

			case <-w.Done:
				fmt.Println("Stop worker")
				return
			}
		}
	}()
}

func (w *Worker) Stop () {
	w.Done <- true
}
