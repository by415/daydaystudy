1.下面代码有什么错误？

```golang
func main() {
    one := 0
    one := 1 
}
```

多次初始化，只能初始化一次，后面采用赋值即可

2.下面代码有什么问题？

```golang
func main() {
    x := []int{
        1,
        2
    }
    _ = x
}
```




3.下面代码输出什么？

```golang
func test(x byte)  {
    fmt.Println(x)
}

func main() {
    var a byte = 0x11 
    var b uint8 = a
    var c uint8 = a + b
    test(c)
}
```

输出34，0x11是16进制的，表示的整形的17
