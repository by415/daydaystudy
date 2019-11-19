1.下面的代码输出什么？

```golang
type N int

func (n *N) test(){
    fmt.Println(*n)
}

func main()  {
    var n N = 10
    p := &n

    n++
    f1 := n.test

    n++
    f2 := p.test

    n++
    fmt.Println(n)

    f1()
    f2()
}
```

13
13
13


2.下面哪一行代码会 panic，请说明原因？

```golang
package main

func main() {
  var m map[int]bool // nil
  _ = m[123]
  var p *[5]string // nil
  for range p {
    _ = len(p)
  }
  var s []int // nil
  _ = s[:]
  s, s[0] = []int{1, 2}, 9
}
```

panic s, s[0] = []int{1, 2}, 9,s[0]中s为nil，s[0]应该报越界错误
