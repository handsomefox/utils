// package option tries to mimick the option data type from rust.
package opt

type Type uint8

const (
	_ = iota
	None
	Some
)

type Option[T any] struct {
	Val T
	Typ Type
}

// WithValue creates a new option with type Some and the provided value.
func WithValue[T any](value T) Option[T] {
	return Option[T]{
		Val: value,
		Typ: Some,
	}
}

// Empty is a shortcut for creating empty option.
func Empty[T any]() Option[T] {
	return Option[T]{
		Typ: None,
	}
}

// Type returns the option type (either option.None or option.Some).
func (opt Option[T]) Type() Type {
	return opt.Typ
}

// Value returns the value under the option.
func (opt Option[T]) Value() T {
	return opt.Val
}

// Unwrap panics if the option has no value, or returns the value.
func (opt Option[T]) Unwrap() T {
	if opt.Typ == None {
		panic("Unwrap() called on empty option")
	}
	return opt.Val
}

// ValueOr returns the option value if it has one, or returns the Or value.
func (opt Option[T]) ValueOr(or T) T {
	if opt.Typ == None {
		return or
	}
	return opt.Val
}

// Expect panics with the given message if the option has no value, or returns the value.
func (opt Option[T]) Expect(message string) T {
	if opt.Typ == None {
		panic(message)
	}
	return opt.Val
}

// HasValue indicates whether the option has a value.
func (opt Option[T]) HasValue() bool {
	return opt.Typ != None
}
