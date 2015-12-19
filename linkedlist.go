package main

import "fmt"


type node struct {
    value int
    next *node
}


type LinkedList struct {
    head *node
}


func (ll *LinkedList) insert (value int) bool {
    if ll.head == nil || ll.head.value > value {
        new_node := new(node)
        new_node.value = value
        new_node.next = ll.head
        ll.head = new_node

        return true
    }

    curr_node := ll.head
    for curr_node.next != nil && curr_node.next.value < value {
        if curr_node.next.value == value {
            return false
        }

        curr_node = curr_node.next
    }

    new_node := new(node)
    new_node.value = value
    new_node.next = curr_node.next
    curr_node.next = new_node

    return true
}

func (ll *LinkedList) show () {
    curr_node := ll.head

    for curr_node != nil {
        fmt.Println(curr_node.value)
        curr_node = curr_node.next
    }
}


func main() {
    ll := new(LinkedList)

    ll.insert(1)
    ll.insert(3)
    ll.insert(5)
    ll.insert(2)
    ll.insert(4)
    ll.show()
}
