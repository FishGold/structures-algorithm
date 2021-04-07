package stack

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	t.Run("IsEmpty()", func(t *testing.T) {
		want := true
		got := stack.IsEmpty()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Push(), Top()", func(t *testing.T) {
		want := 12
		stack.Push(12)
		got, err := stack.Top()
		if err != nil {
			t.Errorf("err:%v", err.Error())
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Pop(), Top()", func(t *testing.T) {
		want := 12
		got, err := stack.Pop()
		if err != nil {
			t.Errorf("err:%v", err.Error())
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("Pop() null stack , Top() null stack", func(t *testing.T) {
		want := "Pop with Empty Stack"
		_, got := stack.Pop()
		if !reflect.DeepEqual(got.Error(), want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run(" Top() null stack", func(t *testing.T) {
		want := "Top with Empty Stack"
		_, got := stack.Top()
		if !reflect.DeepEqual(got.Error(), want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
