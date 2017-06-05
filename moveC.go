package librarian

import(
	"errors"
	"os"
)

type MoveOp struct {
	Origin, Dest string
}

type MoveOpResponse struct {
	Msg error
	Err bool
}

func MoveC (inputChannel <-chan MoveOp, outputChannel chan<- MoveOpResponse) {
	for moveOp := range inputChannel {
		go moveFile(moveOp)
	}
}

func moveFile (job MoveOp, outputChannel chan<- MoveOpResponse) {
	if(!isFile(job.origin) || !isDir(job.Dest)) {
		outputChannel <- MoveOpResponse{ Msg: Errors.New("origin is not file or dest is not directory.", Err: true) }
		return
	}

	err := os.Rename(job.Origin, job.Dest)
	if err != nil {
		outputChannel <- MoveOpResponse{ Msg: err, Err: true }
		return
	}
	outputChannel <- MoveOpResponse{Msg: nil, Err: false}
}