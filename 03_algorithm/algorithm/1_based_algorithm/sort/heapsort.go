package sort

func HeapSort(nums []int) []int {
	h := newMaxHeap(nums)
	ans := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] = h.pop()
	}
	return ans
}

type maxHeap struct {
	slice     []int
	heapSlice int
}

func newMaxHeap(slice []int) *maxHeap {
	heap := &maxHeap{
		slice:     slice,
		heapSlice: len(slice),
	}
	for i := heap.heapSlice / 2; i >= 0; i-- {
		heap.heapify(i)
	}
	return heap
}

func (h *maxHeap) heapify(i int) {
	l, r := 2*i+1, 2*i+2
	max := i
	if l < h.heapSlice && h.slice[l] > h.slice[max] {
		max = l
	}
	if r < h.heapSlice && h.slice[r] > h.slice[max] {
		max = r
	}
	if max != i {
		h.slice[i], h.slice[max] = h.slice[max], h.slice[i]
		h.heapify(max)
	}
}

func (h *maxHeap) pop() int {
	x := h.slice[0]
	h.slice[0], h.slice[h.heapSlice-1] = h.slice[h.heapSlice-1], h.slice[0]
	h.heapSlice--
	h.heapify(0)
	return x
}
