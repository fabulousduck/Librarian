package librarian

import(
	"fmt"
	"io"
)



type WriteOp struct {
	dest string
	content string
}


// a single <-f is to trigger a recieve, no code after
//this will execute before this is done.


//archivist to do the actual operation
func archivist(id int, incoming chan job, outGoing chan job) {
	//assign the incoming job to the job variable
	job := <-incoming
	//execute the job
	job.fn()
	//push the result of job to the outGoing channel
	outGoing <- job
}

func ReadC(paths []string, output chan<- string) {
	//channel to push new jobs onto
	jobChannel := make(chan job)
	//channel to push done jobs restults onto
	finishedJobsChanel := make(chan job)

	//instantiate the workers
	//because they run on channels, they
	//are blocking until a job comes over 
	//the channel
	for i := 0; i < len(paths); i++ {
		go worker(i, jobChannel, finishedJobsChanel)
	}
	//push all jobs onto the job channel
	for i := 0; i < numJobs; i++ {
		jobChannel <- jobs[i]
	}
	//close the job channel when done
	close(jobChannel)

}

func isDir(path string) bool {
	
}