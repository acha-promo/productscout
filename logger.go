package core

import "log"

type (
	DefaultLogger struct{}

	Logger interface {
		Info(args ...interface{})
		Error(args ...interface{})
		Debug(args ...interface{})
	}
)

func (l DefaultLogger) Info(args ...interface{}) {
	log.Println(args...)
}

func (l DefaultLogger) Error(args ...interface{}) {
	log.Println(args...)
}

func (l DefaultLogger) Debug(args ...interface{}) {
	log.Println(args...)
}
