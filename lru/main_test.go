package lru_test

import "testing"

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

func NewNode[V any](val V) *Node[V] {
	return &Node[V]{value: val}
}

type LRU[K comparable, V any] struct {
	length        int
	capacity      int
	head          *Node[V]
	tail          *Node[V]
	lookup        map[K](*Node[V])
	reverseLookup map[*Node[V]]K
}

func NewLRU[K comparable, V any](capacity int) *LRU[K, V] {
	return &LRU[K, V]{
		length:        0,
		capacity:      capacity,
		lookup:        map[K](*Node[V]){},
		reverseLookup: map[*Node[V]]K{},
		head:          nil,
		tail:          nil,
	}
}

func (l *LRU[K, V]) update(key K, val V) {
	node, ok := l.lookup[key]

	if !ok {
		newNode := NewNode(val)
		l.prepend(newNode)
		l.trimCache()
        l.lookup[key] = newNode;
        l.reverseLookup[newNode] = key;
		return
	}

	l.detach(node)
	l.prepend(node)
    node.value = val;
}

func (l *LRU[K, V]) get(key K) (V, bool) {
	node, ok := l.lookup[key]
	var zeroVal V

	if !ok {
		return zeroVal, false
	}

	if l.head == nil {
		return zeroVal, false
	}

	l.detach(node)

	l.prepend(node)

	return node.value, true
}

func (l *LRU[K, V]) detach(node *Node[V]) {
	if node == nil {
		return
	}

	if l.head == node {
		l.head = node.next
	}

	if l.tail == node {
		l.tail = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
		node.prev = nil
	}

	if node.next != nil {
		node.next.prev = node.prev
		node.next = nil
	}

	l.length--
}

func (l *LRU[K, V]) prepend(node *Node[V]) {
	l.length++

	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	l.head.prev = node
	node.next = l.head
	l.head = node
}

func (l *LRU[K, V]) trimCache() {
	if l.length <= l.capacity {
		return
	}

	tail := l.tail

	l.detach(tail)

	key := l.reverseLookup[tail]

	delete(l.lookup, key)
	delete(l.reverseLookup, tail)
}

func Test(t *testing.T) {
    lru := NewLRU[string, int](3)

    _, ok := lru.get("foo")

    if ok {
        t.Error("Key before update")
        return;
    }

    lru.update("foo", 10)
    lru.update("bar", 12)
    lru.update("baz", 4)

    if lru.length != 3 {
        t.Error("Length")
        return;
    }

    val, ok := lru.get("foo")

    if !ok || val != 10 {
        t.Error("Value not valid")
        return
    }

    lru.update("hello", 123)

    _, ok = lru.get("bar")

    if ok {
        t.Error("Value should not be available")
        return
    }
}
