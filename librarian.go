package librarian

import(
	"os"
	"io/ioutil"
)

// <-chan is read only
// chan<- is write only

//archivist to do the actual operation
func readWorker(incoming chan string, outGoing chan<- string) {
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
 

//TODO refactor this to be a concurrent read
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

func expandPaths(paths []string) []string {
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