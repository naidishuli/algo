package sort

import "fmt"

//by default the data is sorted in ascending mode
func BubbleSort(data AInterface) {
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
func SelectionSort(data AInterface) {
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
func InsertionSort(data AInterface) {
	l := data.Len()
	for i := 1; i < l; i++ {
		c := data.Get(i)

		j := i - 1

		jd := data.Get(j)

		for {
			if j >= 0 && data.LessD(c, jd) {
				fmt.Println(c)
				data.Assign(j+1, jd)
				j -= 1
			} else {
				break
			}
		}

		data.Assign(j+1, c)
	}
}

func partition(data AInterface, low int, high int) int {
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

func quickSort(data AInterface, low int, high int) {
	if low < high {
		pi := partition(data, low, high)

		// TODO: this can be parallelized
		quickSort(data, low, pi-1)  // Before pi
		quickSort(data, pi+1, high) // After pi
	}
}

//by default the data is sorted in ascending mode
func QuickSort(data AInterface) {
	// TODO: explore different ways to choose low and high
	// index and measure performance
	quickSort(data, 0, data.Len()-1)
}
