package utils

func Pointer[T any](in T) (out *T) {
	return &in
}
