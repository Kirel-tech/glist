package glist

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewListComparable(t *testing.T) {
	l := NewListComparable[TestStruct]()
	if l == nil {
		t.Error("failed to create new IListComparable")
	}
}

func TestListComparable_Contains(t *testing.T) {
	nowTime := time.Now()
	nowTimePointer := &nowTime
	test := TestStruct{
		Name:          "Test",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	test2 := TestStruct{
		Name:          "T",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	l := NewListComparable[TestStruct]()
	l.Add(test, test2)
	exist := l.Contains(test)
	assert.Equal(t, true, exist)
	notExist := !(l.Contains(TestStruct{}))
	assert.Equal(t, true, notExist)
}

func TestListComparable_Delete(t *testing.T) {
	nowTime := time.Now()
	nowTimePointer := &nowTime
	test := TestStruct{
		Name:          "Test",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	test2 := TestStruct{
		Name:          "T",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	l := NewListComparable[TestStruct](test, test2)
	l.Delete(test2)
	assert.Equal(t, 1, l.Count())
	first, err := l.First()
	assert.NoError(t, err)
	assert.Equal(t, first, test)
}

func TestListComparable_Distinct(t *testing.T) {
	nowTime := time.Now()
	nowTimePointer := &nowTime
	test := TestStruct{
		Name:          "Test",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	test2 := TestStruct{
		Name:          "T",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	test3 := TestStruct{
		Name:          "T",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	l := NewListComparable[TestStruct](test, test2, test3)
	newL := l.Distinct()
	assert.Equal(t, 2, newL.Count())
}

func TestListComparable_Where(t *testing.T) {
	nowTime := time.Now()
	nowTimePointer := &nowTime
	test := TestStruct{
		Name:          "Test",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	test1 := TestStruct{
		Name:          "123456",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	test2 := TestStruct{
		Name:          "T",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	test3 := TestStruct{
		Name:          "T",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	l := NewListComparable[TestStruct](test, test1, test2, test3)
	filteredButSameElementsL := l.Where(func(elem TestStruct) bool {
		if elem.Date == nowTime {
			return true
		}
		return false
	})
	assert.Equal(t, l.Count(), filteredButSameElementsL.Count())
	filteredL := l.Where(func(elem TestStruct) bool {
		if elem.Name == "Test" {
			return true
		}
		return false
	})
	assert.Equal(t, 1, filteredL.Count())
	filteredLNext := l.Where(func(elem TestStruct) bool {
		if elem.Name == "T" {
			return true
		}
		return false
	})
	assert.Equal(t, 2, filteredLNext.Count())
}
