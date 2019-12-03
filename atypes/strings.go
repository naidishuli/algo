package atypes

type AString []string
func (ai AString) Len() int           { return len(ai) }
func (ai AString) Less(i, j int) bool { return ai[i] < ai[j] }
func (ai AString) Swap(i, j int)      { ai[i], ai[j] = ai[j], ai[i] }
func (ai AString) Get(i int) interface{} { return ai[i] }
func (ai AString) LessD(i, j interface{}) bool     { return i.(int) < j.(int) }
func (ai AString) Assign(i int, d interface{}){ai[i] = d.(string)}
func (ai AString) InsertionB (i, j int){c := append(ai[:j], ai[i]);c = append(c, ai[j:i]...);c = append(c, ai[i+1:]...);ai = c}