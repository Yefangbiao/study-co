package sort

import "container/heap"

type HeapNums []int

func (h HeapNums) Len() int {
	return len(h)
}

func (h HeapNums) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h HeapNums) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapNums) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *HeapNums) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func HeapSortGo(nums []int) []int {
	h := &HeapNums{}
	*h = append(*h, nums...)
	heap.Init(h)

	ans := make([]int, 0, len(nums))
	for h.Len() > 0 {
		ans = append(ans, heap.Pop(h).(int))
	}

	return ans
}
