/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
func getCommonStr(str1 string,str2 string) string{
    var lens int = 0
    var ret string = ""
    if(len(str1)>len(str2)){
        lens=len(str2)
    }else {
        lens=len(str1)
    }
    for i:=0;i<lens;i++{
        if(str1[i] != str2[i]){
            break
        }
        ret += string(str1[i])
    }
    return ret
}
func longestCommonPrefix(strs []string) string {
    if(len(strs)==0){
        return ""
    }
    var first string = strs[0]
    var ret string = first
    for i:=1;i<len(strs);i++{
        cur:=getCommonStr(first,strs[i])
        if(len(ret)>len(cur)){
            ret=cur
        }
    }
    return ret
}

func main()
{
    var strs = []string{"fl","flaaa","flwe"}
    fmt.Println(longestCommonPrefix(strs))
}
