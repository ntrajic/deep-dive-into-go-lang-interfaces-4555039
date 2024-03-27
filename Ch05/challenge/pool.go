package pool

type Pool[T any] struct {
}

func New[T any](newFn func() T) *Pool[T] {
	return nil
}

func (p *Pool[T]) Put(v T) {
}

func (p *Pool[T]) Get() (T, bool) {
	var zero T
	return zero, false
}
