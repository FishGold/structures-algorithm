package tree

type ElementType int
type Position *TreeNode

type TreeNode struct {
	Element ElementType
	Left    Position
	Right   Position
}

/*
	二叉查找树：对于树中每个节点X，它的左子树中所有节点的值都小于X，右子树中所有节点的值都大于X
	这里假设所有的节点的值可比较且互异
*/
type SearchTreeInterface interface {
	MakeEmpty()
	Find(x ElementType, t Position) Position
	FindMin(t Position) Position
	FindMax(t Position) Position
	Insert(x ElementType)
	Delete(x ElementType)
	Retrieve(p Position)
}

type SearchTree struct {
	Roof Position
}

func (s *SearchTree) Find(x ElementType, t Position) Position {
	if t == nil {
		return nil
	}

	if x < t.Element {
		return s.Find(x, t.Left)
	} else if x > t.Element {
		return s.Find(x, t.Right)
	}
	return t

}

func (s *SearchTree) FindMin(t Position) Position {
	if t == nil {
		return nil
	} else if t.Left == nil {
		return t
	} else {
		return s.FindMin(t.Left)
	}

}

func (s *SearchTree) FindMax(t Position) Position {
	if t == nil {
		return nil
	} else if t.Right == nil {
		return t
	} else {
		return s.FindMax(t.Right)
	}

}

func (s *SearchTree) Insert(x ElementType) {

}
