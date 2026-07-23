package gotl

type Bimap[T1 comparable, T2 comparable] interface {
	GetByT1(key T1) (T2, error)
	GetByT2(key T2) (T1, error)
}

func NewBimap[T1 comparable, T2 comparable](
	fetchByT1Fn func(T1) (T2, error),
	fetchByT2Fn func(T2) (T1, error),
) Bimap[T1, T2] {
	return &bimapImpl[T1, T2]{
		fetchByT1Fn,
		fetchByT2Fn,
		map[T1]T2{},
		map[T2]T1{},
	}
}

type bimapImpl[T1 comparable, T2 comparable] struct {
	fetchByT1Fn func(T1) (T2, error)
	fetchByT2Fn func(T2) (T1, error)

	t1ToT2 map[T1]T2
	t2ToT1 map[T2]T1
}

func (m *bimapImpl[T1, T2]) GetByT1(key T1) (T2, error) {
	t2, hasKey := m.t1ToT2[key]
	var err error
	if !hasKey {
		t2, err = m.fetchByT1Fn(key)
	}

	return t2, err
}

func (m *bimapImpl[T1, T2]) GetByT2(key T2) (T1, error) {
	t1, hasKey := m.t2ToT1[key]
	var err error
	if !hasKey {
		t1, err = m.fetchByT2Fn(key)
	}

	return t1, err
}
