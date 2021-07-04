package linkedlist

import "fmt"

// var idx int

type Node struct {
	data int
	// index int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Append(n int) {
	second := l.head
	newNode := &Node{data: n}
	l.head = newNode
	l.head.next = second
	// idx++
	l.length++
}

func (l *LinkedList) Show() {
	toPrint := l.head

	for {
		fmt.Printf("%d ", toPrint.data)

		toPrint = toPrint.next
		l.length--

		if l.length == 0 {
			break
		}
	}

	fmt.Printf("\n")
}

// func (l *LinkedList) index(i int) {
// 	toPrint := l.head
// 	var find int

// 	for {
// 		if toPrint.index == i {
// 			find = toPrint.data
// 			break
// 		}

// 		toPrint = toPrint.next
// 		l.length--

// 		if l.length == 0 {
// 			break
// 		}
// 	}

// 	if find == 0 {
// 		fmt.Println("not found index")
// 	}

// 	fmt.Println(find)
// }

func (l *LinkedList) Delete(n int) {
	var count = 0
	for {
		if l.head.data == n {
			l.head.next = l.head.next.next
			l.length--
			break
		}

		l.head = l.head.next
		count++

		if count == l.length {
			break
		}
	}
}

func Try() {
	mylist := LinkedList{}

	mylist.Append(10)
	mylist.Append(14)
	mylist.Append(9)
	mylist.Append(100)
	mylist.Append(200)
	mylist.Append(500)

	mylist.Show()

	mylist.Delete(100)

	mylist.Show()

	// mylist.Delete(14)

	// mylist.Show()

	// mylist.index(0)

}
