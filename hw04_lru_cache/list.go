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
	//elements map[*ListItem]struct{}
	front *ListItem
	back  *ListItem
	len   int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newFront := &ListItem{
		Next:  l.front,
		Prev:  nil,
		Value: v,
	}

	if l.Len() == 0 {
		l.back = newFront
		l.front = newFront
	} else {
		l.front.Prev = newFront
		l.front = newFront
	}

	l.len++
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	newBack := ListItem{
		Next:  nil,
		Prev:  l.back,
		Value: v,
	}

	if l.Len() == 0 {
		l.back = &newBack
		l.front = &newBack
	} else {
		l.back.Next = &newBack
		l.back = &newBack
	}

	l.len++
	return l.back
}

func (l *list) Remove(i *ListItem) {
	if i == l.front {
		l.front = i.Next
	}
	if i == l.back {
		l.back = i.Prev
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.front {
		return
	}
	l.Remove(i)
	i.Prev = nil
	i.Next = l.front
	if l.front != nil {
		l.front.Prev = i
	}
	l.front = i
	if l.back == nil {
		l.back = i
	}
	l.len++
}

func NewList() List {
	return &list{}
}
