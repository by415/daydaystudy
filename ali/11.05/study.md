1.下面的代码有什么问题？

```
func main() {
    const x = 123
    const y = 1.23
    fmt.Println(x)
}
```

2.下面代码输出什么？

```
const (
    x uint16 = 120
    y
    s = "abc"
    z
)

func main() {
    fmt.Printf("%T %v\n", y, y)
    fmt.Printf("%T %v\n", z, z)
}
```

3.下面代码有什么问题？

```
func main() {  
    var x string = nil 

    if x == nil { 
        x = "default"
    }
}
```
