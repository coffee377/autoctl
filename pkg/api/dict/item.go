package dict

import "github.com/coffee377/autoctl/pkg/api/base"

type DictionaryItem[V any] interface {
	IValue[V]
	IText
	IDescription
	base.ISort
}
