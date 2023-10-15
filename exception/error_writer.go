package exception

import (
	"os"
	"sync"
)

var ErrorLogWriter = &ErrorLogWriterImpl{FileName: "storage/logs/error.log"}

type ErrorLogWriterImpl struct {
	FileName string
	mu       sync.Mutex
}

func (writer *ErrorLogWriterImpl) Write(p []byte) (n int, err error) {
	file, err := os.OpenFile(writer.FileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	writer.mu.Lock()
	defer writer.mu.Unlock()

	return file.Write(p)
}
