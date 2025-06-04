package result

func Ok[T any](value T) Result[T] {
	return Result[T]{value: value}
}

func Error[T any](err error) Result[T] {
	return Result[T]{err: err}
}

type Result[T any] struct {
	value T
	err   error
}

func (r Result[T]) Val() T {
	return r.value
}

func (r Result[T]) Err() error {
	return r.err
}

func (r Result[T]) IsError() bool {
	return r.err != nil
}
