package main

//1
//import (
//    "fmt"
//    "log"
//    "time"
//)
//func main() {
//}

//2
func main() {
    x := interface{}(nil)
    y := (*int)(nil)
    a := y == x
    b := y == nil
    _, c := x.(interface{})
    println(a, b, c)
}
