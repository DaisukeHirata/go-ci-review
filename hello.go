package main

import "fmt"

func main() {
    sayHello()
}

func sayHello() {
 fmt.Println("Hello world!!! delete blanks in the beginning of this line") // gofmt test
//    return
//    fmt.Println("Unreachable line") // got vet error test
}
