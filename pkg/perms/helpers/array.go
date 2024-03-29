package helpers

import (
	"reflect"
)

// InArray is the value in the first parameter an element of the array in the second parameter?
// @param interface{}
// @param interface{}
// return bool
func InArray(val interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}

// JoinUintArrays concatenates the given uint arrays and makes them a single array.
// @param ...[]uint
// return []uint
func JoinUintArrays(array ...[]uint) (j []uint) {
	for _, a := range array {
		j = append(j, a...)
	}
	return
}

// RemoveDuplicateValues make singular of repeating values in an array.
// @param []uint
// return []uint
func RemoveDuplicateValues(intSlice []uint) []uint {
	keys := make(map[uint]bool)
	list := []uint{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
