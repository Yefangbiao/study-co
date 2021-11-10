package sort

import "math/rand"

func QuickSort(nums []int) []int {
	return quickSort1(nums)

	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, low, high int) {
	if low >= high {
		return
	}

	pivot := partition(nums, low, high)
	quickSort(nums, low, pivot-1)
	quickSort(nums, pivot+1, high)
}

func partition(nums []int, low, high int) int {
	// 这里可以选择左边，右边，中间，随机
	pivot := rand.Intn(high-low+1) + low
	nums[pivot], nums[high] = nums[high], nums[pivot]
	pivotNum := nums[high]

	i := low - 1
	for j := low; j < high; j++ {
		if nums[j] <= pivotNum {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]

	return i + 1
}

func quickSort1(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// 这里可以选择左边，右边，中间，随机
	pivot := rand.Intn(len(nums))
	pivotNum := nums[pivot]

	lowPart := make([]int, 0)
	middlePart := make([]int, 0)
	highPart := make([]int, 0)

	for _, num := range nums {
		if num < pivotNum {
			lowPart = append(lowPart, num)
		}
		if num == pivotNum {
			middlePart = append(middlePart, num)
		}
		if num > pivotNum {
			highPart = append(highPart, num)
		}
	}

	lowPart = quickSort1(lowPart)
	highPart = quickSort1(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}
