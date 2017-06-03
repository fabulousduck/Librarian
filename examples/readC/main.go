package main

import(
	lr "fabulousduck/librarian"
)

func main () {

}


func createJobs(numJobs int) []job {
	jobs := []job{}
	for i := 0; i < numJobs; i++ {
		jobs = append(jobs, job{doThing, []string{"a","b"}})
	}
	return jobs
}