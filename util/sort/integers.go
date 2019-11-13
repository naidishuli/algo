package sort

type AInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type AInt8 []int8
func (ai AInt8) Len() int           { return len(ai) }
func (ai AInt8) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AInt8) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }

type AInt16 []int8
func (ai AInt16) Len() int           { return len(ai) }
func (ai AInt16) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AInt16) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }

type AInt32 []int8
func (ai AInt32) Len() int           { return len(ai) }
func (ai AInt32) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AInt32) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }

func BubbleSort(data AInterface) {
	for i := 0; i < data.Len(); i++ {
		for j := 0; j < (data.Len() - i - 1); j++ {
			if data.Less(j, j + 1) {
				data.Swap(j, j + 1)
			}
		}
	}
}
