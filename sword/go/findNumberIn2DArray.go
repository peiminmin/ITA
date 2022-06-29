package _go

/*
04. 二维数组中的查找
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

示例:
现有矩阵 matrix 如下：
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target=5，返回true。
给定target=20，返回false。
限制：
0 <= n <= 1000
0 <= m <= 1000

解题思路： 从右上角开始，比target大，往前走，比target小，往下走
*/

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}
	var endX, endY int
	endX = len(matrix) - 1
	if endX >= 0 {
		endY = len(matrix[0]) - 1
		if endY < 0 {
			return false
		}
	} else {
		return false
	}

	curX, curY := 0, endY
	for curX <= endX && curY >= 0 {
		if matrix[curX][curY] > target {
			curY--
		} else if matrix[curX][curY] < target {
			curX++
		} else {
			return true
		}
	}
	return false
}

func findNumber(array []int, target int) bool {
	start, end := 0, len(array)-1
	for start < end-1 {
		if array[start] == target || array[end] == target {
			return true
		}

		mid := (start + end) / 2
		if array[mid] > target {
			end = mid
		} else if array[mid] < target {
			start = mid
		} else {
			return true
		}
	}
	return false
}
