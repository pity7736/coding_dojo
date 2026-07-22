package linkedlists_test

import (
	"data_structures/linked_lists"
	"testing"
)

func TestLinkedListShould(t *testing.T) {
	t.Run("have no head value", func(t *testing.T) {
		ll := linkedlists.New()

		if ll.Head() != nil {
			t.Errorf(`expected value nil, got %p`, ll.Head())
		}
	})

	t.Run("append a value", func(t *testing.T) {
		ll := linkedlists.New()
		ll.Append(1)
		expected := 1

		if ll.Head().Value() != expected {
			t.Errorf("expected value: %d, got %d", expected, ll.Head().Value())
		}
	})

	t.Run("append another value", func(t *testing.T) {
		ll := linkedlists.New()
		ll.Append(2)
		expected := 2

		if ll.Head().Value() != expected {
			t.Errorf("expected value: %d, got %d", expected, ll.Head().Value())
		}
	})

	t.Run("append a second value", func(t *testing.T) {
		ll := linkedlists.New()
		ll.Append(1)
		ll.Append(2)
		expected := 2

		if ll.Tail().Value() != expected {
			t.Errorf("expected value: %d, got %d", expected, ll.Tail().Value())
		}
	})

	t.Run("get value by index", func(t *testing.T) {
		ll := linkedlists.New()
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)

		expected := 2
		got, err := ll.Get(1)

		if got != expected {
			t.Errorf("expected value: %d, got %d", expected, got)
		}
		if err != nil {
			t.Errorf("expected err nil, got: %v", err)
		}
	})

	t.Run("get the last value by index", func(t *testing.T) {
		ll := linkedlists.New()
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)

		expected := 3
		got, err := ll.Get(2)

		if got != expected {
			t.Errorf("expected value: %d, got %d", expected, got)
		}
		if err != nil {
			t.Errorf("expected err nil, got: %v", err)
		}
	})

	t.Run("get error when list is empty", func(t *testing.T) {
		ll := linkedlists.New()

		expected := 0
		got, err := ll.Get(0)

		if got != expected {
			t.Errorf("expected value: %d, got %d", expected, got)
		}
		if err.Error() != "list is empty" {
			t.Errorf("expected err should be list is empty, got %s", err.Error())
		}

	})

	t.Run("get error when index is out of range", func(t *testing.T) {
		ll := linkedlists.New()
		ll.Append(1)

		expected := 0
		got, err := ll.Get(1)

		if got != expected {
			t.Errorf("expected value: %d, got %d", expected, got)
		}
		if err.Error() != "index out of range" {
			t.Errorf("expected err should be index out of range, got %s", err.Error())
		}

	})

	t.Run("get the lenght", func(t *testing.T) {
		ll := linkedlists.New()
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)
		ll.Append(4)

		if ll.Len() != 4 {
			t.Errorf("expected value: 4, got %d", ll.Len())
		}
	})

	t.Run("delete by value", func(t *testing.T) {
		ll := linkedlists.New()
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)
		ll.Append(4)

		ll.Delete(2)

		first, _ := ll.Get(0)
		second, _ := ll.Get(1)
		third, _ := ll.Get(2)

		if first != 1 {
			t.Errorf("expected value: 1, got %d", first)
		}
		if second != 3 {
			t.Errorf("expected value: 3, got %d", second)
		}
		if third != 4 {
			t.Errorf("expected value: 4, got %d", third)
		}

		if ll.Len() != 3 {
			t.Errorf("expected value: 3, got %d", ll.Len())
		}

	})

	t.Run("delete by value the head", func(t *testing.T) {
		ll := linkedlists.New()
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)
		ll.Append(4)

		ll.Delete(1)

		if ll.Head().Value() != 2 {
			t.Errorf("expected head to be 2, got %d", ll.Tail().Value())
		}
	})

	t.Run("delete by value the tail", func(t *testing.T) {
		ll := linkedlists.New()
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)
		ll.Append(4)

		ll.Delete(4)

		if ll.Tail().Value() != 3 {
			t.Errorf("expected tail to be 3, got %d", ll.Tail().Value())
		}
	})

	t.Run("return error when delete on empty list", func(t *testing.T) {
		ll := linkedlists.New()

		err := ll.Delete(1)

		if err.Error() != "list is empty" {
			t.Errorf("expected list is empty error, got %s", err.Error())
		}
	})

	t.Run("return error when delete value is not found", func(t *testing.T) {
		ll := linkedlists.New()
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)
		ll.Append(4)

		err := ll.Delete(5)

		if err.Error() != "value not found" {
			t.Errorf("expected value not found error, got %s", err.Error())
		}
	})

}
