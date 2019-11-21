package main
import "fmt"
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
func twoSum(nums []int, target int) []int {
	var ret []int = make([]int, 2)
    //var ret = []int{0,0}
	for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums); j++ {
            if(i == j) {
                continue
            }
            if(nums[i]+nums[j] == target) {
				ret[0]=i
				ret[1]=j
                return ret
            }
        }
    }
	return ret
}

func main() {
	var p *[]int = new([]int)
	if (*p == nil) {
		fmt.Println("aaaaaaaaa")
	}
	fmt.Println(p,*p)
	var nums = []int{2,7,11,13}
	var target int = 20
	fmt.Println(twoSum(nums,target))
}
//package main

//func main() {
//  var m map[int]bool // nil
//  _ = m[123]
//  var p *[5]string // nil
//  for range p {
//    _ = len(p)
//  }
//  var s []int // nil
//  _ = s[:]
//  s = []int{1,2}
//  s[0] = 9
//}
