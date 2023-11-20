package sorted_set

import "golang.org/x/exp/constraints"

type TreeNode[T constraints.Ordered] struct {
	Left  *TreeNode[T]
	Right *TreeNode[T]
	Index int64
	Value T
}
type SortedSet[T constraints.Ordered] struct {
	BST *TreeNode[T]
}

func (o *SortedSet[T]) Insert(value T) {
	cur := o.BST
	for cur != nil {
		if value <= cur.Value {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	cur = &TreeNode[T]{Value: value}
	if o.BST == nil {
		o.BST = cur
	}
}

func (o *SortedSet[T]) Delete(value T) {

}
