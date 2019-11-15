package sort

type AFloat32 []float32
func (ai AFloat32) Len() int           { return len(ai) }
func (ai AFloat32) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AFloat32) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }

type AFloat64 []float64
func (ai AFloat64) Len() int           { return len(ai) }
func (ai AFloat64) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AFloat64) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }