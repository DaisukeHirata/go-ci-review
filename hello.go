package main

import "fmt"

func main() {
	SayHello2()
}

func SayHello2() {
	fmt.Println("Hello world!!") // gofmt test
	//    return
	//    fmt.Println("Unreachable line") // got vet error test
}
