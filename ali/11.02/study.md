1.下面的代码有什么问题？

```golang
import (  
    "fmt"
    "log"
    "time"
)
func main() {  
}
```

答案：编译报错，imported and not used

2.下面代码输出什么？

```golang
func main() {
    x := interface{}(nil)
    y := (*int)(nil)
    a := y == x
    b := y == nil
    _, c := x.(interface{})
    println(a, b, c)
}
```

答案：false true false


A. true true true
B. false true true
C. true true true
D. false true false

3.下面代码有几处错误的地方？请说明原因。

```golang
func main() {
    var s []int
    s = append(s,1)

    var m map[string]int
    m["one"] = 1 
}
```

答案：不能对一个空的map直接赋值，应该使用make  append
