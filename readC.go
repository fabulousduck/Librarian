package librarian


func ReadC(paths []string, outputChannel chan<- string) {
	//channel to push new jobs onto
	readJobsChannel := make(chan string)
	//set initial value of archivist count
	readJobCount := accountForDirs(paths)
	readPaths := expandPaths(paths)
	//instantiate the workers
	//because they run on channels, they
	//are blocking until a job comes over 
	//the channel
	for i := 0; i < readJobCount; i++ {
		go readWorker(readJobsChannel, outputChannel)
	}
	//push all jobs onto the job channel
	for i := 0; i < readJobCount; i++ {
		readJobsChannel <- readPaths[i]
	}
	//close the job channel when done
	close(readJobsChannel)

}