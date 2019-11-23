/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var two_len int = len(nums1)+ len(nums2)
	ret := make([]int,two_len)
	i := 0
	j := 0
	k := 0
	for {
		if (i == len(nums1) && j < len(nums2)) {
			ret[k] = nums2[j]
			k++
			j++
		}else if (j == len(nums2) && i < len(nums1)) {
			ret[k] = nums1[i]
			k++
			i++
		} else if (i == len(nums1) && j == len(nums2)) {
			break
		} else {
			if (nums1[i]<nums2[j]) {
				ret[k] = nums1[i]
				i++
				k++
			} else {
				ret[k] = nums2[j]
				j++
				k++
			}
		}
	}
	if(two_len%2 == 0) {
		return float64(ret[two_len/2-1]+ret[two_len/2])/2
	}else {
		return float64(ret[(two_len+1)/2-1])
	}
}



func main(){
	var nums1 = []int{1,2,3}
	var nums2 = []int{4,5}
	fmt.Println(findMedianSortedArrays(nums1,nums2))
	var nums3 = []int{1,3}
	var nums4 = []int{2}
	fmt.Println(findMedianSortedArrays(nums3,nums4))
}
