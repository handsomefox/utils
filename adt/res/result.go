// package res tries to mimick the result data type from rust.
package res

type Type uint8

const (
	_ = iota
	Value
	Error
)

type Result[T any] struct {
	Val T
	Err error
}

// WithValue creates a new result with nil error and the provided value.
func WithValue[T any](value T) Result[T] {
	return Result[T]{
		Val: value,
		Err: nil,
	}
}

// WithError is a shortcut for creating Result with (or without) an error.
func WithError[T any](err error) Result[T] {
	return Result[T]{
		Err: err,
	}
}

// Error returns the result err.
func (res Result[T]) Error() error {
	return res.Err
}

// Value returns the value under the result.
func (res Result[T]) Value() T {
	return res.Val
}

// Unwrap panics if the result has an error, or returns the value.
func (res Result[T]) Unwrap() T {
	if res.Err != nil {
		panic(res.Err)
	}
	return res.Val
}

// ValueOr returns the result value if it has one, or returns the Or value.
func (res Result[T]) ValueOr(or T) T {
	if res.Err != nil {
		return or
	}
	return res.Val
}

// Expect panics with the given message if the resulthas no value, or returns the value.
func (res Result[T]) Expect(message string) T {
	if res.Err != nil {
		panic(message)
	}
	return res.Val
}

// IsError indicates whether the result has an error.
func (res Result[T]) IsError() bool {
	return res.Err != nil
}
