package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hpcloud/tail"
)

func main() {
	mongodbURI, exists := os.LookupEnv("MONGODB_CONNECTION_STRING")
	if !exists {
		panic("Env variable MONGODB_CONNECTION_STRING needs to be set")
	}

	filePath := "/sshesame.log"

	t, err := tail.TailFile(filePath, tail.Config{Follow: true})
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	repo := CreateRepository(mongodbURI)

	for line := range t.Lines {
		var logLine LogLine

		err := json.Unmarshal([]byte(line.Text), &logLine)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		logLine.ExtractIpAndPort()

		repo.Insert(logLine)

	}
}
