package listv2

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

func TestList(t *testing.T) {
	tx := testutil.NewTx(t)

	{
		// simple
		vs := []int{1, 2, 3, 4, 5}
		nodeMap := map[int]*Node[int]{}
		ls := New[int]()
		for _, v := range vs {
			node := ls.PushBack(v)
			nodeMap[v] = node
		}
		listVs := ls.Values()
		tx.AssertEqual([]int{1, 2, 3, 4, 5}, listVs)
		//

		// remove front
		ls.Remove(ls.Front())
		listVs = ls.Values()
		tx.AssertEqual([]int{2, 3, 4, 5}, listVs)

		// remove back
		ls.Remove(ls.Back())
		listVs = ls.Values()
		tx.AssertEqual([]int{2, 3, 4}, listVs)

		// move 3 to back
		ls.MoveBack(nodeMap[3])
		listVs = ls.Values()
		tx.AssertEqual([]int{2, 4, 3}, listVs)

		listVs = ls.ValuesReverse()
		tx.AssertEqual([]int{3, 4, 2}, listVs)

		// push some more
		newVs := []int{20, 30, 40}
		for _, v := range newVs {
			node := ls.PushFront(v)
			nodeMap[v] = node
		}
		listVs = ls.Values()
		tx.AssertEqual([]int{40, 30, 20, 2, 4, 3}, listVs)

		n, ok := findFirst(ls, 40)
		tx.AssertEqual(true, ok)
		ls.InsertBefore(n, 42)
		tx.AssertEqual([]int{42, 40, 30, 20, 2, 4, 3}, ls.Values())

		n, ok = findFirst(ls, 3)
		tx.AssertEqual(true, ok)
		ls.InsertAfter(n, 31)
		tx.AssertEqual([]int{42, 40, 30, 20, 2, 4, 3, 31}, ls.Values())

		n, ok = findFirst(ls, 30)
		tx.AssertEqual(true, ok)
		ls.InsertBefore(n, 35)
		tx.AssertEqual([]int{42, 40, 35, 30, 20, 2, 4, 3, 31}, ls.Values())

		n, ok = findFirst(ls, 4)
		tx.AssertEqual(true, ok)
		ls.InsertAfter(n, 666)
		tx.AssertEqual([]int{42, 40, 35, 30, 20, 2, 4, 666, 3, 31}, ls.Values())

		n, ok = findFirst(ls, 2)
		tx.AssertEqual(true, ok)
		ls.Remove(n)
		tx.AssertEqual([]int{42, 40, 35, 30, 20, 4, 666, 3, 31}, ls.Values())

		tx.AssertEqual(9, ls.Count())
	}
}

func findFirst(ls *List[int], v int) (*Node[int], bool) {
	for n := ls.Front(); n != nil; n = n.Next() {
		if n.data == v {
			return n, true
		}
	}
	return nil, false
}
