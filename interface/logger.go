package _interface

type Logger interface {
	Log(message string) error
	Close() error
}
