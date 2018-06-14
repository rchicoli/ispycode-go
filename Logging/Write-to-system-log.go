package main

import (
	"log"
	"log/syslog"
)

func main() {

	logwriter, e := syslog.New(syslog.LOG_NOTICE, "myprog")
	if e == nil {
		log.SetOutput(logwriter)
	}

	log.Print("Hello System Log")
}
