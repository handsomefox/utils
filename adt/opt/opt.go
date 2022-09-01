// package option tries to mimick the option data type from rust.
package opt

type Type uint8

const (
	_ = iota
	None
	Some
)

type Option[T any] struct {
	value T
	typ   Type
}

// WithValue creates a new option with type Some and the provided value.
func WithValue[T any](value T) Option[T] {
	return Option[T]{
		value: value,
		typ:   Some,
	}
}

// Empty is a shortcut for creating empty option.
func Empty[T any]() Option[T] {
	return Option[T]{
		typ: None,
	}
}

// Type returns the option type (either option.None or option.Some).
func (option Option[T]) Type() Type {
	return option.typ
}

// Value returns the value under the option.
func (option Option[T]) Value() T {
	return option.value
}

// Unwrap panics if the option has no value, or returns the value.
func (option Option[T]) Unwrap() T {
	if option.typ == None {
		panic("Unwrap() called on empty option")
	}
	return option.value
}

// ValueOr returns the option value if it has one, or returns the Or value.
func (option Option[T]) ValueOr(or T) T {
	if option.typ == None {
		return or
	}
	return option.value
}

// Expect panics with the given message if the option has no value, or returns the value.
func (option Option[T]) Expect(message string) T {
	if option.typ == None {
		panic(message)
	}
	return option.value
}

// HasValue indicates whether the option has a value.
func (option Option[T]) HasValue() bool {
	return option.typ != None
}
