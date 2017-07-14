# Librarian
A tiny Go package for concurrent file operations.

## How to use

Librarian is simple in use because it has very few functions.

### ` ReadC([]string, chan string) `

ReadC is for reading multiple files concurrently
By providing a `string slice` containing the paths to the files.

If a given path is a directory, all files in it will be read and returned
on the outgoing channel.

*example*

```go
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
```   
   
### `WriteC(chan lr.WriteOp, chan bool)`

WriteC is for writing to a large amount of files concurrently
by pushing `lr.writeOp`'s to a channel passed to `WriteBatch`.

The channel will return  `n` amount of `boolean`'s representing if
the operation succeeded or not.

example

```go
   package main
   
   import(
      lr "fabulousduck/librarian"
   )
   
   func main() {
      operationOutcomes := make(chan, bool)
      inputs := make(chan, lr.WriteOp)
      go lr.WriteC(inputs, operationOutcomes)
      for i := 0; i < 100; i++ {
        inputs <- lr.WriteOp{ dest: "~/Desktop/invoices/"+i+".txt", content: "invoice #"+i }
      }  
      outcomes := <- operationOutcomes
   }
