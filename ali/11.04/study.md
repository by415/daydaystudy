1.下面代码有什么错误？

```golang
func main() {
    one := 0
    one := 1 
}
```

多次声明一个变量，不能再单独的声明中重复声明一个变量，但是在多变量声明的时候可以重复声明，但是必须保证至少有一个变量是新声明的。
eg
```
func main() {
	one := 1
	one, two := 2,3
}
```

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
编译错误，第四行代码没有逗号。用字面量初始化数组、slice 和 map 时，最好是在每个元素后面加上逗号，即使是声明在一行或者多行都不会出错。



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
