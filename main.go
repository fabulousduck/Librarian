package main

import(
	"fmt"
	"math/rand"
)

type job struct {
	message int
}

// a single <-f is to trigger a recieve, no code after
//this will execute before this is done.


func worker(id int, input chan job, output chan job) {
	job := <-input
	fmt.Println("worker executing : ", job.message )
	output <- job
}

func main() {
	jobChannel := make(chan job)
	finishedJobsChanel := make(chan job)
	numJobs := 10
	jobs := createJobs(numJobs)

	for i := 0; i < numJobs; i++ {
		go worker(i, jobChannel, finishedJobsChanel)
	}
	for i := 0; i < numJobs; i++ {
		jobChannel <- jobs[i]
	}
	close(jobChannel)

}

func createJobs(numJobs int) []job {
	jobs := []job{}
	for i := 0; i < numJobs; i++ {
		jobs = append(jobs, job{message: rand.Intn(500)})
	}
	return jobs
}