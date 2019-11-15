package sort


type AByte []byte
func (ai AByte) Len() int           { return len(ai) }
func (ai AByte) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AByte) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }