package comparators

func Int64CompareFunc(a, b int64) int {
	return int(a - b)
}

func IntCompareFunc(a, b int) int {
	return a - b
}
