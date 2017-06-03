package main

import(
	"fmt"
)

// a single <-f is to trigger a recieve, no code after
//this will execute before this is done.


func worker(id int, input chan string, output chan string) {
	job := <-input
	fmt.Println("worker executing : ", job )
	output <- job
}

func main() {
	jobChannel := make(chan string)
	finishedJobsChanel := make(chan string)
	numJobs := 10
	for i := 0; i < numJobs; i++ {
		go worker(i, jobChannel, finishedJobsChanel)
	}
	for i := 0; i < numJobs; i++ {
		jobChannel <- "hello"  + string(i)
	}
	close(jobChannel)

}