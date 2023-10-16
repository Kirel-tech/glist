package glist

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewList(t *testing.T) {
	l := NewList[TestStruct]()
	if l == nil {
		t.Error("failed to create new IList")
	}
}

func TestList_Where(t *testing.T) {
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
	l := NewList[TestStruct](test, test1, test2, test3)
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
