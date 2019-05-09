package main

import (
	"log"
	"os"
	"time"
)

func setLog(level string, msg string) {
	l := log.New(os.Stdout, "", 0)
	l.SetPrefix(time.Now().Format(defaultTimeFormat) + " [" + level + "] ")
	if level == "critical" {
		l.Fatal(msg)
	} else {
		l.Print(msg)
	}
}
