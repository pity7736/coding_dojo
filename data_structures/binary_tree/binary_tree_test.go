package binarytree_test

import (
	"data_structures/binary_tree"
	"testing"
)

func TestBinaryTreeShould(t *testing.T) {
	t.Run("has nil root when tree is empty", func(t *testing.T) {
		bt := binarytree.New()

		if bt.Root() != nil {
			t.Errorf("expected root be nil, got %p", bt.Root())
		}
	})

	t.Run("insert value", func(t *testing.T) {
		bt := binarytree.New()
		bt.Insert(2)

		if bt.Root().Value() != 2 {
			t.Errorf("expected value: 2, got %d", bt.Root().Value())
		}
	})

	t.Run("insert higher value at right", func(t *testing.T) {
		bt := binarytree.New()
		bt.Insert(2)
		expected := 3
		bt.Insert(expected)

		value := bt.Root().Right().Value()

		if value != expected {
			t.Errorf("expected right value: %d, got %d", expected, value)
		}
	})

	t.Run("insert lower value at left", func(t *testing.T) {
		bt := binarytree.New()
		bt.Insert(2)
		bt.Insert(3)
		expected := 1
		bt.Insert(expected)

		value := bt.Root().Left().Value()

		if value != expected {
			t.Errorf("expected left value: %d, got %d", expected, value)
		}
	})

	t.Run("insert three values", func(t *testing.T) {
		bt := binarytree.New()
		bt.Insert(4)
		bt.Insert(5)
		bt.Insert(3)
		expected := 2
		bt.Insert(expected)

		value := bt.Root().Left().Left().Value()

		if value != expected {
			t.Errorf("expected value: %d, got %d", expected, value)
		}
	})

	t.Run("ignore duplicate values", func(t *testing.T) {
		bt := binarytree.New()
		bt.Insert(2)
		bt.Insert(2)

		if bt.Root().Right() != nil || bt.Root().Left() != nil {
			t.Errorf("expected no duplicate to be inserted")
		}
	})
}

func TestBinaryTreeIterativeInsertShould(t *testing.T) {
	t.Run("insert value", func(t *testing.T) {
		bt := binarytree.New()
		bt.InsertIterative(2)

		if bt.Root().Value() != 2 {
			t.Errorf("expected value: 2, got %d", bt.Root().Value())
		}
	})

	t.Run("insert higher value at right", func(t *testing.T) {
		bt := binarytree.New()
		bt.InsertIterative(2)
		expected := 3
		bt.InsertIterative(expected)

		value := bt.Root().Right().Value()

		if value != expected {
			t.Errorf("expected right value: %d, got %d", expected, value)
		}
	})

	t.Run("insert lower value at left", func(t *testing.T) {
		bt := binarytree.New()
		bt.InsertIterative(2)
		bt.InsertIterative(3)
		expected := 1
		bt.InsertIterative(expected)

		value := bt.Root().Left().Value()

		if value != expected {
			t.Errorf("expected left value: %d, got %d", expected, value)
		}
	})

	t.Run("insert three values", func(t *testing.T) {
		bt := binarytree.New()
		bt.InsertIterative(4)
		bt.InsertIterative(5)
		bt.InsertIterative(3)
		expected := 2
		bt.InsertIterative(expected)

		value := bt.Root().Left().Left().Value()

		if value != expected {
			t.Errorf("expected value: %d, got %d", expected, value)
		}
	})

	t.Run("ignore duplicate values", func(t *testing.T) {
		bt := binarytree.New()
		bt.InsertIterative(2)
		bt.InsertIterative(2)

		if bt.Root().Right() != nil || bt.Root().Left() != nil {
			t.Errorf("expected no duplicate to be inserted")
		}
	})
}

func TestBinaryTreeSearchShould(t *testing.T) {
	t.Run("return true when value exists", func(t *testing.T) {
		bt := binarytree.New()
		bt.Insert(4)
		bt.Insert(2)
		bt.Insert(6)

		if !bt.Search(2) {
			t.Errorf("expected true, got false")
		}
	})

	t.Run("return false when value does not exist", func(t *testing.T) {
		bt := binarytree.New()
		bt.Insert(4)
		bt.Insert(2)
		bt.Insert(6)

		if bt.Search(5) {
			t.Errorf("expected false, got true")
		}
	})

	t.Run("return false when tree is empty", func(t *testing.T) {
		bt := binarytree.New()

		if bt.Search(1) {
			t.Errorf("expected false, got true")
		}
	})
}

func TestBinaryTreeSearchIterativeShould(t *testing.T) {
	t.Run("return true when value exists", func(t *testing.T) {
		bt := binarytree.New()
		bt.Insert(4)
		bt.Insert(2)
		bt.Insert(6)

		if !bt.SearchIterative(2) {
			t.Errorf("expected true, got false")
		}
	})

	t.Run("return false when value does not exist", func(t *testing.T) {
		bt := binarytree.New()
		bt.Insert(4)
		bt.Insert(2)
		bt.Insert(6)

		if bt.SearchIterative(5) {
			t.Errorf("expected false, got true")
		}
	})

	t.Run("return false when tree is empty", func(t *testing.T) {
		bt := binarytree.New()

		if bt.SearchIterative(1) {
			t.Errorf("expected false, got true")
		}
	})
}

func TestBinaryTreeTraversalShould(t *testing.T) {
	t.Run("in-order returns sorted values", func(t *testing.T) {
		bt := binarytree.New()
		bt.Insert(4)
		bt.Insert(2)
		bt.Insert(6)
		bt.Insert(1)
		bt.Insert(3)

		expected := []int{1, 2, 3, 4, 6}
		got := bt.InOrder()

		for i, v := range expected {
			if got[i] != v {
				t.Errorf("expected %d at index %d, got %d", v, i, got[i])
			}
		}
	})
}
