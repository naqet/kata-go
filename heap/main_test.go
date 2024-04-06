package heap_test

import (
	"errors"
	"reflect"
	"testing"
)

type Heap struct {
	length int
	data   []int
}

func NewHeap() *Heap {
    return &Heap{}
}

func (h *Heap) insert(val int) {
	h.data = append(h.data, val)
	h.heapifyUp(len(h.data) - 1)
	h.length++
}

func (h *Heap) delete() (int, error) {
	if h.length == 0 {
		return 0, errors.New("Heap is empty");
	}

    if h.length == 1 {
        head := h.data[0];
        h.data = []int{}
        h.length--;
        return head, nil;
    }

    head := h.data[0];
    last := h.data[h.length - 1];

    h.data[0] = last;

    h.data = h.data[:h.length - 1]

    h.length--;
    h.heapifyDown(0)

	return head, nil
}

func (h *Heap) parent(idx int) int {
	return (idx - 1) / 2
}

func (h *Heap) leftChild(idx int) int {
	return 2*idx + 1
}

func (h *Heap) rightChild(idx int) int {
	return 2*idx + 2
}

func (h *Heap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	val := h.data[idx]

	parentIdx := h.parent(idx)
	pval := h.data[parentIdx]

	if val < pval {
		h.data[idx] = pval
		h.data[parentIdx] = val
		h.heapifyUp(parentIdx)
	}
}

func (h *Heap) heapifyDown(idx int) {
	if idx >= h.length {
		return
	}

	lIdx := h.leftChild(idx)

	if lIdx >= h.length {
		return
	}

	val := h.data[idx]
	lVal := h.data[lIdx]
	rIdx := h.rightChild(idx)

	if rIdx >= h.length {
        if lVal < val {
            h.data[lIdx] = val
            h.data[idx] = lVal
        }

		return
	}

	rVal := h.data[rIdx]

	if rVal < lVal && rVal < val {
		h.data[rIdx] = val
		h.data[idx] = rVal
		h.heapifyDown(rIdx)
	} else if (lVal < rVal && lVal < val) || (lVal == rVal && lVal < val) {
		h.data[lIdx] = val
		h.data[idx] = lVal
		h.heapifyDown(lIdx)
    }
}

func TestHeap(t *testing.T) {
    heap := NewHeap();

    heap.insert(1)
    heap.insert(2)
    heap.insert(3)

    if !reflect.DeepEqual(heap.data, []int{1,2,3}) || heap.length != 3 {
        t.Error(heap.data, []int{1,2,3});
        return
    }

    heap.insert(4)
    deleted, err := heap.delete()

    if err != nil || deleted != 1 || heap.length != 3 || !reflect.DeepEqual(heap.data, []int{2,4,3}) {
        t.Error(heap.data, []int{2,4,3});
        return;
    }

    heap = NewHeap();

    heap.insert(1)
    heap.insert(2)
    heap.insert(3)
    heap.insert(4)
    heap.insert(5)
    deleted, err = heap.delete()

    if err != nil || deleted != 1 || heap.length != 4 || !reflect.DeepEqual(heap.data, []int{2,4,3,5}) {
        t.Error(heap.data, []int{2,4,3,5});
        return;
    }

    heap = NewHeap();

    heap.insert(1)
    heap.insert(1)
    heap.insert(1)
    heap.insert(4)
    heap.insert(5)
    deleted, err = heap.delete()

    if err != nil || deleted != 1 || heap.length != 4 || !reflect.DeepEqual(heap.data, []int{1,4,1,5}) {
        t.Error(heap.data, []int{2,4,3,5});
        return;
    }
}
