package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	for {
		data1, err := ioutil.ReadFile("/data1/random.txt")
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(string(data1))
		data2, err := ioutil.ReadFile("/data2/random.txt")
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(string(data2))
		time.Sleep(1 * time.Second)
	}
}
