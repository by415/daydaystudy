package main
//import "fmt"
//type N int
//
//func (n *N) test(){
//    fmt.Println(*n)
//}
//
//func main()  {
//    var n N = 10
//    p := &n
//
//    n++
//    f1 := n.test
//
//    n++
//    f2 := p.test
//
//    n++
//    fmt.Println(n)
//
//    f1()
//    f2()
//}

//package main

func main() {
  var m map[int]bool // nil
  _ = m[123]
  var p *[5]string // nil
  for range p {
    _ = len(p)
  }
  var s []int // nil
  _ = s[:]
  s = []int{1,2}
  s[0] = 9
}
