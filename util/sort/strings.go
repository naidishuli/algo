package sort

type AString []string
func (ai AString) Len() int           { return len(ai) }
func (ai AString) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AString) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }