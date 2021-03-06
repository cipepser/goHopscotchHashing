package hopscotch

import "testing"

func TestLookupSingleKey(t *testing.T) {
	h := NewHopscotch()

	err := h.Insert(1)
	for err != nil {
		h = h.Reconstruct()
		err = h.Insert(1)
	}

	isExists := h.Lookup(1)
	if !isExists {
		t.Error("1 is not found.")
	}
}

func TestLookupMultiKeys(t *testing.T) {
	h := NewHopscotch()

	keys := make([]int64, int(N))
	for k := 1; k <= int(N); k++ {
		keys[k-1] = int64(k)
	}

	for _, k := range keys {
		err := h.Insert(k)
		for err != nil {
			h = h.Reconstruct()
			err = h.Insert(k)
		}
	}
	for _, k := range keys {
		isExists := h.Lookup(k)
		if !isExists {
			t.Errorf("%d is not found.", k)
		}
	}
}

func TestLookupNonExistentKey(t *testing.T) {
	h := NewHopscotch()

	isExists := h.Lookup(999)
	if isExists {
		t.Error("999 is found.")
	}
}

func TestDeleteSingleKey(t *testing.T) {
	h := NewHopscotch()

	err := h.Insert(1)
	for err != nil {
		h = h.Reconstruct()
		err = h.Insert(1)
	}

	h.Delete(1)
	isExists := h.Lookup(1)
	if isExists {
		t.Error("1 is found.")
	}
}

func TestDeleteMultipleKeys(t *testing.T) {
	h := NewHopscotch()

	keys := make([]int64, N)
	for k := 1; k <= int(N); k++ {
		keys[k-1] = int64(k)
	}
	for _, k := range keys {
		err := h.Insert(k)
		if err != nil {
			h = h.Reconstruct()
		}
	}
	for _, k := range keys {
		h.Delete(k)
		isExists := h.Lookup(k)
		if isExists {
			t.Errorf("%d not found.", k)
		}
	}

}
