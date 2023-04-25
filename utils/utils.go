package utils

import "reflect"

func IsFilterPresent[T any](filter T) bool {
	return !reflect.ValueOf(&filter).Elem().IsZero()
}
