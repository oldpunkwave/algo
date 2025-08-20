package list

import "fmt"

type (
	List struct {
		head *Node
	}

	Node struct {
		val  int
		next *Node
	}
)

func New() *List {
	return &List{}
}

func (l *List) Reverse() {
	var prev *Node
	curr := l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	l.head = prev
}

func (l *List) Push(val int) {
	nn := &Node{val: val}

	if l.head == nil {
		l.head = nn

		return
	}

	head := l.head

	for head.next != nil {
		head = head.next
	}

	head.next = nn
}

func (l *List) PushFront(val int) {
	nn := &Node{val: val}

	if l.head == nil {
		l.head = nn

		return
	}

	nn.next = l.head
	l.head = nn
}

func (l *List) Remove(val int) {
	if l.head == nil {
		return
	}

	if l.head.val == val {
		l.head = l.head.next

		return
	}

	prev, curr := l.head, l.head.next

	for curr != nil {
		if curr.val == val {
			prev.next = curr.next

			return
		}

		prev = curr
		curr = curr.next
	}
}

func (l *List) Traverse() {
	head := l.head

	for head != nil {
		fmt.Printf("%d -> ", head.val)

		head = head.next
	}

	fmt.Println("nil")
}
