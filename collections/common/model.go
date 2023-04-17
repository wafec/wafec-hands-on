package common

type CompareFunc[T comparable] func(a, b T) int
