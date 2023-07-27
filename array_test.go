package array

import (
	"fmt"
	"testing"
)

func TestArray_1(t *testing.T) {
	arr := New[int]()
	arr.Push(1)
	arr.Push(2)
	arr.Push(3)

	for i := 0; i < 3; i++ {
		val := arr.Lookup(i)
		if val != i+1 {
			t.FailNow()
		}
	}
}

func TestArray_2(t *testing.T) {
	expected := []string{"a", "b", "c"}

	arr := New[string]()
	arr.Push("a")
	arr.Push("b")
	arr.Push("c")

	for i := 0; i < 3; i++ {
		val := arr.Lookup(i)
		if val != expected[i] {
			t.FailNow()
		}
	}
}

func TestArray_3(t *testing.T) {
	expected := []string{"c", "b", "a"}

	arr := New[string]()
	arr.Push("a")
	arr.Push("b")
	arr.Push("c")
	arr.Reverse()

	for i := 0; i < 3; i++ {
		val := arr.Lookup(i)
		if val != expected[i] {
			t.FailNow()
		}
	}
}

func TestArray_4(t *testing.T) {
	arr := New[int]()

	arr.Push(1)
	arr.Push(2)
	arr.Push(3)
	value := arr.Pop()
	if value != 3 {
		t.FailNow()
	}
	value = arr.Pop()
	if value != 2 {
		t.FailNow()
	}
	value = arr.Pop()
	if value != 1 {
		t.FailNow()
	}
	count := arr.Length()
	if count != 0 {
		t.FailNow()
	}
}

func TestArray_Copy(t *testing.T) {
	arr := New[int]()
	arr.Push(1)
	arr.Push(2)
	arr.Push(3)
	arrCopy := arr.Copy()

	arr.ForEach(func(i int) {
		fmt.Println(i)
	})

	arrCopy.ForEach(func(i int) {
		fmt.Println(i)
	})
}

func TestArray_Clear(t *testing.T) {
	arr := New[int]()
	for i := 0; i < 10; i++ {
		arr.Push(i)
	}

	if arr.Length() != 10 {
		t.FailNow()
	}

	arr.Clear()

	if arr.Length() > 0 {
		t.FailNow()
	}
}

func BenchmarkArray_Push(b *testing.B) {
	arr := New[int]()
	for n := 0; n < b.N; n++ {
		arr.Push(n)
	}
}

func BenchmarkArray_Copy(b *testing.B) {
	arr := New[int]()
	arr.Push(1)
	arr.Push(2)
	arr.Push(3)
	for n := 0; n < b.N; n++ {
		arr.Copy()
	}
}
