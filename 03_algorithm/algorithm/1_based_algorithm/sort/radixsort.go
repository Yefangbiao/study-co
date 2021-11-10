package sort

func RadixSort(nums []int) []int {
	//分为正数和负数,对于负数,变成正数最后reverse
	negative := make([]int, 0)
	nonNegative := make([]int, 0)
	for _, num := range nums {
		if num < 0 {
			negative = append(negative, -num)
		} else {
			nonNegative = append(nonNegative, num)
		}
	}
	negative = radixSort(negative)
	i, j := 0, len(negative)-1
	for i <= j {
		negative[i], negative[j] = -negative[j], -negative[i]
		i++
		j--
	}
	negative = append(negative, radixSort(nonNegative)...)
	return negative
}

func radixSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	maxNum := maxElem(nums)
	radix := 1
	for ; maxNum/radix > 0; radix *= 10 {
		nums = countSort(nums, radix)
	}
	return nums
}

func maxElem(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// 具体的排序使用计数排序
func countSort(nums []int, radix int) []int {
	count := make([]int, 10)
	for _, num := range nums {
		count[(num/radix)%10]++
	}
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	ans := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		ans[count[(num/radix)%10]-1] = num
		count[(num/radix)%10]--
	}

	return ans
}
