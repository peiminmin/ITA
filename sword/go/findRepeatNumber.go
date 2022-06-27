package _go

/*
03. 数组中重复的数字

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3


限制：

2 <= n <= 100000


解题思路：
1. 首先想到的方法是用map过滤 或者 创建一个长度为n的 bool 数组，遍历数组设置对应index为true，寻找重复的数字 ，时间复杂度O(n),空间复杂度O(n)
2. 原地交换:遍历数组把元素调整到和下标相同的位置，时间复杂度O(n),空间复杂度O(1)
3.


*/

func findRepeatNumber(nums []int) int {

	for i := 0; i < len(nums); {
		if nums[i] == i {
			i++ //下标对齐，才往后移动
			continue
		}
		if nums[nums[i]] != nums[i] {
			swap(nums, i, nums[i])
		} else {
			return nums[i]
		}
	}
	return -1
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp

}
