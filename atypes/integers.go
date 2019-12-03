package atypes

type AInt []int
func (ai AInt) Len() int              { return len(ai) }
func (ai AInt) Less(i, j int) bool    { return ai[i] < ai[j] }
func (ai AInt) Swap(i, j int)         { ai[i], ai[j] = ai[j], ai[i] }
func (ai AInt) Get(i int) interface{} { return ai[i] }
func (ai AInt) LessD(i, j interface{}) bool     { return i.(int) < j.(int) }
func (ai AInt) Assign(i int, d interface{}){ai[i] = d.(int)}
func (ai AInt) InsertionB (i, j int){c := append(ai[:j], ai[i]);c = append(c, ai[j:i]...);c = append(c, ai[i+1:]...);ai = c}

type AInt8 []int8
func (ai AInt8) Len() int           { return len(ai) }
func (ai AInt8) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AInt8) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }
func (ai AInt8) Get(i int) interface{} { return ai[i] }
func (ai AInt8) LessD(i, j interface{}) bool     { return i.(int) < j.(int) }
func (ai AInt8) Assign(i int, d interface{}){ai[i] = d.(int8)}
func (ai AInt8) InsertionB (i, j int){c := append(ai[:j], ai[i]);c = append(c, ai[j:i]...);c = append(c, ai[i+1:]...);ai = c}

type AInt16 []int16
func (ai AInt16) Len() int           { return len(ai) }
func (ai AInt16) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AInt16) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }
func (ai AInt16) Get(i int) interface{} { return ai[i] }
func (ai AInt16) LessD(i, j interface{}) bool     { return i.(int) < j.(int) }
func (ai AInt16) Assign(i int, d interface{}){ai[i] = d.(int16)}
func (ai AInt16) InsertionB (i, j int){c := append(ai[:j], ai[i]);c = append(c, ai[j:i]...);c = append(c, ai[i+1:]...);ai = c}

type AInt32 []int32
func (ai AInt32) Len() int           { return len(ai) }
func (ai AInt32) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AInt32) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }
func (ai AInt32) Get(i int) interface{} { return ai[i] }
func (ai AInt32) LessD(i, j interface{}) bool     { return i.(int) < j.(int) }
func (ai AInt32) Assign(i int, d interface{}){ai[i] = d.(int32)}
func (ai AInt32) InsertionB (i, j int){c := append(ai[:j], ai[i]);c = append(c, ai[j:i]...);c = append(c, ai[i+1:]...);ai = c}

type AInt64 []int64
func (ai AInt64) Len() int           { return len(ai) }
func (ai AInt64) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AInt64) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }
func (ai AInt64) Get(i int) interface{} { return ai[i] }
func (ai AInt64) LessD(i, j interface{}) bool     { return i.(int) < j.(int) }
func (ai AInt64) Assign(i int, d interface{}){ai[i] = d.(int64)}
func (ai AInt64) InsertionB (i, j int){c := append(ai[:j], ai[i]);c = append(c, ai[j:i]...);c = append(c, ai[i+1:]...);ai = c}

type AUInt []uint
func (ai AUInt) Len() int           { return len(ai) }
func (ai AUInt) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AUInt) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }
func (ai AUInt) Get(i int) interface{} { return ai[i] }
func (ai AUInt) LessD(i, j interface{}) bool     { return i.(int) < j.(int) }
func (ai AUInt) Assign(i int, d interface{}){ai[i] = d.(uint)}
func (ai AUInt) InsertionB (i, j int){c := append(ai[:j], ai[i]);c = append(c, ai[j:i]...);c = append(c, ai[i+1:]...);ai = c}

type AUInt8 []uint8
func (ai AUInt8) Len() int           { return len(ai) }
func (ai AUInt8) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AUInt8) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }
func (ai AUInt8) Get(i int) interface{} { return ai[i] }
func (ai AUInt8) LessD(i, j interface{}) bool     { return i.(int) < j.(int) }
func (ai AUInt8) Assign(i int, d interface{}){ai[i] = d.(uint8)}
func (ai AUInt8) InsertionB (i, j int){c := append(ai[:j], ai[i]);c = append(c, ai[j:i]...);c = append(c, ai[i+1:]...);ai = c}

type AUInt16 []uint16
func (ai AUInt16) Len() int           { return len(ai) }
func (ai AUInt16) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AUInt16) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }
func (ai AUInt16) Get(i int) interface{} { return ai[i] }
func (ai AUInt16) LessD(i, j interface{}) bool     { return i.(int) < j.(int) }
func (ai AUInt16) Assign(i int, d interface{}){ai[i] = d.(uint16)}
func (ai AUInt16) InsertionB (i, j int){c := append(ai[:j], ai[i]);c = append(c, ai[j:i]...);c = append(c, ai[i+1:]...);ai = c}

type AUInt32 []uint32
func (ai AUInt32) Len() int           { return len(ai) }
func (ai AUInt32) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AUInt32) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }
func (ai AUInt32) Get(i int) interface{} { return ai[i] }
func (ai AUInt32) LessD(i, j interface{}) bool     { return i.(int) < j.(int) }
func (ai AUInt32) Assign(i int, d interface{}){ai[i] = d.(uint32)}
func (ai AUInt32) InsertionB (i, j int){c := append(ai[:j], ai[i]);c = append(c, ai[j:i]...);c = append(c, ai[i+1:]...);ai = c}

type AUInt64 []uint64
func (ai AUInt64) Len() int           { return len(ai) }
func (ai AUInt64) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AUInt64) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }
func (ai AUInt64) Get(i int) interface{} { return ai[i] }
func (ai AUInt64) LessD(i, j interface{}) bool     { return i.(int) < j.(int) }
func (ai AUInt64) Assign(i int, d interface{}){ai[i] = d.(uint64)}
func (ai AUInt64) InsertionB (i, j int){c := append(ai[:j], ai[i]);c = append(c, ai[j:i]...);c = append(c, ai[i+1:]...);ai = c}