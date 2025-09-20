package main

import (
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"
)

func main() {
	if logger, err := zap.NewProduction(); err != nil {
		log.Fatalf("Could not use production zap logger")
	} else {
		loggerSync := func() {
			if err := logger.Sync(); err != nil {
				panic(err)
			}
		}
		defer loggerSync()
		s := fmt.Sprintf("Successfully containerized the backend code - %s", time.Now().String())
		logger.Info(s)
	}
}
