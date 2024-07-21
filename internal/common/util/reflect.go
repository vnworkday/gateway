package util

import (
	"reflect"

	"github.com/gookit/goutil/reflects"
)

func GetOrDefault[T any](value T, defaultValue T) T {
	if reflects.IsEmpty(reflect.ValueOf(value)) {
		return defaultValue
	}

	return value
}

func ToPointer[T any](value T) *T {
	return &value
}
