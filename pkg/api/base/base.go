package base

import "time"

type Identify[T any] interface {
	GetId() T

	SetId(id T)
}

type ISort interface {
	GetSort() int
}

type IFlag interface {
	GetFlag() int
}

type IStatus interface {
	GetStatus() int
}

type (
	ICreatedBy[ID any] interface {
		GetCreatedBy() ID
	}

	ICreatedAt interface {
		GetCreatedAt() *time.Time
	}
)

type (
	IUpdatedBy[ID any] interface {
		GetUpdatedBy() ID
	}

	IUpdatedAt interface {
		GetUpdatedAt() *time.Time
	}
)

type (
	ITree[ID any, C any] interface {
		Identify[ID]

		GetParentId() ID

		SetParentId(id ID)

		GetChildren() []C
	}

	ITreeSort[ID any, C any] interface {
		ITree[ID, C]
		ISort
	}
)
