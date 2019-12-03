package sort

import "github.com/naidishuli/algo/atypes"

//by default the data is sorted in ascending mode
func BubbleSort(data atypes.AInterface) {
	l := data.Len()
	for i := 0; i < l; i++ {
		for j := 0; j < (l - i - 1); j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
			}
		}
	}
}

//by default the data is sorted in ascending mode
func SelectionSort(data atypes.AInterface) {
	l := data.Len()
	for i := 0; i < l; i++ {
		minIndex := i
		for j := i + 1; j < l; j++ {
			if data.Less(j, minIndex) {
				minIndex = j
			}
		}
		data.Swap(i, minIndex)
	}
}

//by default the data is sorted in ascending mode
func InsertionSortBinary(data atypes.AInterface) {
	for i := 1; i < data.Len(); i++{
		//val := data.Get(i)
		j := insertionBinary(data, i, 0, i-1)
		//arr = arr[:j] + [val] + arr[j:i] + arr[i+1:]
		data.InsertionB(i, j)
	}
}

func insertionBinary(data atypes.AInterface, i int, start, end int) int {
	if start == end{
		if data.Less(i, start){
			return start
		}else{
			return start + 1
		}
	}

	if start > end{
		return start
	}

	mid := (start + end) / 2

	if data.Less(mid, i){
		return insertionBinary(data, i, mid + 1, end)
	}

	if data.Less(i, mid){
		return insertionBinary(data, i, start, mid - 1)
	}

	return mid
}

//by default the data is sorted in ascending mode
func InsertionSort(data atypes.AInterface) {
	var n = data.Len()
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if data.Less(j, j-1){
				data.Swap(j-1,j)
			}
			j = j - 1
		}
	}
}

//by default the data is sorted in ascending mode
func QuickSort(data atypes.AInterface) {
	// TODO: explore different ways to choose low and high
	// index and measure performance
	quickSort(data, 0, data.Len()-1)
}

func partition(data atypes.AInterface, low int, high int) int {
	pivot := data.Get(high)

	i := (low - 1)

	for j := low; j <= high-1; j++ {
		if data.LessD(data.Get(j), pivot) {
			i++
			data.Swap(i, j)
		}
	}
	data.Swap(i+1, high)
	return i + 1
}

func quickSort(data atypes.AInterface, low int, high int) {
	if low < high {
		pi := partition(data, low, high)

		// TODO: this can be parallelized
		quickSort(data, low, pi-1)  // Before pi
		quickSort(data, pi+1, high) // After pi
	}
}