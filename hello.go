package main

import "fmt"

func main() {
    SayHello()
}

func SayHello() {
 fmt.Println("Hello world!!") // gofmt test
//    return
//    fmt.Println("Unreachable line") // got vet error test
}
