package common

type CompareTo[T comparable] func(a, b T) int64

func Int64CompareTo(a, b int64) int64 {
	return a - b
}
