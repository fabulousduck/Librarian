package main

import(
	"fmt"
	lr "github.com/fabulousduck/librarian"
)

func main () {
	paths := []string{
		"../exampleFiles/A.txt",
		"../exampleFiles/B.txt",
		"../exampleFiles/C.txt",
	}

	fileOutput := make(chan string)
	go lr.ReadC(paths, fileOutput)
	for i := 0; i < 3; i++ {
		content := <-fileOutput
		fmt.Println("file contents : ", content)
	}
}