package main

import(
	"fmt"
)


type job struct {
	fn func(string, string) string
	from, to string
}

// a single <-f is to trigger a recieve, no code after
//this will execute before this is done.

//worker to perform the task its given
func worker(id int, incoming chan job, outGoing chan job) {
	//assign the incoming job to the job variable
	job := <-incoming
	//execute the job
	job.fn(job.from, job.to)
	//push the result of job to the outGoing channel
	outGoing <- job
}

func main() {
	//channel to push new jobs onto
	jobChannel := make(chan job)
	//channel to push done jobs restults onto
	finishedJobsChanel := make(chan job)
	//amount of jobs to be done
	numJobs := 2000
	//create the jobs
	jobs := createJobs(numJobs)

	//instantiate the workers
	//because they run on channels, they
	//are blocking until a job comes over 
	//the channel
	for i := 0; i < numJobs; i++ {
		go worker(i, jobChannel, finishedJobsChanel)
	}
	//push all jobs onto the job channel
	for i := 0; i < numJobs; i++ {
		jobChannel <- jobs[i]
	}
	//close the job channel when done
	close(jobChannel)

}

//helper function to generate n amount of jobs
func createJobs(numJobs int) []job {
	jobs := []job{}
	for i := 0; i < numJobs; i++ {
		jobs = append(jobs, job{doThing, "a", "b"})
	}
	return jobs
}

//job function to be performed
func doThing(f string, t string) string{
	fmt.Println("moving file from ", f, "to", t)
	return t
}