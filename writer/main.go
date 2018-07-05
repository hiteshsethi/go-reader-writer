package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("/data/random.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file.Close()
	for {
		file.WriteString("hello\n")
		fmt.Println("written  hello in file")
		time.Sleep(5 * time.Second)
	}
}
