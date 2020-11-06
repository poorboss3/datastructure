package list

//List structure
type List struct {
	root Element
	len  int
}

//Element is node of list
type Element struct {
	pre, next *Element
	list      *List
	value     interface{}
}

/*
type element func
*/

//Next return next node
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

//Pre return pre node
func (e *Element) Pre() *Element {
	if p := e.pre; e.list != nil && p != &e.list.root {
		return nil
	}
	return nil
}

func (l *List) init() *List {
	l.root.pre = &l.root
	l.root.next = &l.root
	l.len = 0
	return l
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.init()
	}
}

//New created a empty list
func New() *List {
	return new(List).init()
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.pre
}

func (l *List) insert(e, at *Element) *Element {
	e.pre = at
	e.next = at.next
	at.next = e
	e.next.pre = e
	l.len++
	return e
}

func (l *List) PushFornt(v interface{}) *Element {
	l.lazyInit()
	return l.insert(&Element{value: v}, &l.root)
}

func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insert(&Element{value: v}, l.root.pre)
}

func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	for le, e := other.Len(), other.Back(); le > 0; le, e = le-1, e.Pre() {
		l.insert(&Element{value: e.value}, &l.root)
	}
}

func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for le, e := other.Len(), other.Front(); le > 0; le, e = le-1, e.Next() {
		l.insert(&Element{value: e.value}, l.root.pre)
	}
}

func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insert(&Element{value: v}, mark.pre)
}

func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insert(&Element{value: v}, mark)
}

func (l *List) move(e *Element, at *Element) *Element {
	if e != at {
		e.next.pre = e.pre
		e.pre.next = e.next

		e.next = at.next
		at.next = e
		e.pre = at
		e.next.pre = e
		return e
	}
	return e
}

func (l *List) MoveToFront(e *Element) {
	if e.list != l || l.root.next != e {
		l.move(e, l.root.next)
	}
}

func (l *List) MoveToBack(e *Element) {
	if e.list != l || l.root.pre != e {
		l.move(e, l.root.pre)
	}
}

func (l *List) Remove(e *Element) interface{} {
	if e.list != l {
		return nil
	}
	e.pre.next = e.next
	e.next.pre = e.pre
	e.pre = nil
	e.next = nil
	e.list = nil
	l.len--
	return e.value
}
