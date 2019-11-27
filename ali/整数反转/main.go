/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main
import "fmt"

const INT_MAX = int(^uint32(0) >> 1)
const INT_MIN = ^INT_MAX
func reverse(x int) int {
	var ret int = 0
    for x != 0 {
		var pop int = x%10
		x /= 10
		if (ret > (INT_MAX - pop) / 10) {
			return 0
		}
		if (ret < (INT_MIN - pop) / 10 ) {
			return 0
		}
		ret = ret * 10 + pop
	}
	return ret
}



func main() {
	var test1 int = 123
	var test2 int = -123
	var test3 int = 120
	var test4 int = -120
	var test5 int = 1534236469
	fmt.Println(reverse(test1))
	fmt.Println(reverse(test2))
	fmt.Println(reverse(test3))
	fmt.Println(reverse(test4))
	fmt.Println(reverse(test5))
	fmt.Println(INT_MAX)
	fmt.Println(INT_MIN)
}
