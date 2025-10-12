package concurrency

import (
	"fmt"
	"sync"
)

type Config struct {
	DBConfig *DBConfig
}

type DBConfig struct {
	Port     int
	User     string
	Database string
	Host     string
}

var once sync.Once
var Instance *Config

func getConfig() *Config {
	once.Do(func() {
		d := &DBConfig{
			// read from the disk
			Port:     8080,
			User:     "localuser",
			Database: "local_db",
			Host:     "localhost",
		}
		Instance = &Config{
			DBConfig: d,
		}
		fmt.Println("Initialized once")
	})
	return Instance
}

func ThreadSafeSingleton() {
	var wg sync.WaitGroup
	w := func() *Config {
		defer wg.Done()
		return getConfig()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		fmt.Printf("Worker %d calling getConfig()\n", i)
		go w()
	}
	wg.Wait()
}
