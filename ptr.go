package goptr

func Ptr[T any](v T) *T {
	return &v
}

func Unptr[T any](ptr *T) T {
	return *ptr
}
