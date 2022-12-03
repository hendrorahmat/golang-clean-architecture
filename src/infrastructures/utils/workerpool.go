package utils

import (
	"context"
	"fmt"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/constants"
	"sync"
	"time"
)

type task func(ctx context.Context) (any, error)
type result struct {
	data  any
	error error
}
type WorkerPool struct {
	tasks           []task
	results         []result
	jobs            chan task
	wg              *sync.WaitGroup
	ctx             context.Context
	TotalSuccessJob int
	workersCount    int
}

var resultPool = sync.Pool{New: func() any {
	return new(result)
}}

func (wa *WorkerPool) WaitWorker() {
	wa.wg.Wait()
}

func worker(jobs <-chan task, wa *WorkerPool) <-chan result {
	wa.wg.Add(wa.workersCount)
	chanOut := make(chan result)
	var rp *result
	for i := 0; i < wa.workersCount; i++ {
		go func(workerId int) {
			fmt.Printf("Worker id %d started \n", workerId)
			for job := range jobs {
				select {
				case <-wa.ctx.Done():
					rp.data = nil
					rp.error = wa.ctx.Err()
					resultPool.Put(rp)
					fmt.Println("Break")
					chanOut <- *rp
					break
				default:
					fmt.Printf("Worker id %d processing \n", workerId)
					rp = resultPool.Get().(*result)
					data, err := job(wa.ctx)
					fmt.Printf("Worker id %d finished \n", workerId)
					rp.data = data
					rp.error = err
					resultPool.Put(rp)
					chanOut <- *rp
				}
			}
			fmt.Printf("Worker id %d closed \n", workerId)
			wa.wg.Done()
		}(i)
	}

	go func() {
		fmt.Println("Wait group")
		wa.WaitWorker()
		fmt.Println("Wait group done")
		close(chanOut)
	}()
	return chanOut
}

func (wa *WorkerPool) dispatchJob() <-chan task {
	wa.jobs = make(chan task)
	go func() {
		for i, t := range wa.tasks {
			time.Sleep(1 * time.Second)
			select {
			case <-wa.ctx.Done():
				fmt.Println("Job dispatch canceled")
				break
			default:
				i++
				fmt.Printf("dispatching job %d ... \n", i)
				wa.jobs <- t
			}
		}
		fmt.Println("Job dispatched successfully")
		close(wa.jobs)
		fmt.Println("Job channel was closed")
	}()
	fmt.Println("return jobs")
	return wa.jobs
}

func (wa *WorkerPool) Run(ctx context.Context) {
	wa.ctx = ctx
	done := make(chan int)
	go func() {
		jobs := wa.dispatchJob()
		w := worker(jobs, wa)
		counterSuccess := 0
		for data := range w {

			counterSuccess++
			wa.results = append(wa.results, data)
			fmt.Println("Total processed job ", counterSuccess)
		}
		done <- counterSuccess
	}()

	select {
	case <-wa.ctx.Done():
		fmt.Println("Proccess Stopped ", ctx.Err())
	case totalSuccess := <-done:
		close(done)
		wa.TotalSuccessJob = totalSuccess
		fmt.Println("Total Keseluruhan ", totalSuccess)
	}
}

func (wa *WorkerPool) AddTask(task task) {
	wa.tasks = append(wa.tasks, task)
}

func (wa *WorkerPool) SetWorkerNumber(total int) {
	wa.workersCount = total
}

func (wa *WorkerPool) GetTotalSuccessJob() int {
	return wa.TotalSuccessJob
}
func (wa *WorkerPool) GetResults() []result {
	return wa.results
}

type IWorkerPool interface {
	Run(ctx context.Context)
	WaitWorker()
	GetResults() []result
	AddTask(task task)
	SetWorkerNumber(total int)
	GetTotalSuccessJob() int
}

func NewWorkerPool() IWorkerPool {
	return &WorkerPool{wg: new(sync.WaitGroup), workersCount: constants.TotalWorkerMax}
}
