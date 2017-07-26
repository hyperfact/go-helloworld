package main

import (
	log "github.com/op/go-logging"
	"os"
)

var logger = log.MustGetLogger("net")

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func main() {
	fmtr := log.MustStringFormatter(`%{color}[%{level:.5s}]%{time:2006-01-02 15:04:05} %{shortfile} ▶ %{id:03x}%{color:reset} %{message}`)
	var be log.Backend
	be = log.NewLogBackend(os.Stdout, "", 0)
	be = log.NewBackendFormatter(be, fmtr)
	lbe := log.AddModuleLevel(be)
	lbe.SetLevel(log.DEBUG, "")
	log.SetBackend(be)
	logger.ExtraCalldepth = 1

	Debug("this is a debug log")
}
