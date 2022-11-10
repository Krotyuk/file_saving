package file_logger

import "os"

type FileLogger struct{ file *os.File }

func New(fileName string) (*FileLogger, error) {
	f, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return &FileLogger{f}, nil
}
func (f FileLogger) Log(message string) error {
	_, err := f.file.WriteString(message + "\n")
	return err
}
func (f FileLogger) Close() error {
	return f.file.Close()
}
