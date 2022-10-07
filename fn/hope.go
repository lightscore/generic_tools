package fn

import (
	"fmt"
	"strings"
)

// one-liner panicking on errors

func Hope(f func() error, msg ...string) {
	if err := f(); err != nil {
		panic(fmt.Sprintf("PANIC %s | %s", strings.Join(msg, " | "), err))
	}
}

func HopeUno[T, U any](f func(T) (U, error), t T, msg ...string) U {
	u, err := f(t)
	if err != nil {
		panic(fmt.Sprintf("PANIC %s | %s", strings.Join(msg, " | "), err))
	}
	return u
}

func HopeDos[T, U, V any](f func(T, U) (V, error), t T, u U, msg ...string) V {
	v, err := f(t, u)
	if err != nil {
		panic(fmt.Sprintf("PANIC %s | %s", strings.Join(msg, " | "), err))
	}
	return v
}

func HopeTres[T, U, V, W any](f func(T, U, V) (W, error), t T, u U, v V, msg ...string) W {
	w, err := f(t, u, v)
	if err != nil {
		panic(fmt.Sprintf("PANIC %s | %s", strings.Join(msg, " | "), err))
	}
	return w
}
