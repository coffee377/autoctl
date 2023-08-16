package authority

type FunctionPoints []FunctionPoint

func (f FunctionPoints) Len() int           { return len(f) }
func (f FunctionPoints) Less(i, j int) bool { return f[i].GetPosition() < f[j].GetPosition() }
func (f FunctionPoints) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
