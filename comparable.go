package glist

type listComparable[T comparable] struct {
	*list[T]
}

func NewListComparable[T comparable](elems ...T) IListComparable[T] {
	return &listComparable[T]{
		&list[T]{
			&listBase[T]{s: elems},
		},
	}
}

func (l *listComparable[T]) Where(whereF func(elem T) bool) IListComparable[T] {
	newList := NewListComparable[T]()
	l.ForEach(func(elem T) {
		if whereF(elem) {
			newList.Add(elem)
		}
	})
	return newList
}

func (l *listComparable[T]) Contains(elem T) bool {
	for _, el := range l.s {
		if el == elem {
			return true
		}
	}
	return false
}

func (l *listComparable[T]) Delete(elem T) {
	for index, el := range l.s {
		if el == elem {
			l.s[index] = l.s[len(l.s)-1]
			l.s = l.s[:len(l.s)-1]
			return
		}
	}
}

func (l *listComparable[T]) Distinct() IListComparable[T] {
	newList := NewListComparable[T]()
	l.ForEach(func(elem T) {
		if !newList.Contains(elem) {
			newList.Add(elem)
		}
	})
	return newList
}
