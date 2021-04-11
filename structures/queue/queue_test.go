package queue

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	q := New(4)
	t.Run("IsEmpty()", func(t *testing.T) {
		want := true
		got := q.IsEmpty()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("EnQueue() DeQueue()", func(t *testing.T) {
		want := 1
		q.EnQueue(1)
		q.EnQueue("hello")
		got, _ := q.DeQueue()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("DeQueue() ", func(t *testing.T) {
		want := "hello"
		got, _ := q.DeQueue()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("DeQueue on null queue", func(t *testing.T) {
		want := "empty queue"
		_, got := q.DeQueue()

		if !reflect.DeepEqual(want, got.Error()) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("DeQueue() ", func(t *testing.T) {
		want := "full queue"
		q.EnQueue(2)
		q.EnQueue(3)
		q.EnQueue(4)
		q.EnQueue(34.32)

		got := q.EnQueue("sun")
		if !reflect.DeepEqual(want, got.Error()) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("IsFill() ", func(t *testing.T) {
		want := true
		got := q.IsFull()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("MakeEmpty() ", func(t *testing.T) {
		want := true
		q.MakeEmpty()
		got := q.IsEmpty()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}
