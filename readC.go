package librarian

import(
	"io/ioutil"
)

func ReadC(paths []string, outputChannel chan<- string) {
	//channel to push new jobs onto
	readJobsChannel := make(chan string)
	readJobCount := accountForDirs(paths)
	readPaths := expandPaths(paths)

	for i := 0; i < readJobCount; i++ {
		go readWorker(readJobsChannel, outputChannel)
	}

	for i := 0; i < readJobCount; i++ {
		readJobsChannel <- readPaths[i]
	}

	close(readJobsChannel)

}

func readWorker(incoming chan string, outGoing chan<- string) {
	filePath := <-incoming

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	outGoing <- string(data)
}