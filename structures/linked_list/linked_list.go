package linked_list

type ElementType interface{}
type Position *Node

type Node struct {
	Element ElementType
	Next    Position
}

type linkedListInterface interface {
	IsEmpty() bool
	IsLast(p Position) bool
	Find(x ElementType) Position
	Delete(x ElementType)
	Insert(x ElementType, p Position) Position
	InsertAtBegin(x ElementType) Position
	InsertAtEnd(x ElementType) Position
	First() Position
	Last() Position
	Retrieve(p Position) ElementType
}

/* Single LinkedList implementation place a header */
type LinkedList struct {
	Header Position
}

func New() *LinkedList {
	return &LinkedList{
		Header: &Node{},
	}
}

/* Return true if  link_list is empty */
func (l *LinkedList) IsEmpty() bool {
	return l.Header.Next == nil
}

/* Return true if p is the last position of list*/
func (l *LinkedList) IsLast(p Position) bool {
	return p.Next == nil
}

/* Return Position of x in list; nil if not found */
func (l *LinkedList) Find(x ElementType) Position {
	var p = l.Header
	for p != nil {
		if p.Element == x {
			return p
		}
		p = p.Next
	}
	return nil
}

// Delete first occurrence of x from list
func (l *LinkedList) Delete(x ElementType) {
	var next Position
	var p = l.Header
	for p != nil {
		next = p.Next
		if next.Element == x {
			p.Next = next.Next
			p = nil // set p nil
			return
		}
		p = next
	}

}

/* Insert after legal position p*/
func (l *LinkedList) Insert(x ElementType, p Position) Position {
	var tmp_node = &Node{
		Element: x,
		Next:    p.Next,
	}
	p.Next = tmp_node
	return p.Next
}

/* Return the position of first node */
func (l *LinkedList) First() Position {
	return l.Header.Next
}

/* Return the position of last node */
func (l *LinkedList) Last() Position {
	var p = l.Header
	for p.Next != nil {
		p = p.Next
	}
	return p
}

/* Insert x at the begin of list */
func (l *LinkedList) InsertAtBegin(x ElementType) Position {
	first := l.First()
	if first == nil {
		var tmp_node = &Node{
			Element: x,
		}
		l.Header.Next = tmp_node
		return tmp_node
	}
	return l.Insert(x, first)
}

/* Insert x at the begin of list */
func (l *LinkedList) InsertAtEnd(x ElementType) Position {
	last := l.Last()
	return l.Insert(x, last)
}

/* Retrieve node value of p */
func (l *LinkedList) Retrieve(p Position) ElementType {
	return p.Element
}
