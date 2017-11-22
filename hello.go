package main

import "fmt"

func main() {
    sayHello()
}

func sayHello() {
    fmt.Println("Hello world!!")
    return
    fmt.Println("Unreachable line") // got vet error test
}
