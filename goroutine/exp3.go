package main

import (
	"fmt"
	"time"
)

func main(){
	go g1()
	go g2()
	time.Sleep(time.Second)
}

func g1(){
	for i:='A'; i<='Z';i++ {
		go fmt.Printf("%c-", i)
	}
}

func g2(){
	for i:='a';i<='z';i++ {
		go fmt.Printf("%c-", i)
	}
}
