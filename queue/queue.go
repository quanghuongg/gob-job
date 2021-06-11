package queue

import (
	"demo-job/job"
	"demo-job/worker"
	"fmt"
)

type JobQueue struct {
	Workers    []*worker.Worker
	JobRunning chan job.Job
	Done       chan bool
}

func NewJobQueue(size int) JobQueue {
	workers := make([]*worker.Worker, size, size)
	jobRunning := make(chan job.Job)

	for i := 0; i < size; i++ {
		workers[i] = worker.NewWorker(i, jobRunning)
	}
	return JobQueue{
		Workers:    workers,
		JobRunning: jobRunning,
		Done:       make(chan bool),
	}
}
func (jp *JobQueue) Push(job job.Job) {
	jp.JobRunning <- job
}

func (jp *JobQueue) Start() {
	go func() {
		for i := 0; i < len(jp.Workers); i++ {
			jp.Workers[i].Run()
		}
	}()
	go func() {
		for {
			select {
			case <-jp.Done:
				for i := 0; i < len(jp.Workers); i++ {
					fmt.Println("Stop worker ", jp.Workers[i].WorkerId)
					jp.Workers[i].Stop()
				}
				return
			}

		}
	}()

}
func (jp *JobQueue) Stop()  {
	jp.Done <- true
}
