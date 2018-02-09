package main

import (
    "fmt"
    "time"
    "math/rand"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func main() {
    fmt.Println("Hello Міша")
    fmt.Println("The time is", time.Now())
    fmt.Println("My favorite number is", rand.Intn(10))
}
