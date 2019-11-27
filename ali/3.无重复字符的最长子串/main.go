package main
import "fmt"
/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func lengthOfLongestSubstring(s string) int {
    if (len(s) == 0) {
        return 0
    }
    var ret int = 1
    for i := 0; i < len(s); i++ {
        var sub_len int = 1
        var sign bool = true
        for j := i + 1; j < len(s); j++ {
            for k := i; k < j; k++ {
                if (s[j] == s[k]) {
                    sign = false
                }
            }
            if (sign == true) {
                sub_len += 1
            } else{
                sign = true
                break
            }
        }
        if (sub_len > ret) {
            ret = sub_len
        }
    }
    return ret
}



func main() {
	var s string = "abcabcabc"
	var a string = " "
	var b string = ""
	var c string = "au"
	fmt.Println(s,lengthOfLongestSubstring(s))
	fmt.Println(a,lengthOfLongestSubstring(a))
	fmt.Println(b,lengthOfLongestSubstring(b))
	fmt.Println(c,lengthOfLongestSubstring(c))
}
