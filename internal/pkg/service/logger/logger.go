package logger

import (
	"fmt"

	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/loggers"
)

var (
	Logger = new(loggers.Logger)
	Debug  = Logger.Debug
	Error  = Logger.Error
	Fatal  = Logger.Fatal
	Info   = Logger.Info
)

func Starting(service string) {
	Info(fmt.Sprintf("starting %s...", service))
}

func StartingModule(module string) {
	Starting("module:" + module)
}

func Started(service string) {
	Info(fmt.Sprintf("%s started", service))
}

func StartedModule(module string) {
	Started("module:" + module)
}

func Stopping(service string) {
	Info(fmt.Sprintf("stopping %s...", service))
}

func Stopped(service string) {
	Info(fmt.Sprintf("%s stopped", service))
}

func EstablishingConnectionWith(service string) {
	Info(fmt.Sprintf("establishing connection with %s...", service))
}

func ConnectionEstablishedWith(service string) {
	Info(fmt.Sprintf("connection established with %s", service))
}

func ClosingConnectionWith(service string) {
	Info(fmt.Sprintf("closing connection with %s...", service))
}

func ConnectionClosedWith(service string) {
	Info(fmt.Sprintf("connection closed with %s", service))
}
