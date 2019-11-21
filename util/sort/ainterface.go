package sort

type AInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Assign(i int, d interface{})
	Get(i int) interface{}
	LessD(i, j interface{}) bool
	InsertionB(i, j int)
}
