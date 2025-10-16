package dict

import "github.com/coffee377/autoctl/pkg/api/base"

type DictionaryItem[V string | int | bool] interface {
	IValue[V]
	IText
	IDescription
	base.ISort
}
