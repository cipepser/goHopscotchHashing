package hopscotch

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	h := NewHopscotch()

	h.Insert(1)
	fmt.Println(h)

}

func TestLookup(t *testing.T) {
	h := NewHopscotch()

	// single key
	h.Insert(1)
	isExists := h.Lookup(1)
	if !isExists {
		t.Error("1 is not found.")
	}

	// multiple keys
	keys := []int64{2, 3, 4, 5, 6, 7}
	for _, k := range keys {
		h.Insert(k)
	}
	for _, k := range keys {
		isExists := h.Lookup(k)
		if !isExists {
			t.Errorf("%d is not found.", k)
		}
	}

	// not exist key
	isExists = h.Lookup(0)
	if isExists {
		t.Error("0 is found.")
	}
}

func TestDelete(t *testing.T) {
	h := NewHopscotch()

	// single key
	h.Insert(1)
	h.Delete(1)
	isExists := h.Lookup(1)
	if isExists {
		t.Error("1 is found.")
	}

	// multiple keys
	keys := []int64{2, 3, 4, 5, 6, 7}
	for _, k := range keys {
		h.Insert(k)
	}
	for _, k := range keys {
		h.Delete(k)
		isExists := h.Lookup(k)
		if isExists {
			t.Errorf("%d not found.", k)
		}
	}

}
