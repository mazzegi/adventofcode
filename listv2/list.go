package listv2

func New[T any]() *List[T] {
	return &List[T]{}
}

type Node[T any] struct {
	data T
	prev *Node[T]
	next *Node[T]
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

func (n *Node[T]) Data() T {
	return n.data
}

func (n *Node[T]) SetData(t T) {
	n.data = t
}

type List[T any] struct {
	front *Node[T]
	back  *Node[T]
}

func (list *List[T]) Front() *Node[T] {
	return list.front
}

func (list *List[T]) Back() *Node[T] {
	return list.back
}

func (list *List[T]) PushBack(data T) *Node[T] {
	node := &Node[T]{data: data}
	list.PushBackNode(node)
	return node
}

func (list *List[T]) PushBackNode(node *Node[T]) {
	if list.back == nil { // front is also nil in this case
		list.back = node
		list.front = node
		return
	}
	list.back.next = node
	node.prev = list.back
	list.back = node
}

func (list *List[T]) PushFront(data T) *Node[T] {
	node := &Node[T]{data: data}
	list.PushFrontNode(node)
	return node
}

func (list *List[T]) PushFrontNode(node *Node[T]) {
	if list.front == nil { // back is also nil in this case
		list.front = node
		list.back = node
		return
	}
	list.front.prev = node
	node.next = list.front
	list.front = node
}

func (list *List[T]) InsertBefore(before *Node[T], data T) *Node[T] {
	if before == list.front {
		return list.PushFront(data)
	}
	node := &Node[T]{data: data}
	node.prev = before.prev
	node.next = before

	before.prev.next = node
	before.prev = node

	return node
}

func (list *List[T]) InsertAfter(after *Node[T], data T) *Node[T] {
	if after == list.back {
		return list.PushBack(data)
	}
	node := &Node[T]{data: data}
	node.prev = after
	node.next = after.next

	after.next.prev = node
	after.next = node

	return node
}

func (list *List[T]) Remove(node *Node[T]) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		// node was front
		list.front = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		// node was back
		list.back = node.prev
	}
}

func (list *List[T]) MoveBack(node *Node[T]) {
	list.Remove(node)
	node.prev = nil
	node.next = nil
	list.PushBackNode(node)
}

func (list *List[T]) Values() []T {
	var vs []T
	curr := list.front
	for curr != nil {
		vs = append(vs, curr.Data())
		curr = curr.next
	}
	return vs
}

func (list *List[T]) ValuesReverse() []T {
	var vs []T
	curr := list.back
	for curr != nil {
		vs = append(vs, curr.Data())
		curr = curr.prev
	}
	return vs
}

func (list *List[T]) Count() int {
	var cnt int
	for n := list.Front(); n != nil; n = n.Next() {
		cnt++
	}
	return cnt
}
