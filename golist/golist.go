/*
   Implementation of Linked List API
*/

package golist

type List struct {
    head *Node
}

type Node struct {
    next  *Node
    value interface{}
}

func Create() *List {
    return &List{nil}
}

func (l *List) Insert(value interface{}) {
    node := &Node{
        next:  l.head,
        value: value,
    }
    l.head = node
}

func (l *List) Head() *Node {
    /*
       Get head of the List
    */
    return l.head
}

func (l *List) Tail() *Node {
    /*
       Get tail of the List
    */
    current := l.head
    for current.next != nil {
        current = current.next
    }
    return current
}

func (l *List) Find(value interface{}) *Node {
    /*
       Find first equal to value node or return nil
    */
    current := l.head
    for i := 0; current != nil; i++ {
        if current.value == value {
            return current
        }
        current = current.next
    }
    return nil
}
