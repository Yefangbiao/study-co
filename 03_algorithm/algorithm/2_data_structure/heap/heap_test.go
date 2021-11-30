package heap

import (
	"testing"
)

func TestHeap(t *testing.T) {
	nums := []int{5, 7, 2, 9, 1}
	want := []int{9, 7, 5, 2, 1}

	h := NewHeap(nums)
	for i := 0; i < len(nums); i++ {
		x := h.Pop()
		if x != want[i] {
			t.Errorf("got: %v,want: %v", x, want[i])
		}
	}

	h = NewHeap(nums)
	h.Push(11)
	h.Push(8)
	x := h.Pop()
	if x != 11 {
		t.Errorf("got: %v,want: %v", x, 11)
	}
	x = h.Pop()
	if x != 9 {
		t.Errorf("got: %v,want: %v", x, 9)
	}

}
