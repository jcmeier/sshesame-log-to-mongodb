package main

import (
	"encoding/json"
	"fmt"

	"github.com/hpcloud/tail"
)

func main() {
	filePath := "sshesame.log"

	t, err := tail.TailFile(filePath, tail.Config{Follow: true})
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	repo := CreateRepository()

	for line := range t.Lines {
		var logLine LogLine

		err := json.Unmarshal([]byte(line.Text), &logLine)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		repo.Insert(logLine)

	}
}
