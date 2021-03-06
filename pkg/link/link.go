package link

import (
	"bytes"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Next  *Node
}

func New(lst ...int) *Node {
	head := &Node{}
	tmp := head
	for _, i := range lst {
		tmp.Next = &Node{Value: i}
		tmp = tmp.Next
	}
	return head.Next
}

// String >= StringV2 > StringV3
func (l *Node) String() string {
	if l == nil {
		return ""
	}

	var w bytes.Buffer
	w.WriteString(strconv.Itoa(l.Value))
	l = l.Next

	for l != nil {
		w.WriteString("->")
		w.WriteString(strconv.Itoa(l.Value))
		l = l.Next
	}
	return w.String()
}

func (l *Node) StringV2() string {
	if l == nil {
		return ""
	}

	var w bytes.Buffer
	for {
		w.WriteString(strconv.Itoa(l.Value))
		if l.Next != nil {
			w.WriteString("->")
			l = l.Next
			continue
		}
		break
	}
	return w.String()
}

func (l *Node) StringV3() string {
	lst := make([]string, 0)
	for l != nil {
		lst = append(lst, strconv.Itoa(l.Value))
		l = l.Next
	}
	return strings.Join(lst, "->")
}

func (l *Node) Reverse() *Node {
	var pre *Node
	cur := l
	for cur != nil {
		cur, cur.Next, pre = cur.Next, pre, cur
	}
	return pre
}

func (l *Node) Cmp(o *Node) bool {
	t1, t2 := l, o
	for t1 != nil && t2 != nil {
		if t1.Value != t2.Value {
			return false
		}
		t1 = t1.Next
		t2 = t2.Next
	}
	return t1 == nil && t2 == nil
}
