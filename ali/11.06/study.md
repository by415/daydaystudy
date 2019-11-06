1.下面的代码有什么问题？

```
func main() {
    data := []int{1,2,3}
    i := 0
    ++i
    fmt.Println(data[i++])
}
```

编译出错，go没有前置++

2.下面代码最后一行输出什么？请说明原因。

```
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
```

输出
1
1
2 2
1 
{}作用域
