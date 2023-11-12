package main

import (
	"fmt"

	"github.com/hpcloud/tail"
)

func main() {
	filePath := "test.log"

	t, err := tail.TailFile(filePath, tail.Config{Follow: true})
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
