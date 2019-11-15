package sort

func BubbleSort(data AInterface) {
	for i := 0; i < data.Len(); i++ {
		for j := 0; j < (data.Len() - i - 1); j++ {
			if data.Less(j, j + 1) {
				data.Swap(j, j + 1)
			}
		}
	}
}
