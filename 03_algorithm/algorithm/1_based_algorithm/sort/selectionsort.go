package sort

func SelectionSort(nums []int) []int {
	n := len(nums)
	var minIndex int

	for i := 0; i < n; i++ {
		minIndex = i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}

	return nums
}
