package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	elements map[*ListItem]struct{}
	front    *ListItem
	back     *ListItem
}

func (l *list) Len() int {
	return len(l.elements)
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newFront := ListItem{
		Next:  l.front,
		Prev:  nil,
		Value: v,
	}

	if len(l.elements) == 0 {
		l.back = &newFront
		l.front = &newFront
	} else {
		l.front.Prev = &newFront
		l.front = &newFront
	}

	l.elements[&newFront] = struct{}{}
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	newBack := ListItem{
		Next:  nil,
		Prev:  l.back,
		Value: v,
	}

	if len(l.elements) == 0 {
		l.back = &newBack
		l.front = &newBack
	} else {
		l.back.Next = &newBack
		l.back = &newBack
	}

	l.elements[&newBack] = struct{}{}
	return l.back
}

func (l *list) Remove(i *ListItem) {
	if len(l.elements) == 0 {
		return
	}

	if _, exists := l.elements[i]; exists {
		switch {
		case i.Prev == nil && i.Next == nil:
			l.front = nil
			l.back = nil
		case i.Prev == nil:
			l.back = i.Next
		case i.Next == nil:
			l.front = i.Prev
		default:
			i.Next.Prev = i.Prev
			i.Prev.Next = i.Next
		}

		delete(l.elements, i)
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if _, exists := l.elements[i]; exists {
		switch {
		case i == l.front:
			return
		case i == l.back:
			l.back = i.Prev
			i.Prev.Next = nil
			i.Next = l.front
			l.front.Prev = i
		default:
			i.Prev = nil
			i.Next = l.front
			i.Next.Prev = i.Prev
			l.front.Prev = i
		}

		l.front = i
	}
}

func NewList() List {
	return &list{
		elements: make(map[*ListItem]struct{}),
	}
}
