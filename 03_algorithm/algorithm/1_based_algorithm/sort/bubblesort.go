package sort

func BubbleSort(nums []int) []int {
	flag := true
	n := len(nums)

	for i := 0; i < n-1; i++ {
		flag = true
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return nums
}
