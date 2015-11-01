package util_test

import (
	"testing"

	"github.com/daime/redis-metrics/util"
)

const (
	ELEMENT = "element1"
)

var (
	ELEMENT_ARRAY = [4]string{
		"element1",
		"element2",
		"element3",
		"element4",
	}

	ELEMENT_SLICE = []string{
		"element1",
		"element2",
		"element3",
		"element4",
	}
)

func Test_Append_NoExisting_IncludeAElement(t *testing.T) {
	set := util.NewSet()
	set.Append(ELEMENT)
	if !set.Contains(ELEMENT) {
		t.Fail()
	}
}

func Test_AppendAll_NoExistingArray_IncludeAllElements(t *testing.T) {
	set := util.NewSet()
	set.AppendAll(ELEMENT_ARRAY)
	for _, element := range ELEMENT_ARRAY {
		if !set.Contains(element) {
			t.Fail()
		}
	}
}

func Test_AppendAll_NoExistingSlice_IncludeAllElements(t *testing.T) {
	set := util.NewSet()
	set.AppendAll(ELEMENT_SLICE)
	for _, element := range ELEMENT_SLICE {
		if !set.Contains(element) {
			t.Fail()
		}
	}
}

func Test_AppendAll_SimpleElement_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if r != "AppendAll only takes Slices or Arrays." {
				t.Fail()
			}
		}
	}()
	set := util.NewSet()
	set.AppendAll(ELEMENT)
	t.Fail()
}

func Test_Remove_Existing_RemoveTheElement(t *testing.T) {
	set := util.NewSet()
	set.Append(ELEMENT)
	set.Remove(ELEMENT)
	if set.Contains(ELEMENT) {
		t.Fail()
	}
}
