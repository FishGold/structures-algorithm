package stack

import "errors"

type ElementType interface{}
type PrtToNode *Node

type Node struct {
	Element ElementType
	Next    PrtToNode
}

type StackInterface interface {
	IsEmpty() bool
	Push(x ElementType)
	Top() (ElementType, error)
	Pop() ElementType
}

func NewStack() *Stask {
	return &Stask{
		Header: &Node{},
	}
}

type Stask struct {
	Header PrtToNode
}

/*Return true if stack is empty*/
func (s *Stask) IsEmpty() bool {
	return s.Header.Next == nil
}

/* Push Elment x */
func (s *Stask) Push(x ElementType) {
	var tmp_node PrtToNode = &Node{
		Element: x,
		Next:    s.Header.Next,
	}
	s.Header.Next = tmp_node
}

/* Return top Element */
func (s *Stask) Top() (ElementType, error) {
	if s.IsEmpty() {
		return nil, errors.New("Top with Empty Stack")
	}
	return s.Header.Next.Element, nil
}

/* Return top Element  and remove if from stack */
func (s *Stask) Pop() (ElementType, error) {
	if s.IsEmpty() {
		return nil, errors.New("Pop with Empty Stack")
	}
	var top_elment ElementType = s.Header.Next.Element
	s.Header.Next = s.Header.Next.Next
	return top_elment, nil
}
