package sort

import "math"

func CountingSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	// 统计最大值和最小值
	min, max := math.MaxInt32, math.MinInt32
	for _, num := range nums {
		min = minInt(min, num)
		max = maxInt(max, num)
	}

	// 统计变形
	count := make([]int, max-min+1)
	for _, num := range nums {
		count[num-min]++
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	// 填充结果数组
	res := make([]int, n)
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		res[count[num-min]-1] = num
		count[num-min]--
	}
	return res
}
