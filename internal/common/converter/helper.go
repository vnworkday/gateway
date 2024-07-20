package converter

type ConvertFunc[T any] func(T) error
