package util

import (
	"reflect"
)

type Set map[interface{}]bool

func NewSet() Set {
	return make(Set)
}

func (set *Set) Append(element interface{}) *Set {
	(*set)[element] = true
	return set
}

func (set *Set) AppendAll(element interface{}) *Set {
	s := reflect.ValueOf(element)
	if s.Kind() == reflect.Slice || s.Kind() == reflect.Array {
		// Cycles throgh the slice, appending each inner element to the set
		for i := 0; i < s.Len(); i++ {
			set.Append(s.Index(i).Interface())
		}
	} else {
		panic("AppendAll only takes Slices or Arrays.")
	}
	return set
}

func (set *Set) Remove(element interface{}) *Set {
	delete(*set, element)
	return set
}

func (set *Set) Contains(element interface{}) bool {
	return (*set)[element] == true
}
