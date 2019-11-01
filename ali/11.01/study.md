### 1.下面代码编译能通过吗？

```
func main()  
{ 
    fmt.Println("hello world")
}
```

答案解析：编译错误，go中大括号不能放在单独的一行

```
syntax error: unexpected semicolon or newline before {
```

应该为
```
func main() { 
    fmt.Println("hello world")
}
```


### 2.下面这段代码输出什么？

```
var x = []int{2: 2, 3, 0: 1}

func main() {
    fmt.Println(x)
}
```

答案解析：输出[1 0 2 3],切片可以指定索引初始化，没有指定的所以默认值为0，所以第一个元素为1，第二个没有指定则为默认值0，第三个指定为2，未指定索引的索引依次加1。

### 3.下面这段代码输出什么？

```
func incr(p *int) int {
    *p++
    return *p
}
func main() {
    v := 1
    incr(&v)
    fmt.Println(v)
}
```

答案解析：输出2，p 是指针变量，指向变量 v，*p++操作的意思是取出变量v的值并执行加一操作，所以v的最终值是 2。go语言中++,--只能作为语句不能作为表达式，也就是a=n++是不允许的，a++和a--表示a += 1;a -= 1--->a = a+1;a = a - 1。
