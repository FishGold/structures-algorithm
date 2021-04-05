package linked_list

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	// 初始化链表
	list := New()

	t.Run("Test IsEmpty()", func(t *testing.T) {
		var want = true
		got := list.IsEmpty()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("Test InsertAtBegin()、First()、Retrieve()", func(t *testing.T) {
		var want = 1
		list.InsertAtBegin(1)
		got := list.Retrieve(list.First())
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test InsertAtEnd()", func(t *testing.T) {
		var want = "single"
		list.InsertAtEnd(want)
		got := list.Retrieve(list.Last())
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Find()", func(t *testing.T) {
		var want = 12.11
		list.InsertAtBegin(want)
		got := list.Retrieve(list.Find(want))
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Delete() with end", func(t *testing.T) {
		var want = "single"
		list.Delete(12.11)
		got := list.Retrieve(list.Last())
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Delete() with begin", func(t *testing.T) {
		want := []int64{1, 2, 3}
		list.InsertAtBegin(want)
		list.Delete(1)
		got := list.Retrieve(list.First())
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}
