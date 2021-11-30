package heap

// Heap 大根堆
type Heap struct {
	nums    []int
	heapNum int
}

func NewHeap(nums []int) *Heap {
	heap := &Heap{
		nums:    nums,
		heapNum: len(nums),
	}
	for i := heap.heapNum / 2; i >= 0; i-- {
		heap.down(i)
	}

	return heap
}

func (h *Heap) Pop() int {
	x := h.nums[0]

	h.nums[0], h.nums[h.heapNum-1] = h.nums[h.heapNum-1], h.nums[0]
	h.nums = h.nums[:len(h.nums)-1]
	h.heapNum--
	h.down(0)

	return x
}

func (h *Heap) Push(x int) {
	h.nums = append(h.nums, x)
	h.heapNum++

	h.up(h.heapNum - 1)

}

func (h *Heap) down(x int) {
	max := x
	l, r := 2*x+1, 2*x+2
	if l < h.heapNum && h.nums[l] > h.nums[max] {
		max = l
	}
	if r < h.heapNum && h.nums[r] > h.nums[max] {
		max = r
	}
	if max != x {
		h.nums[max], h.nums[x] = h.nums[x], h.nums[max]
		h.down(max)
	}
}

func (h *Heap) up(x int) {
	for x != 0 && h.nums[x] > h.nums[(x-1)/2] {
		h.nums[x], h.nums[(x-1)/2] = h.nums[(x-1)/2], h.nums[x]
		x = (x - 1) / 2
	}
}
