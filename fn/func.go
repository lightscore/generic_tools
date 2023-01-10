package fn

// functional stuff aka higher-order functions

func Map[T, U any](xs []T, f func(T) U) []U {
	ys := make([]U, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

func IndexMap[T, U any](xs []T, f func(int, T) U) []U {
	ys := make([]U, 0, len(xs))
	for i, x := range xs {
		ys = append(ys, f(i, x))
	}
	return ys
}

func Filter[T any](xs []T, f func(T) bool) []T {
	ys := make([]T, 0)
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

func Reduce[T, U any](init T, xs []U, f func(T, U) T) T {
	for _, x := range xs {
		init = f(init, x)
	}
	return init
}

func Any[T any](xs []T, f func(T) bool) bool {
	for _, x := range xs {
		if f(x) {
			return true
		}
	}
	return false
}

func All[T any](xs []T, f func(T) bool) bool {
	for _, x := range xs {
		if !f(x) {
			return false
		}
	}
	return true
}
