package sort

import "testing"

func testFramework(t *testing.T, sortFunc func([]int) []int) {
	for _, testCase := range sortTestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := sortFunc(testCase.Input)
			pos, sorted := compareSlices(got, testCase.Except)
			if !sorted {
				if pos == -1 {
					t.Errorf("test %s failed due to slice length changing", testCase.Name)
				}
				t.Errorf("test %s failed at index %d", testCase.Name, pos)
			}
		})
	}
}

func compareSlices(a []int, b []int) (int, bool) {
	if len(a) != len(b) {
		return -1, false
	}
	for pos := range a {
		if a[pos] != b[pos] {
			return pos, false
		}
	}
	return -1, true
}

func TestBubbleSort(t *testing.T) {
	testFramework(t, BubbleSort)
}

func TestInsertionSort(t *testing.T) {
	testFramework(t, InsertionSort)
}

func TestBucketSort(t *testing.T) {
	testFramework(t, BucketSort)
}

func TestCountingSort(t *testing.T) {
	testFramework(t, CountingSort)
}

func TestMergeSort(t *testing.T) {
	testFramework(t, MergeSort)
}

func TestRadixSort(t *testing.T) {
	testFramework(t, RadixSort)
}

// TestSelectionSort very slow
//func TestSelectionSort(t *testing.T) {
//	testFramework(t, SelectionSort)
//}

func TestShellSort(t *testing.T) {
	testFramework(t, ShellSort)
}

func TestHeapSort(t *testing.T) {
	testFramework(t, HeapSort)
}

func TestHeapSortGo(t *testing.T) {
	testFramework(t, HeapSortGo)
}

func TestQuickSort(t *testing.T) {
	testFramework(t, QuickSort)
}
