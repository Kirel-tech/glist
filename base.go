package glist

import "fmt"

type listBase[T any] struct {
	s []T
}

func NewListBase[T any](elems ...T) IListBase[T] {
	return &listBase[T]{s: elems}
}

func (l *listBase[T]) Add(elems ...T) {
	l.s = append(l.s, elems...)
}

func (l *listBase[T]) Count() int {
	return len(l.s)
}

func (l *listBase[T]) ForEach(eachF func(elem T)) {
	for _, elem := range l.s {
		eachF(elem)
	}
}

func (l *listBase[T]) First() (T, error) {
	if l.Count() < 1 {
		return *new(T), fmt.Errorf(ErrNoElements)
	}
	return l.s[0], nil
}

func (l *listBase[T]) FindFirst(findF func(elem T) bool) (T, error) {
	if l.Count() < 1 {
		return *new(T), fmt.Errorf(ErrNoElements)
	}
	for _, elem := range l.s {
		if findF(elem) {
			return elem, nil
		}
	}
	return *new(T), fmt.Errorf(ErrNotFound)
}

func (l *listBase[T]) Last() (T, error) {
	if l.Count() < 1 {
		return *new(T), fmt.Errorf(ErrNoElements)
	}
	return l.s[l.Count()-1], nil
}

func (l *listBase[T]) FindLast(findF func(elem T) bool) (T, error) {
	if l.Count() < 1 {
		return *new(T), fmt.Errorf(ErrNoElements)
	}
	for i := len(l.s) - 1; i >= 0; i-- {
		if findF(l.s[i]) {
			return l.s[i], nil
		}
	}
	return *new(T), fmt.Errorf(ErrNotFound)
}

func (l *listBase[T]) ToSlice() []T {
	return l.s
}
