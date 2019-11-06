1.下面的代码有什么问题？

```
func main() {
    const x = 123
    const y = 1.23
    fmt.Println(x)
}
```
常量的定义格式：const identifier [type] = value,在 Go 语言中，你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型,所以上面的程序可以编译通过；

常量的值必须是能够在编译时就能够确定的；你可以在其赋值表达式中涉及计算过程，但是所有用于计算的值必须在编译期间就能获得。自定义的函数不能用于常量赋值，但是内置函数可以。

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

输出：
unit16 120
string abc

常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
	


3.下面代码有什么问题？

```
func main() {  
    var x string = nil 

    if x == nil { 
        x = "default"
    }
}
```

不可以将 nil 分配给 string 类型的变量:cannot use nil as type string in assignment

任何类型在未初始化时都对应一个零值：布尔类型是false，整型是0，字符串是""，而指针、函数、interface、slice、channel和map的零值都是nil。

nil没有默认的类型，尽管它是多个类型的零值，必须显式或隐式指定每个nil用法的明确类型。
