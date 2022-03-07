package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go func() {
		for {
			select {
			case <- done:
				fmt.Println("exit go routineeeeee")
				return
			default:
				fmt.Println("monitoringgggggg...")
				time.Sleep(1 * time.Second)
			}
			fmt.Println("aaaaa")
		}
	}()
	time.Sleep(3 * time.Second)
	done <- true
	time.Sleep(5 * time.Second)
	fmt.Println("exit product")
	time.Sleep(5 * time.Second)
}