package hw04_lru_cache //nolint:golint, stylecheck
import (
	"fmt"
	"unsafe"
)

type ListItem struct {
	Value interface{}
	Prev  *ListItem
	Next  *ListItem
}

type List interface {
	Len() int                          // длина списка
	Front() *ListItem                  // первый ListItem
	Back() *ListItem                   // последний ListItem
	PushFront(v interface{}) *ListItem // добавить значение в начало
	PushBack(v interface{}) *ListItem  // добавить значение в конец
	Remove(i *ListItem)                // удалить элемент
	MoveToFront(i *ListItem)           // переместить элемент в начало

	Values() []interface{} // вернуть все значения списка без учета порядка следования
	Items() []*ListItem    // вернуть все элементы без учета порядка следования
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
	items map[*ListItem]int
}

func (d list) Len() int {
	return d.len
}

func (d list) Front() *ListItem {
	return d.front
}

func (d list) Back() *ListItem {
	return d.back
}

func (d *list) makeFront(li *ListItem) {
	li.Prev = d.front
	if li.Prev != nil {
		li.Prev.Next = li
	}
	d.front = li
	if d.back == nil {
		d.back = d.front
	}
}

func (d *list) PushFront(v interface{}) *ListItem {
	li := ListItem{v, nil, nil}
	d.makeFront(&li)
	d.items[d.front] = d.len
	d.len++
	return d.front
}

func (d *list) PushBack(v interface{}) *ListItem {
	li := ListItem{v, nil, nil}

	li.Next = d.back
	if li.Next != nil {
		li.Next.Prev = &li
	}

	d.back = &li
	if d.front == nil {
		d.front = d.back
	}

	d.items[d.back] = d.len
	d.len++

	return d.back
}

func (d *list) unlink(li *ListItem) {
	if _, ok := d.items[li]; !ok {
		panic(fmt.Sprintf("No item in list: %v %+v", unsafe.Pointer(li), li))
	}

	if li != d.back && li != d.front {
		li.Prev.Next = li.Next
		li.Next.Prev = li.Prev
	}
	if li == d.back {
		d.back = d.back.Next
		if li.Next != nil {
			li.Next.Prev = nil
		}
	}
	if li == d.front {
		d.front = d.front.Prev
		if li.Prev != nil {
			li.Prev.Next = nil
		}
	}
}

func (d *list) Remove(li *ListItem) {
	d.unlink(li)
	delete(d.items, li)
	d.len--
}

func (d *list) MoveToFront(li *ListItem) {
	if li == d.front {
		return
	}

	d.unlink(li)
	li.Next, li.Prev = nil, nil
	d.makeFront(li)
}

func (d list) Values() []interface{} {
	values := make([]interface{}, 0, d.len)
	for li := range d.items {
		values = append(values, li.Value)
	}
	return values
}

func (d list) Items() []*ListItem {
	items := make([]*ListItem, 0, d.len)
	for li := range d.items {
		items = append(items, li)
	}
	return items
}

func NewList() List {
	return &list{items: map[*ListItem]int{}}
}
