package main
import "fmt"
//1
//func main() {
//	fmt.Println("hello world")
//}

//2
//var x = []int{2: 2, 3, 0: 1}
//var y = []int{1,2,3,4}
//func main() {
//    fmt.Println(x)
//    fmt.Println(y)
//}

//3
//func incr(p *int) int {
//    *p++
//    return *p
//}
//func main() {
//    v := 1
//    incr(&v)
//    fmt.Println(v)
//}


//type ConfigOne struct {
//    Daemon string
//}
//
//func (c *ConfigOne) String() string {
//    return fmt.Sprintf("print: %v", c)
//}
//
//func main() {
//    c := &ConfigOne{}
//    c.String()
//}

func main() {
    var a = []int{1, 2, 3, 4, 5}
    var r = make([]int, 0)

    for i, v := range a {
        if i == 0 {
			fmt.Println(a)
        }

        r = append(r, v)
    }

    fmt.Println(r)
}
