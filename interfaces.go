package glist

// IListWhere defines interface for where'able IList
type iListWhere[T any, TList IListBase[T]] interface {
	// Where returns new TList with applied filter
	Where(func(elem T) bool) TList
}

// IListBase defines base interface for list implementation
type IListBase[T any] interface {
	// Add elements to list
	Add(elems ...T)
	// Count elements to list
	Count() int
	// ForEach iterates over every element in list
	ForEach(func(elem T))
	// First element from list
	First() (T, error)
	// FindFirst element with filtering
	FindFirst(func(elem T) bool) (T, error)
	// Last element from list
	Last() (T, error)
	// FindLast element with filtering
	FindLast(func(elem T) bool) (T, error)
	// ToSlice returns slice
	ToSlice() []T
}

// IList defines list interface
type IList[T any] interface {
	IListBase[T]
	iListWhere[T, IList[T]]
}

// IListComparable defines list interface for comparable elements
type IListComparable[T comparable] interface {
	IListBase[T]
	iListWhere[T, IListComparable[T]]
	// Contains returns a flag that the element is in the list
	Contains(elem T) bool
	// Delete element from list
	Delete(elem T)
	// Distinct return new list of elements without duplicates
	Distinct() IListComparable[T]
}
