package optional

type Optional[T any] struct {
	present bool
	value   T
}

func Of[T any](value T) Optional[T] {
	return Optional[T]{true, value}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

func (o Optional[T]) Present() bool {
	return o.present
}

func (o Optional[T]) Empty() bool {
	return !o.present
}

func (o Optional[T]) Value() T {
	if !o.present {
		panic("can't get value of empty optional")
	}
	return o.value
}

func (o Optional[T]) Map(mapper func(T) T) Optional[T] {
	if o.Empty() {
		return o
	}

	return Of(mapper(o.Value()))
}
