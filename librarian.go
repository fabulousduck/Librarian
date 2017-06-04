package librarian

import(
	"os"
	"io/ioutil"
)


//archivist to do the actual operation
func archivist(incoming chan string, outGoing chan<- string) {
	//assign the incoming job to the job variable
	filePath := <-incoming
	//execute the job
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	//push the result of job to the outGoing channel
	outGoing <- string(data)
}

func ReadC(paths []string, outputChannel chan<- string) {
	//channel to push new jobs onto
	readJobsChannel := make(chan string)
	//set initial value of archivist count
	readJobCount := accountForDirs(paths)
	readPaths := getReadPaths(paths)
	//instantiate the workers
	//because they run on channels, they
	//are blocking until a job comes over 
	//the channel
	for i := 0; i < readJobCount; i++ {
		go archivist(readJobsChannel, outputChannel)
	}
	//push all jobs onto the job channel
	for i := 0; i < readJobCount; i++ {
		readJobsChannel <- readPaths[i]
	}
	//close the job channel when done
	close(readJobsChannel)

}

func accountForDirs(paths []string) int {
	totalPathCount := 0
	for i := 0; i < len(paths); i++ {
		if(isDir(paths[i])) {
			totalPathCount += getNumFilesInDir(paths[i])
		} else {
			totalPathCount++
		}
	}
	return totalPathCount
}

func isDir(path string) bool {
	fileStat, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	mode := fileStat.Mode()

	if(mode.IsDir()) {
		return true
	} else {
		return false
	}
}

func getNumFilesInDir(path string) int {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	return len(files)
}

func getReadPaths(paths []string) []string {
	allPaths := []string{}

	for i := 0; i < len(paths); i++ {
		if(isDir(paths[i])) {
			files, err := ioutil.ReadDir(paths[i])
			if err != nil {
				panic(err)
			}
			for _, file := range files {
				allPaths = append(allPaths, string(paths[i] + file.Name()))
			}
		} else {
			allPaths = append(allPaths, paths[i])
		}
	}

	return allPaths
}