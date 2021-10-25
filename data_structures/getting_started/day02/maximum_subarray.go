package main

import "fmt"

//示例 1：
//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

//示例 2:
//输入：nums = [1]
//输出：1

//示例 3:
//输入：nums = [0]
//输出：0

//示例 4：
//输入：nums = [-1]
//输出：-1

//示例 5：
//输入：nums = [-100000]
//输出：-100000

// 提示：
// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
//
// 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/maximum-subarray
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// maxSubArray xxxxx [-2,1,-3,4,-1,2,1,-5,4]
func maxSubArray(nums []int) int {
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		fmt.Println("开始循环", nums[i])
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] = nums[i-1] + nums[i]
			fmt.Println("计算后值", nums[i])
		} else {
			fmt.Println("跳出循环值", nums[i])
		}
		if nums[i] > sum {
			sum = nums[i]
		}
		fmt.Println("sum", sum)
	}
	return sum
}
