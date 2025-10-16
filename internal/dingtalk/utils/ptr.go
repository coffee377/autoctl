package utils

// ToPtr 将任意类型的值转换为其指针类型
// 适用于所有可寻址的类型，包括基本类型、结构体、切片元素等
func ToPtr[T any](v T) *T {
	return &v
}

// FromPtr 将指针类型转换为其对应的值类型
// 如果指针为nil，则返回该类型的零值
func FromPtr[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}
