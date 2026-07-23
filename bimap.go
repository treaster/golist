package gotl

import (
	"fmt"
)

type Bimap[T1 comparable, T2 comparable] interface {
	Add(T1, T2) error
	GetByT1(key T1) (T2, bool)
	GetByT2(key T2) (T1, bool)
}

func NewBimap[T1 comparable, T2 comparable](allowOverwrite bool) Bimap[T1, T2] {
	return &bimapImpl[T1, T2]{
		allowOverwrite,
		map[T1]T2{},
		map[T2]T1{},
	}
}

type bimapImpl[T1 comparable, T2 comparable] struct {
	allowOverwrite bool

	t1ToT2 map[T1]T2
	t2ToT1 map[T2]T1
}

func (m *bimapImpl[T1, T2]) Add(t1 T1, t2 T2) error {
	existingT2, ok1 := m.t1ToT2[t1]
	if ok1 && !m.allowOverwrite {
		return fmt.Errorf("attempt to overwrite bimap entry disallowed (%v, %v) -> (%v, %v)", t1, existingT2, t1, t2)
	}

	existingT1, ok2 := m.t2ToT1[t2]
	if ok2 && !m.allowOverwrite {
		return fmt.Errorf("attempt to overwrite bimap entry disallowed (%v, %v) -> (%v, %v)", existingT1, t2, t1, t2)
	}

	if ok1 && ok2 && existingT1 == t1 && existingT2 == t2 {
		return nil
	}

	if ok1 && existingT2 != t2 || ok2 && existingT1 != t1 {
		delete(m.t1ToT2, existingT1)
		delete(m.t2ToT1, existingT2)
	}

	m.t1ToT2[t1] = t2
	m.t2ToT1[t2] = t1

	return nil
}

func (m *bimapImpl[T1, T2]) GetByT1(key T1) (T2, bool) {
	t2, ok := m.t1ToT2[key]
	return t2, ok
}

func (m *bimapImpl[T1, T2]) GetByT2(key T2) (T1, bool) {
	t1, ok := m.t2ToT1[key]
	return t1, ok
}
