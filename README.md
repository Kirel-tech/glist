[![codecov](https://codecov.io/gh/Kirel-tech/glist/branch/main/graph/badge.svg)](https://codecov.io/gh/Kirel-tech/glist)
[![Go Report Card](https://goreportcard.com/badge/github.com/Kirel-tech/glist)](https://goreportcard.com/report/github.com/Kirel-tech/glist)

# GList
GList is an additional layer of abstraction over slice that provides some new functionality for data manipulation.
This implementation is implemented according to OOP principles, and was written for internal purposes.

## Interfaces
There are several interfaces for managing data based on its type and purpose. Each interface extends previous one.
* `IListBase` - base interface definition for basic data read nad add functionality.
* `IList` - interface definition for **not comparable** elements.
* `IListComparable` - interface definition for **comparable** elements.

### interfaces methods
```
IListBase.Add(elems ...T)
IListBase.Count() int
IListBase.ForEach(func(elem T))
IListBase.First() (T, error)
IListBase.FindFirst(func(elem T) bool) (T, error)
IListBase.Last() (T, error)
IListBase.FindLast(func(elem T) bool) (T, error)
IListBase.ToSlice() []T
IList.Where(func(elem T) bool) IList
IListComparable.Where(func(elem T) bool) IListComparable
IListComparable.Contains(elem T) bool
IListComparable.Delete(elem T)
IListComparable.Distinct() IListComparable[T]
```

## Implementations
Each implementation has constructor:
* `NewListBase()` - creates new `IListBase` implementation.
* `NewList()` - creates new `IList` implementation.
* `NewListComparable()` - creates new `IListComparable` implementation.

## Examples
You can found a lot of examples in tests.
### Select emulation
We cannot declare an additional type in the declaration of a separate interface method, so implementing the Select method is impossible. 
To do something like this, I propose a trivial solution:
```go
l := NewListComparable[SomeSourceType]()
selcetedElems := NewListComparable[SomeDestType]()
l.ForEach(func(elem SomeSourceType) {
    selcetedElems.Add(SomeDestType{
	    // fill data	
    })
})
```
