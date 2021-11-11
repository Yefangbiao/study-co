package sort

func BucketSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	min, max := nums[0], nums[0]
	for _, num := range nums {
		min = minInt(min, num)
		max = maxInt(max, num)
	}
	// 桶的间距。(防止出现[1,1])的情况
	d := maxInt(1, (max-min)/(len(nums)-1))
	// 桶的个数
	bucketSize := (max-min)/d + 1
	// 桶
	buckets := make([][]int, bucketSize)
	for _, num := range nums {
		buckets[(num-min)/d] = append(buckets[(num-min)/d], num)
	}

	// 每个桶排序
	for _, bucket := range buckets {
		// 这里可以用计数排序
		bucket = bucketCountingSort(bucket)
	}

	ans := make([]int, 0, len(nums))
	for _, bucket := range buckets {
		ans = append(ans, bucket...)
	}
	return ans
}

func bucketCountingSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	min, max := nums[0], nums[0]
	for _, num := range nums {
		min = minInt(min, num)
		max = maxInt(max, num)
	}

	count := make([]int, max-min+1)
	for _, num := range nums {
		count[num-min]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	ans := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		ans[count[nums[i]-min]-1] = nums[i]
		count[nums[i]-min]--
	}
	return ans
}
