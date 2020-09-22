package service

import (
	"fmt"
	"os"
)

type Job interface {
	Do() error
}

// 定义一下job队列和work池类型
// define job channel
type JobChan chan Job

// define worker channer
type WorkerChan chan JobChan

var (
	JobQueue   JobChan
	WorkerPool WorkerChan

	MaxWorker = os.Getenv("MAX_WORKERS")
	MaxQueue  = os.Getenv("MAX_QUEUE")
)

type Worker struct {
	JobChannel JobChan
	quit       chan bool
}

func (w *Worker) Start() {
	go func() {
		for {
			// regist current job channel to worker pool
			WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				if err := job.Do(); err != nil {
					fmt.Printf("excute job failed with err: %v", err)
				}
			case <-w.quit:
				return
			}
		}
	}()
}

type Dispatcher struct {
	Workers []*Worker
	quit    chan bool
}

/*func (d *Dispatcher) Run() {
	for i := 0; i < len(MaxWorker); i++ {
		worker := newWorker()
		d.Workers = append(d.Workers, worker)
		worker.Start()
	}

	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				jobChan := <-WorkerPool
				jobChan <- job
			}(job)
		// stop dispatcher
		case <-d.quit:
			return
		}
	}
}*/
