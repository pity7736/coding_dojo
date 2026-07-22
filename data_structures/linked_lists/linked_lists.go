package linkedlists

import "errors"

type node struct {
	value int
	next  *node
}

func newNode(value int, next *node) *node {
	return &node{value: value, next: next}
}

func (self *node) Value() int {
	return self.value
}

type LinkedList struct {
	head   *node
	tail   *node
	lenght int
}

func New() *LinkedList {
	return &LinkedList{}
}

func (self *LinkedList) Head() *node {
	return self.head
}

func (self *LinkedList) Append(value int) {
	n := newNode(value, nil)
	if self.head == nil {
		self.head = n
		self.tail = n
	} else {
		t := self.tail
		t.next = n
		self.tail = n
	}
	self.lenght++
}

func (self *LinkedList) Tail() *node {
	return self.tail
}

func (self *LinkedList) Get(i int) (int, error) {
	if self.head == nil {
		return 0, errors.New("list is empty")
	}
	if i > self.lenght-1 {
		return 0, errors.New("index out of range")
	}

	result := self.head
	for range i {
		result = result.next
	}
	return result.value, nil
}

func (self *LinkedList) Delete(value int) error {
	if self.head == nil {
		return errors.New("list is empty")
	}
	dummy := &node{next: self.head}
	current := dummy
	for current.next != nil && current.next.value != value {
		current = current.next
	}
	if current.next == nil {
		return errors.New("value not found")
	}
	current.next = current.next.next
	if current.next == nil {
		self.tail = current
	}
	self.head = dummy.next
	self.lenght--
	return nil
}

func (self *LinkedList) Len() int {
	return self.lenght
}
