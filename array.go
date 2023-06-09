package array

type Array[T any] struct {
	length int
	values []T
}

// New creates a new array.
func New[T any]() *Array[T] {
	return &Array[T]{
		length: 0,
		values: make([]T, 0),
	}
}

// Length returns the number of values in the array.
func (arr *Array[T]) Length() int {
	return arr.length
}

// Lookup the value at a given index.
func (arr *Array[T]) Lookup(index int) T {
	var value T
	if index < 0 || index > arr.length {
		return value
	} else {
		value = arr.values[index]
	}
	return value
}

// Push adds an item to the end of the array and returns the length of the array.
func (arr *Array[T]) Push(value T) int {
	arr.values = append(arr.values, value)
	arr.length++
	return arr.length
}

// Pop removes the last value in the array and returns the item.
func (arr *Array[T]) Pop() T {
	lastValue := arr.values[arr.length-1]
	// :a.length-1 goes up to (but excluding the last item)
	arr.values = arr.values[:arr.length-1]
	arr.length--
	return lastValue
}

// Delete the value at a given index.
func (arr *Array[T]) Delete(index int) {
	// If index is less than zero or greater than array size do nothing
	if index < 0 || index > arr.length {
		return
	}
	// start at the index of the item to delete and loop until end of the array
	for i := index; i < arr.length-1; i++ {
		// move the item to the left by one
		arr.values[i] = arr.values[i+1]
	}
	// exclude the last item
	arr.values = arr.values[:arr.length-1]
	arr.length--
}

// Reverse the order of values in the array.
func (arr *Array[T]) Reverse() {
	// if array is empty or has one item just return
	if arr.length < 2 {
		return
	}
	// create a new array that is same size
	array := make([]T, arr.length)
	// start at zero of new array
	n := 0
	// move from right to left
	for i := arr.length; i > 0; i-- {
		// last item moves to front of new array
		array[n] = arr.values[i-1]
		// move one to the right of the new array
		n++
	}
	arr.values = array
}
