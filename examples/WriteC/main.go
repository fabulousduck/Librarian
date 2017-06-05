package main

import(
	lr "github.com/fabulousduck/librarian"
	"fmt"
)

func main() {
	writeOpCount := 100;
	operationOutcomes, inputs := make(chan lr.WriteOpResponse), make(chan lr.WriteOp)
	go lr.WriteC(inputs, operationOutcomes)
	for i := 0; i < writeOpCount; i++ {
		inputs <- lr.WriteOp{ Dest: fmt.Sprintf("../exampleFiles/createdFiles/%d.txt", i), Content: fmt.Sprintf("Invoice #%d", i) }
	}
	close(inputs)  
	for i := 0; i < writeOpCount; i++ {
		writeResult := <-operationOutcomes
		fmt.Println("Response from write operation : ", writeResult.Msg, "err ", writeResult.Err, "bytes written : ", writeResult.BytesWritten)
	}
	close(operationOutcomes)
}