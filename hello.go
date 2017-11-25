package main

import "fmt"

func main() {
    sayHello()
}

func sayHello() {
    fmt.Println("Hello world!!! circle ci - build pull request setting - on")
//    return
//    fmt.Println("Unreachable line") // got vet error test
}
