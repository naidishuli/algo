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
			if j >= 0 && data.LessD(c, jd){
				fmt.Println(c)
				data.Assign(j + 1, jd)
				j -= 1
			}else {
				break
			}
		}

		data.Assign(j + 1, c)
	}
}
