package main

import "fmt"

func main() {
    SayHello()
}

func SayHello() {
 fmt.Println("Hello world!!! delete blanks in the beginning of this line!!") // gofmt test
//    return
//    fmt.Println("Unreachable line") // got vet error test
}
