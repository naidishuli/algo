package sort

type AInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}