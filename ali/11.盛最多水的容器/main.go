/*
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。



图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 


输入: [1,8,6,2,5,4,8,3,7]
输出: 49

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
/*
暴力法  时间复杂度为O(len(height)*len(height))
*/
func maxArea(height []int) int {
    var cur int = 0
    var max int = 0
    for i:=0;i<len(height);i++ {
        for j:=i+1;j<len(height);j++ {
            if (height[i] < height[j]) {
                cur = (j-i) * height[i]
            } else {
                cur = (j-i) * height[j]
            }
            if (cur > max) {
                max = cur
            }
        }
    }
    return max   
}

/*
面积最大    先去长为最长的，然后移动高较小的垂直线，时间复杂度为O(len(height))
*/
func maxAreaBest(height []int) int {
    cur,max := 0,0
    i,j:=0,len(height)-1
    for i<j  {
        if(height[i] < height[j]) {
            cur = (j-i)*height[i]
            i++
        } else {
            cur = (j-i)*height[j]
            j--
        }
        if (cur > max) {
            max = cur
        }
    }
    return max
}

func main(){
	var test1 = []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(test1))
	fmt.Println(maxAreaBest(test1))
}
