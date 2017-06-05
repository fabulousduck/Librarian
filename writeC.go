package librarian

import(
	"os"
	"fmt"
)

type WriteOp struct {
	Dest, Content string
}

type WriteOpResponse struct {
	Msg error
	Err bool
	BytesWritten int
}

func WriteC (inputChannel <-chan WriteOp, outputChannel chan<- WriteOpResponse) {
	for workOp := range inputChannel {
		go writeWorker(workOp, outputChannel)
	}

}

func writeWorker (job WriteOp, outGoing chan<- WriteOpResponse) {
	file, err := os.OpenFile(job.Dest, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("err : ", err)
		outGoing <- WriteOpResponse{ Msg: err, Err: true, BytesWritten: 0 }
		return 
	}
	bytesWritten , err := file.WriteString(job.Content)
	if err != nil {
		outGoing <- WriteOpResponse{ Msg: err, Err: true, BytesWritten: 0 }
		return
	}
	outGoing <- WriteOpResponse{ Msg: nil, Err: false, BytesWritten: bytesWritten }
	return 
}