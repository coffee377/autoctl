package auth

type FunctionPoints []FunctionPoint

func (f FunctionPoints) Len() int           { return len(f) }
func (f FunctionPoints) Less(i, j int) bool { return f[i].GetSort() < f[j].GetSort() }
func (f FunctionPoints) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
