package main

import "fmt"

// Структура двусвязного списка
type DoublyLinkedList struct {
	head  *Node
	tail  *Node
	count int
}

// Структура ветви
type Node struct {
	value int
	next  *Node
	prev  *Node
}

func initDoublyList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Количество элементов в двусвязном списке
func (list *DoublyLinkedList) Size() int {
	return list.count
}

// Проверка списка на наличие элементов
func (list *DoublyLinkedList) IsEmpty() bool {
	return list.count == 0
}

// Первый элемент списка
func (list *DoublyLinkedList) Peek() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListError")
		return 0, false
	}
	return list.head.value, true
}

// Добавление элемента в начало списка
func (list *DoublyLinkedList) AddHead(value int) {
	newNode := &Node{value, nil, nil}
	if list.count == 0 {
		list.tail = newNode
		list.head = newNode
	} else {
		list.head.prev = newNode
		newNode.next = list.head
		list.head = newNode
	}
	list.count++
}

// Добавление элемента в конец списка
func (list *DoublyLinkedList) AddTail(value int) {
	newNode := &Node{value, nil, nil}
	if list.count == 0 {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
	list.count++
}

// Удалить первый элемент списка
func (list *DoublyLinkedList) RemoveHead() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListError")
		return 0, false
	}
	value := list.head.value
	list.head = list.head.next
	if list.head == nil {
		list.tail = nil
	} else {
		list.head.prev = nil
	}
	list.count--
	return value, true
}

// Удалить узел по значению
func (list *DoublyLinkedList) RemoveNode(key int) bool {
	curr := list.head
	if curr == nil { // пустой список
		return false
	}
	if curr.value == key { // если удаляемый элемент первый
		curr = curr.next
		list.count--
		if curr != nil {
			list.head = curr
			list.head.prev = nil
		} else {
			list.tail = nil // если в списке только один элемент
		}
		return true
	}
	for curr.next != nil {
		if curr.next.value == key {
			curr.next = curr.next.next
			if curr.next == nil { // если удаляемый элемент последний
				list.tail = curr
			} else {
				curr.next.prev = curr
			}
			list.count--
			return true
		}
		curr = curr.next
	}
	return false
}

// Поиск элемента в списке
func (list *DoublyLinkedList) IsPresent(key int) bool {
	temp := list.head
	for temp != nil {
		if temp.value == key {
			return true
		}
		temp = temp.next
	}
	return false
}

// Удаление всех элементов в списке
func (list *DoublyLinkedList) FreeList() {
	list.tail = nil
	list.head = nil
	list.count = 0
}

// Печать всех элементов списка
func (list *DoublyLinkedList) Print() {
	temp := list.head
	for temp != nil {
		fmt.Println(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

// Перевернуть список
func (list *DoublyLinkedList) ReverseList() {
	curr := list.head
	var tempNode *Node
	for curr != nil {
		tempNode = curr.next
		curr.next = curr.prev
		curr.prev = tempNode
		if curr.prev == nil {
			list.tail = list.head
			list.head = curr
			return
		}
		curr = curr.prev
	}
	return
}

// Скопировать список в обратном порядке
func (list *DoublyLinkedList) CopyListReversed(dll *DoublyLinkedList) {
	curr := list.head
	for curr != nil {
		dll.AddHead(curr.value)
		curr = curr.next
	}
}

// Просто скопировать список
func (list *DoublyLinkedList) CopyList(dll *DoublyLinkedList) {
	curr := list.head
	for curr != nil {
		dll.AddTail(curr.value)
		curr = curr.next
	}
}

// Вставка в список с сортировкой
func (list *DoublyLinkedList) SortedInsert(value int) {
	temp := &Node{value, nil, nil}
	curr := list.head
	if curr == nil { // first element
		list.head = temp
		list.tail = temp
	}
	if list.head.value <= value { // at the begining
		temp.next = list.head
		list.head.prev = temp
		list.head = temp
	}
	for curr.next != nil && curr.next.value > value { // treversal
		curr = curr.next
	}
	if curr.next == nil { // at the end
		list.tail = temp
		temp.prev = curr
		curr.next = temp
	} else { // all other
		temp.next = curr.next
		temp.prev = curr
		curr.next = temp
		temp.next.prev = temp
	}
}

// Удаление повторяющихся значений в отсортированном списке
func (list *DoublyLinkedList) RemoveDuplicate() {
	curr := list.head
	var deleteMe *Node
	for curr != nil {
		if (curr.next != nil) && curr.value == curr.next.value {
			deleteMe = curr.next
			curr.next = deleteMe.next
			curr.next.prev = curr
			if deleteMe == list.tail {
				list.tail = curr
			}
		} else {
			curr = curr.next
		}
	}
}

// Удалить все значения меньше последнего добавленного
func (list *DoublyLinkedList) DeleteLess() {
	if list.count != 1 {
		temp := list.tail
		for temp.value >= temp.prev.value && temp.prev != list.head {
			list.RemoveNode(temp.prev.value)

		}
		if temp.prev == list.head && temp.value >= list.head.value {
			list.RemoveHead()
		}

		//fmt.Println()
	}
}

// Разница между текущим и следующим
// Удалить все значения меньше последнего добавленного
func (list *DoublyLinkedList) Difference() {
	temp := list.head
	for temp != list.tail {
		temp.value = temp.value - temp.next.value
		temp = temp.next
	}

}

func main() {
	myList := initDoublyList()
	myList.AddHead(7)
	myList.DeleteLess()
	myList.AddTail(3)
	myList.DeleteLess()
	myList.AddTail(4)
	myList.DeleteLess()
	myList.AddTail(2)
	myList.DeleteLess()
	myList.AddTail(2)
	myList.DeleteLess()

	myList.Difference()
	myList.Print()

}
