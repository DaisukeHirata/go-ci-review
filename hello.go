package main

import "fmt"

func main() {
    sayHello()
}

func sayHello() {
    fmt.Println("Hello world!!! go vet ok version")
//    return
//    fmt.Println("Unreachable line") // got vet error test
}
