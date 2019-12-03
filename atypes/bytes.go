package atypes


type AByte []byte
func (ai AByte) Len() int           { return len(ai) }
func (ai AByte) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AByte) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }
func (ai AByte) Get(i int) interface{} { return ai[i] }
func (ai AByte) LessD(i, j interface{}) bool     { return i.(int) < j.(int) }
func (ai AByte) Assign(i int, d interface{}){ai[i] = d.(byte)}
func (ai AByte) InsertionB (i, j int){c := append(ai[:j], ai[i]);c = append(c, ai[j:i]...);c = append(c, ai[i+1:]...);ai = c}