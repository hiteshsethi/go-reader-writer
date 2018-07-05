package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	for {
		data, err := ioutil.ReadFile("/data/random.txt")
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(string(data))
		time.Sleep(1 * time.Second)
	}
}
