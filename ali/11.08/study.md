1.下面代码有什么问题？

```golang
type foo struct {
    bar int
}

func main() {
    var f foo
    f.bar, tmp := 1, 2
}
```

参考答案及解析：编译错误：

non-name f.bar on left side of :=

:= 操作符不能用于结构体字段赋值。

2.下面的代码输出什么？

```golang
func main() {  
    fmt.Println(~2) 
}
```

编译出错：

invalid character U+007E '~'

很多语言都是采用 ~ 作为按位取反运算符，Go 里面采用的是 ^ ,另外需要注意的是，如果作为二元运算符，^ 表示按位异或,重点介绍下这个操作符 &^，按位置零


3.下面代码输出什么？

```golang
func main() {
    var ch chan int
    select {
    case v, ok := <-ch:
        println(v, ok)
    default:
        println("default") 
    }
}
```

输出default，因为ch为nil，所有的读写都会阻塞

4.下面这段代码输出什么？

```golang
type People struct {
    name string `json:"name"`
}

func main() {
    js := `{
        "name":"seekload"
    }`
    var p People
    err := json.Unmarshal([]byte(js), &p)
    if err != nil {
        fmt.Println("err: ", err)
        return
    }
    fmt.Println(p)
}
```

输出 {}。结构体访问控制，因为 name 首字母是小写，导致其他包不能访问，所以输出为空结构体
