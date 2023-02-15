package list

import (
	"strconv"
	"strings"
	"testing"
)

func String(l *List[int]) string {
	if l == nil {
		return "nil"
	}
	var sb strings.Builder
	for ; l != nil; l = l.next {
		sb.WriteString(strconv.Itoa(l.value))
		if l.next != nil {
			sb.WriteString("->")
		}
	}
	return sb.String()
}

func TestString(t *testing.T) {
	l := New([]int{1, 2, 3})
	expected := "1->2->3"
	actual := String(l)
	if expected != actual {
		t.Fatalf("TestString = %s, want %s", actual, expected)
	}
}

func TestIsEmpty(t *testing.T) {
	var l *List[int]
	if !l.IsEmpty() {
		t.Fail()
	}
	l = l.PushHead(1)
	if l.IsEmpty() {
		t.Fail()
	}
}

func TestLength(t *testing.T) {
	l := New([]int{1, 2, 3})
	expected := 3
	actual := l.Length()
	if expected != actual {
		t.Fatalf("TestLength = %d, want %d", actual, expected)
	}
}

func TestLengthOfNil(t *testing.T) {
	var l *List[int]
	expected := 0
	actual := l.Length()
	if expected != actual {
		t.Fatalf("TestLengthOfNil = %d, want %d", actual, expected)
	}
}

func TestTail(t *testing.T) {
	l := New([]int{1, 2, 3})
	expected := 3
	actual := l.Tail().value
	if expected != actual {
		t.Fatalf("TestTail = %d, want %d", actual, expected)
	}
}

func TestTailOfNil(t *testing.T) {
	var l *List[int]
	var expected *List[int] = nil
	actual := l.Tail()
	if expected != actual {
		t.Fatalf("TestTailOfNil = %p, want %p", actual, expected)
	}
}

func TestTailPredecessor(t *testing.T) {
	l := New([]string{"a", "b", "c", "d"})
	expected := "c"
	actual := l.TailPredecessor().value
	if expected != actual {
		t.Fatalf("TestTailPredecessor = %s, want %s", actual, expected)
	}
}

func TestPushHead(t *testing.T) {
	l := New([]int{1, 2, 3})
	l = l.PushHead(4)
	expected := "4->1->2->3"
	actual := String(l)
	if expected != actual {
		t.Fatalf("TestPushHead = %s, want %s", actual, expected)
	}
}

func TestPushHeadOfNil(t *testing.T) {
	var l *List[int]
	l = l.PushHead(4)
	expected := "4"
	actual := String(l)
	if expected != actual {
		t.Fatalf("TestPushHeadOfNil = %s, want %s", actual, expected)
	}
}

func TestPushTail(t *testing.T) {
	l := New([]int{1, 2, 3})
	l = l.PushTail(4)
	expected := "1->2->3->4"
	actual := String(l)
	if expected != actual {
		t.Fatalf("TestPushTail = %s, want %s", actual, expected)
	}
}

func TestPushTailOfNil(t *testing.T) {
	var l *List[int]
	l = l.PushTail(4)
	expected := "4"
	actual := String(l)
	if expected != actual {
		t.Fatalf("TestPushTailOfNil = %s, want %s", actual, expected)
	}
}

func TestPopHead(t *testing.T) {
	l := New([]int{1, 2, 3})
	l, actualValue := l.PopHead()
	expectedValue := 1
	if expectedValue != actualValue {
		t.Fatalf("TestPopHead = %d, want %d", actualValue, expectedValue)
	}
	expectedList := "2->3"
	actualList := String(l)
	if expectedList != actualList {
		t.Fatalf("TestPopHead = %s, want %s", actualList, expectedList)
	}
}

func TestPopTail(t *testing.T) {
	l := New([]int{1, 2, 3})
	l, actualValue := l.PopTail()
	expectedValue := 3
	if expectedValue != actualValue {
		t.Fatalf("TestPopTail = %d, want %d", actualValue, expectedValue)
	}
	expectedList := "1->2"
	actualList := String(l)
	if expectedList != actualList {
		t.Fatalf("TestPopTail = %s, want %s", actualList, expectedList)
	}
}

func TestArray(t *testing.T) {
	l := New([]int{1, 2, 3})
	expected := []int{1, 2, 3}
	actual := l.Array()
	if len(expected) != len(actual) {
		t.Fatalf("TestArray len = %d, want %d", len(actual), len(expected))
	}
}

func TestReverse(t *testing.T) {
	l := New([]int{1, 2, 3})
	l = l.Reverse()
	expected := "3->2->1"
	actual := String(l)
	if expected != actual {
		t.Fatalf("TestReverse = %s, want %s", actual, expected)
	}
}

func TestReverseInPlace(t *testing.T) {
	l := New([]int{1, 2, 3, 4, 5})
	l = l.ReverseInPlace()
	expected := "5->4->3->2->1"
	actual := String(l)
	if expected != actual {
		t.Fatalf("TestReverseInPlace = %s, want %s", actual, expected)
	}
}

func TestReverseInPlaceOneValue(t *testing.T) {
	l := New([]int{1})
	l = l.ReverseInPlace()
	expected := "1"
	actual := String(l)
	if expected != actual {
		t.Fatalf("TestReverseInPlaceOneValue = %s, want %s", actual, expected)
	}
}

func TestReverseInPlaceTwoValues(t *testing.T) {
	l := New([]int{1, 2})
	l = l.ReverseInPlace()
	expected := "2->1"
	actual := String(l)
	if expected != actual {
		t.Fatalf("TestReverseInPlaceTwoValues = %s, want %s", actual, expected)
	}
}
