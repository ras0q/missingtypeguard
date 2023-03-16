package missingtypeguard

import (
	"go/types"

	"golang.org/x/tools/go/types/typeutil"
)

type typedMap[T any] struct {
	typeutil.Map
}

func (m *typedMap[T]) At(key types.Type) T {
	var t T
	if v := m.Map.At(key); v != nil {
		t = v.(T)
	}

	return t
}

func (m *typedMap[T]) Set(key types.Type, val T) {
	m.Map.Set(key, val)
}

func (m *typedMap[T]) Iterate(f func(key types.Type, value T)) {
	m.Map.Iterate(func(key types.Type, value interface{}) {
		f(key, value.(T))
	})
}
