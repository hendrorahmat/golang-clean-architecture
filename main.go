package main

import "github.com/hendrorahmat/golang-clean-architecture/cmd"

func main() {
	cmd.Execute()

	//blocking := make(chan bool)
	//worker := utils.NewWorkerPool()
	//const numTasks = 10
	//type Job struct {
	//	id   int
	//	name string
	//}
	//for i := 0; i < numTasks; i++ {
	//	result := fmt.Sprintf("Data ke %d ", i)
	//	job := Job{name: "job " + result, id: i}
	//
	//	worker.AddTask(func(ctx context.Context) (any, error) {
	//		return job, nil
	//	})
	//}
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*411)
	//defer cancel()
	//worker.Run(ctx)
	//fmt.Println(worker.GetResults())
	//<-blocking
}
