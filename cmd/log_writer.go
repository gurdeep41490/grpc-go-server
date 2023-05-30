package main

import (
	"fmt"
	"time"
)

type logWriter struct{}

func (writer *logWriter) Write(bytes []byte) (int, error) {
	return fmt.Println(time.Now().Format("15:04:02") + " " + string(bytes))
}
