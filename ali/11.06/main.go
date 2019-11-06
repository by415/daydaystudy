package main

import "fmt"
//
//func main() {
//    data := []int{1,2,3}
//    i := 0
//    ++i
//    fmt.Println(data[i++])
//}

func main() {
    x := 1
    fmt.Println(x)
    {
        fmt.Println(x)
        i,x := 2,2
        fmt.Println(i,x)
    }
    fmt.Println(x)  // print ?
}
