package glist

type list[T any] struct {
	*listBase[T]
}

func NewList[T any](elems ...T) IList[T] {
	return &list[T]{
		&listBase[T]{s: elems},
	}
}

func (l *list[T]) Where(whereF func(elem T) bool) IList[T] {
	newList := NewList[T]()
	l.ForEach(func(elem T) {
		if whereF(elem) {
			newList.Add(elem)
		}
	})
	return newList
}
