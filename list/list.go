package list

type List[T any] struct {
	value T
	next  *List[T]
}

func New[T any](values []T) (result *List[T]) {
	if len(values) < 1 {
		return nil
	}
	result = &List[T]{value: values[0]}
	curr := result
	for i := 1; i < len(values); i++ {
		prev := curr
		curr = &List[T]{value: values[i]}
		prev.next = curr
	}
	return
}

func (l *List[T]) IsEmpty() bool {
	return l == nil
}

func (l *List[T]) Length() (result int) {
	if l == nil {
		return 0
	}
	for ; l != nil; l = l.next {
		result += 1
	}
	return
}

func (l *List[T]) Tail() *List[T] {
	for ; l != nil; l = l.next {
		if l.next == nil {
			return l
		}
	}
	return nil
}

func (l *List[T]) TailPredecessor() *List[T] {
	for ; l != nil && l.next != nil; l = l.next {
		if l.next.next == nil {
			return l
		}
	}
	return nil
}

func (l *List[T]) PushHead(value T) *List[T] {
	return &List[T]{
		value: value,
		next:  l,
	}
}

func (l *List[T]) PushTail(value T) *List[T] {
	tail := l.Tail()
	next := &List[T]{value: value}
	if tail == nil {
		return next
	}
	tail.next = next
	return l
}

func (l *List[T]) PopHead() (*List[T], T) {
	return l.next, l.value
}

func (l *List[T]) PopTail() (*List[T], T) {
	tailPrev := l.TailPredecessor()
	result := tailPrev.next.value
	tailPrev.next = nil
	return l, result
}

func (l *List[T]) Array() (result []T) {
	for ; l != nil; l = l.next {
		result = append(result, l.value)
	}
	return
}

func (l *List[T]) Reverse() (result *List[T]) {
	for ; l != nil; l = l.next {
		result = result.PushHead(l.value)
	}
	return
}

func (l *List[T]) ReverseInPlace() *List[T] {
	var prev, prevPrev *List[T]
	for ; l != nil; l = l.next {
		if prev != nil {
			prev.next = prevPrev
		}
		prevPrev = prev
		prev = l
	}
	if prev != nil {
		prev.next = prevPrev
	}
	return prev
}
