package goptr

import "reflect"

// Ptr returns a pointer to the given value.
func Ptr[T any](v T) *T {
	return &v
}

// Unptr returns the value pointed to by ptr.
func Unptr[T any](ptr *T) T {
	return *ptr
}

// PtrOrNil returns a pointer to the given value, or nil if the value is a zero value.
func PtrOrNil[T any](v T) *T {
	if reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface()) {
		return nil
	}
	return &v
}

// UnptrOrZero returns the value pointed to by ptr, or the zero value if the pointer is nil.
func UnptrOrZero[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}
