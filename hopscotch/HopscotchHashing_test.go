package hopscotch

import "testing"

func TestLookup(t *testing.T) {
	h := NewHopscotch()

	// single key
	h.Insert(1)
	isExists := h.Lookup(1)
	if !isExists {
		t.Error("1 is not found.")
	}

	// multiple keys
	// keys := []int64{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	keys := make([]int64, 40)
	for k := 1; k <= int(40); k++ {
		keys[k-1] = int64(k)
	}

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
	isExists = h.Lookup(999)
	if isExists {
		t.Error("999 is found.")
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
	// keys := []int64{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	keys := make([]int64, 40)
	for k := 1; k <= int(40); k++ {
		keys[k-1] = int64(k)
	}
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
