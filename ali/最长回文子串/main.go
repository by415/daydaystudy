package main
import "fmt"

// 返回最长回文子串的长度
func longestPalindrome(s string) string {
    if (len(s) == 1 || len(s) == 0) {
        return s
    }
    var max_len int = 1
	var min int = 0
	var max int = 0
    for i := 1; i < len(s); i++ {
        var min_index int = i
        var max_index int = i
        var cur_len int = 1
        for min_index > 0 {
            if (s[i] == s[min_index-1]) {
                if (min_index > 0) {
                    cur_len++
                    min_index--
                }
            } else {
                break
            }
        }
        for max_index < len(s)-1 {
            if (s[i] == s[max_index+1]) {
                if(max_index < len(s) -1) {
                    cur_len++
                    max_index++
                }
            } else {
                break
            }
        }
		var ret_min int = min_index
		var ret_max int = max_index
        for j,k := min_index-1,max_index+1; j >=0 && k < len(s); j,k=j-1,k+1 {
            if (s[j] == s[k]) {
				ret_min = j
				ret_max = k
                cur_len += 2
            } else {
                break
            }
        }
        if (cur_len > max_len) {
            max_len = cur_len
			min = ret_min
			max = ret_max
        }
    }
	fmt.Println(min,max)
    return s[min:max+1]
}



func main() {
	var test1 string = "aa"
	var test2 string = "bcabbbbbbbbbbbacb"
	var test3 string = "abcccc"
	var test4 string = "addbc"
	fmt.Println(longestPalindrome(test1))
	fmt.Println(longestPalindrome(test2))
	fmt.Println(longestPalindrome(test3))
	fmt.Println(longestPalindrome(test4))
}
