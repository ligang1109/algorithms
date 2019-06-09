package problem

import "math"

type IntBTree struct {
	Value int
	Left  *IntBTree
	Right *IntBTree
}

func ValidIntBst(btree *IntBTree) bool {
	return validIntBst(btree, math.MinInt64, math.MaxInt64)
}

func validIntBst(btree *IntBTree, min, max int) bool {
	if btree == nil {
		return true
	}

	if btree.Value <= min || btree.Value >= max {
		return false
	}

	if !validIntBst(btree.Left, min, btree.Value) {
		return false
	}

	return validIntBst(btree.Right, btree.Value, max)
}
