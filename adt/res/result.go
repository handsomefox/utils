// package res tries to mimick the result data type from rust.
package res

type Type uint8

const (
	_ Type = iota
	Value
	Error
)

type Result[T any] struct {
	value T
	err   error
}

// WithValue creates a new result with nil error and the provided value.
func WithValue[T any](value T) Result[T] {
	return Result[T]{
		value: value,
		err:   nil,
	}
}

// WithError is a shortcut for creating Result with (or without) an error.
func WithError[T any](err error) Result[T] {
	return Result[T]{
		err: err,
	}
}

// Error returns the result err.
func (result Result[T]) Error() error {
	return result.err
}

// Value returns the value under the result.
func (result Result[T]) Value() T {
	return result.value
}

// Unwrap panics if the result has an error, or returns the value.
func (result Result[T]) Unwrap() T {
	if result.err != nil {
		panic(result.err)
	}
	return result.value
}

// ValueOr returns the result value if it has one, or returns the Or value.
func (result Result[T]) ValueOr(or T) T {
	if result.err != nil {
		return or
	}
	return result.value
}

// Expect panics with the given message if the resulthas no value, or returns the value.
func (result Result[T]) Expect(message string) T {
	if result.err != nil {
		panic(message)
	}
	return result.value
}

// IsError indicates whether the result has an error.
func (result Result[T]) IsError() bool {
	return result.err != nil
}
