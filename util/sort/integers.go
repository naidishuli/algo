package sort

type AInt []int

func (ai AInt) Len() int              { return len(ai) }
func (ai AInt) Less(i, j int) bool    { return ai[i] < ai[j] }
func (ai AInt) Swap(i, j int)         { ai[i], ai[j] = ai[j], ai[i] }
func (ai AInt) Get(i int) interface{} { return ai[i] }
func (ai AInt) LessD(i, j interface{}) bool     { return i.(int) < j.(int) }
func (ai AInt) Assign(i int, d interface{}){
	ai[i] = d.(int)
}
func (ai AInt) InsertionB (i, j int){
	c := append(ai[:j], ai[i])
	c = append(c, ai[j:i]...)
	c = append(c, ai[i+1:]...)
	ai = c
}


type AInt8 []int8

func (ai AInt8) Len() int           { return len(ai) }
func (ai AInt8) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AInt8) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }

type AInt16 []int16

func (ai AInt16) Len() int           { return len(ai) }
func (ai AInt16) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AInt16) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }

type AInt32 []int32

func (ai AInt32) Len() int           { return len(ai) }
func (ai AInt32) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AInt32) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }

type AInt64 []int64

func (ai AInt64) Len() int           { return len(ai) }
func (ai AInt64) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AInt64) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }

type AUInt []uint

func (ai AUInt) Len() int           { return len(ai) }
func (ai AUInt) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AUInt) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }

type AUInt8 []uint8

func (ai AUInt8) Len() int           { return len(ai) }
func (ai AUInt8) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AUInt8) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }

type AUInt16 []uint16

func (ai AUInt16) Len() int           { return len(ai) }
func (ai AUInt16) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AUInt16) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }

type AUInt32 []uint32

func (ai AUInt32) Len() int           { return len(ai) }
func (ai AUInt32) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AUInt32) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }

type AUInt64 []uint64

func (ai AUInt64) Len() int           { return len(ai) }
func (ai AUInt64) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AUInt64) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }
