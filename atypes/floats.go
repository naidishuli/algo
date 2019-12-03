package atypes

type AFloat32 []float32
func (ai AFloat32) Len() int           { return len(ai) }
func (ai AFloat32) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AFloat32) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }
func (ai AFloat32) Get(i int) interface{} { return ai[i] }
func (ai AFloat32) LessD(i, j interface{}) bool     { return i.(int) < j.(int) }
func (ai AFloat32) Assign(i int, d interface{}){ai[i] = d.(float32)}
func (ai AFloat32) InsertionB (i, j int){c := append(ai[:j], ai[i]);c = append(c, ai[j:i]...);c = append(c, ai[i+1:]...);ai = c}

type AFloat64 []float64
func (ai AFloat64) Len() int           { return len(ai) }
func (ai AFloat64) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AFloat64) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }
func (ai AFloat64) Get(i int) interface{} { return ai[i] }
func (ai AFloat64) LessD(i, j interface{}) bool     { return i.(int) < j.(int) }
func (ai AFloat64) Assign(i int, d interface{}){ai[i] = d.(float64)}
func (ai AFloat64) InsertionB (i, j int){c := append(ai[:j], ai[i]);c = append(c, ai[j:i]...);c = append(c, ai[i+1:]...);ai = c}