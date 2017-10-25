package main

import (
	"fmt"
	"time"
)

func heartBeat(){
	for range time.Tick(time.Second *1){
		fmt.Println("Foo")
	}
}

func main() {
	go heartBeat()
}