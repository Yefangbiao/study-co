package sort

func ShellSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	for d := len(nums) / 2; d > 0; d /= 2 {
		for i := d; i < len(nums); i++ {
			for j := i; j >= d && nums[j-d] > nums[j]; j -= d {
				nums[j], nums[j-d] = nums[j-d], nums[j]
			}
		}
	}
	return nums
}
