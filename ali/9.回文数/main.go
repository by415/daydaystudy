/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
func isPalindrome(x int) bool {
	if (x < 0) {
		return false
	}
	var tmpx int = x
	var ret int = 0
	var cur int = 0
	for tmpx != 0 {
		cur = tmpx % 10
		ret = ret*10 + cur
		tmpx /= 10
	}
	if (ret == x) {
		return true
	}
	return false
}


func main() {
	a := 100
	b := 202
	c := -123
	d := 10
	fmt.Println(isPalindrome(a))
	fmt.Println(isPalindrome(b))
	fmt.Println(isPalindrome(c))
	fmt.Println(isPalindrome(d))
}
