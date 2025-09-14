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
		defer logger.Sync()
		s := fmt.Sprintf("Successfully containerized the backend code - %s", time.Now().String())
		logger.Info(s)
	}
}
