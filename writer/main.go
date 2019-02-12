package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file1, err := os.Create("/data1/random.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	file2, err := os.Create("/data2/random.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file1.Close()
	defer file2.Close()
	for {
		file1.WriteString("hello data 1\n")
		file2.WriteString("hello data 2\n")
		fmt.Println("written  hello in file")
		time.Sleep(5 * time.Second)
	}
}
