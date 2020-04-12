//Домашнее задание
//Двусвязный список
//Цель: https://en.wikipedia.org/wiki/Doubly_linked_list?Ожидаемые типы (псевдокод):?```
//List // тип контейнер
//Len() // длинна списка
//First() // первый
//Item Last() // последний
//Item PushFront(v interface{}) // добавить значение в начало
//PushBack(v interface{}) // добавить значение в конец
//Remove(i Item) // удалить элемент
//
//Item // элемент списка
//Value() interface{} // возвращает значение
//Nex() *Item // следующий
//Item Prev() *Item // предыдущий ```
//Реализовать двусвязанный список на языке Go

package main

import (
	"fmt"
)

type Item struct {
	Value int
	Next *Item
	Prev *Item
}

type List struct {
	First *Item
	Last *Item
}


// ITEM METHODS //
func (item_ *Item) GetItemValue() int {
	return item_.Value
}

func (item_ *Item) GetNextItem() *Item {
	return item_.Next
}

func (item_ *Item) GetPrevItem() *Item {
	return item_.Prev
}

// LIST METHODS //
func (list_ *List) InitList() {
	list_ = &List{nil, nil}
}

func (list_ *List) GetLen() int {
	len_ := 0
	FirstNode := list_.GetFirstNode()
	if FirstNode == nil {
		return len_
	} else {
		CurrentElement  := FirstNode
		for true {
			len_ += 1
			if CurrentElement.Next == nil {
				return len_
			} else {
				CurrentElement = CurrentElement.Next
			}
		}
	}
	return len_
}

func (list_ *List) GetFirstNode() *Item {
	return list_.First
}

func (list_ *List) GetLastNode() *Item {
	return list_.Last
}

func (list_ *List) PushFront(item *Item) {
	if list_.GetFirstNode() == nil {
		list_.First = item
		list_.Last = item
	} else {
		list_.First.Prev = item
		item.Next = list_.First
		list_.First = item
	}
}

func (list_ *List) PushBack(item *Item) {
	if list_.GetLastNode() == nil {
		list_.First = item
		list_.Last = item
	} else {
		list_.Last.Next = item
		item.Prev = list_.Last
		list_.Last = item
	}
}

func (list_ *List) Remove(elementPos int) {
	//
	CurrentNode := list_.GetFirstNode()
	if CurrentNode != nil {
		i := 0
		for true {
			if i == 0 && CurrentNode.Next == nil {
				// IF LIST HAVE ONLY ONE ITEM
				list_.InitList()
			} else {
				// IF LIST HAVE MORE THAN ONE ITEM
				if i == elementPos {
					if CurrentNode.Prev == nil {
						// IF REMOVE FIRST ITEM IN LIST
						CurrentNode.Next.Prev = nil
						break
					} else if CurrentNode.Next == nil {
						// IF REMOVE LAST ITEM IN LIST
						CurrentNode.Prev.Next = nil
						break
					} else {
						CurrentNode.Prev.Next = CurrentNode.Next
						CurrentNode.Next.Prev = CurrentNode.Prev
						CurrentNode = CurrentNode.Next
						break
					}
				}
				CurrentNode = CurrentNode.Next
			}
			i += 1
		}
	} else {
		// IF LIST EMPTY
		list_.InitList()
	}
}

func main() {
	tryItem1 := Item{1, nil, nil}
	tryItem2 := Item{666, nil, nil}
	tryItem3 := Item{3, nil, nil}
	tryItem4 := Item{4, nil, nil}

	var list_ List
	list_.InitList()
	fmt.Printf("len now %v\n", list_.GetLen())

	list_.PushFront(&tryItem1)
	list_.PushFront(&tryItem2)
	list_.PushFront(&tryItem3)
	list_.PushBack(&tryItem4)
	fmt.Printf("len now %v\n", list_.GetLen())
	list_.Remove(2)
	fmt.Printf("len now %v\n", list_.GetLen())
}
