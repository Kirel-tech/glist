package glist

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type TestStruct struct {
	Name          string
	Date          time.Time
	DateNullable  *time.Time
	Number        int
	NumberPointer *int
}

func TestNewListBase(t *testing.T) {
	l := NewListBase[TestStruct]()
	if l == nil {
		t.Error("failed to create new IListBase")
	}
}

func TestListBase_Add(t *testing.T) {
	nowTime := time.Now()
	nowTimePointer := &nowTime
	testStruct := TestStruct{
		Name:          "Test",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	l := &listBase[TestStruct]{}
	l.Add(testStruct)
	assert.Equal(t, 1, len(l.s))
	assert.Equal(t, l.s[0], testStruct)
	l.Add(testStruct)
	assert.Equal(t, 2, len(l.s))
	assert.Equal(t, l.s[1], testStruct)
	l.Add(testStruct, testStruct)
	assert.Equal(t, 4, len(l.s))
}

func TestListBase_Count(t *testing.T) {
	nowTime := time.Now()
	nowTimePointer := &nowTime
	testStruct := TestStruct{
		Name:          "Test",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	l := &listBase[TestStruct]{}
	l.Add(testStruct)
	assert.Equal(t, 1, l.Count())
	l.Add(testStruct)
	assert.Equal(t, 2, l.Count())
	l.Add(testStruct, testStruct)
	assert.Equal(t, 4, len(l.s))
}

func TestListBase_FindFirst(t *testing.T) {
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
	test1 := TestStruct{
		Name:          "Test1",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	test3 := TestStruct{
		Name:          "",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	l := &listBase[TestStruct]{}

	_, err := l.FindFirst(func(elem TestStruct) bool {
		if len(elem.Name) > 4 {
			return true
		}
		return false
	})
	assert.Error(t, err)
	assert.Equal(t, ErrNoElements, err.Error())

	l.Add(test, test1, test2, test3)
	first, err := l.FindFirst(func(elem TestStruct) bool {
		if len(elem.Name) > 4 {
			return true
		}
		return false
	})
	assert.NoError(t, err)
	assert.Equal(t, test1, first)
	first, err = l.FindFirst(func(elem TestStruct) bool {
		if len(elem.Name) > 3 {
			return true
		}
		return false
	})
	assert.NoError(t, err)
	assert.Equal(t, test, first)
	first, err = l.FindFirst(func(elem TestStruct) bool {
		if len(elem.Name) > 5 {
			return true
		}
		return false
	})
	assert.Error(t, err)
	assert.Equal(t, ErrNotFound, err.Error())
}

func TestListBase_First(t *testing.T) {
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
		Name:          "T",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	l := &listBase[TestStruct]{}
	_, err := l.First()
	assert.Error(t, err)
	assert.Equal(t, ErrNoElements, err.Error())

	l.Add(test, test1)
	first, err := l.First()
	assert.NoError(t, err)
	assert.Equal(t, test, first)
}

func TestListBase_Last(t *testing.T) {
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
		Name:          "T",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	l := &listBase[TestStruct]{}
	_, err := l.Last()
	assert.Error(t, err)
	assert.Equal(t, ErrNoElements, err.Error())

	l.Add(test, test1)
	last, err := l.Last()
	assert.NoError(t, err)
	assert.Equal(t, test1, last)
}

func TestListBase_FindLast(t *testing.T) {
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
	test1 := TestStruct{
		Name:          "Test1",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	test3 := TestStruct{
		Name:          "",
		Date:          nowTime,
		DateNullable:  nowTimePointer,
		Number:        0,
		NumberPointer: nil,
	}
	l := &listBase[TestStruct]{}
	_, err := l.FindLast(func(elem TestStruct) bool {
		if len(elem.Name) > 4 {
			return true
		}
		return false
	})
	assert.Error(t, err)
	assert.Equal(t, ErrNoElements, err.Error())

	l.Add(test, test1, test2, test3)
	last, err := l.FindLast(func(elem TestStruct) bool {
		if len(elem.Name) > 4 {
			return true
		}
		return false
	})
	assert.NoError(t, err)
	assert.Equal(t, test1, last)
	last, err = l.FindLast(func(elem TestStruct) bool {
		if len(elem.Name) > 3 {
			return true
		}
		return false
	})
	assert.NoError(t, err)
	assert.Equal(t, test1, last)
	last, err = l.FindLast(func(elem TestStruct) bool {
		if len(elem.Name) > 5 {
			return true
		}
		return false
	})
	assert.Error(t, err)
	assert.Equal(t, ErrNotFound, err.Error())
}

func TestListBase_ForEach(t *testing.T) {
	callCount := 0
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
	l := &listBase[TestStruct]{}
	l.Add(test, test2)
	l.ForEach(func(elem TestStruct) {

		switch callCount {
		case 0:
			assert.Equal(t, test, elem)
		case 1:
			assert.Equal(t, test2, elem)
		}
		callCount++
	})
}

func TestListBase_ToSlice(t *testing.T) {
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
	source := []TestStruct{test, test2}
	l := &listBase[TestStruct]{}
	l.Add(test, test2)
	result := l.ToSlice()
	assert.Equal(t, source, result)
}
