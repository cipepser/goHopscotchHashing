package hopscotch

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"math/big"
)

var (
	N int64 = 10 // テーブルの大きさ
)

const (
	H = 2 // bitmapのサイズ
)

func hash(key int64) int64 {
	hasher := md5.New()

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(key))
	hasher.Write(b)

	h := hex.EncodeToString(hasher.Sum(nil))

	t, _ := new(big.Int).SetString(h, 16)

	return t.Rem(t, big.NewInt(N)).Int64()
}

type bucket struct {
	item   int64
	bitmap [H]bool
}

type Hopscotch []bucket

func NewHopscotch() Hopscotch {
	h := Hopscotch(make([]bucket, N))
	return h
}

func (h Hopscotch) Reconstruct() Hopscotch {
	N = N * 2
	nh := Hopscotch(make([]bucket, N))

	for _, b := range h {
		if b.item != 0 {
			err := nh.Insert(b.item)
			for err != nil {
				nh = nh.Reconstruct()
				err = nh.Insert(b.item)
			}
		}
	}

	return nh
}

func (h Hopscotch) Lookup(key int64) bool {
	idx := int(hash(key))

	for i := 0; i < H; i++ {
		if h[(idx+i)%int(N)].item == key {
			return true
		}
	}

	return false
}

func (h Hopscotch) Insert(key int64) error {
	idx := int(hash(key))
	if h[idx].item == 0 {
		h[idx].item = key
		h[idx].bitmap[0] = true
		return nil
	}

	// linear probing for an empty backet
	i := idx + 1
	for h[i%int(N)].item != 0 {
		i++
		if i%int(N) == idx-1 {
			return errors.New("no empty bucket, you have to reconstruct backets with larger table size more than N.")
		}
	}

	// be able to insert the key within H-1
	if i < idx+H-1 {
		h[i%int(N)].item = key
		h[idx].bitmap[i-idx] = true
		return nil
	}

	// back to an empty backet util encounts the index within H-1 from the idx
	j := i - H + 1

	for i > idx+H-1 {
		k := 0
		for l, b := range h[j%int(N)].bitmap {
			if b {
				k = l
				h[j%int(N)].bitmap[k], h[j%int(N)].bitmap[H-1] = h[j%int(N)].bitmap[H-1], h[j%int(N)].bitmap[k]
				h[(j+k)%int(N)].item, h[(j+H-1)%int(N)].item = h[(j+H-1)%int(N)].item, h[(j+k)%int(N)].item
				break
			}

			if l == H-1 {
				return errors.New("no swapable bucket, you have to reconstruct backets with larger table size more than N.")
			}
		}
		i = j + k
		j = i - H + 1
	}

	h[i%int(N)].item = key
	h[idx].bitmap[i-idx] = true

	return nil
}

func (h Hopscotch) Delete(key int64) {
	idx := int(hash(key))

	for i := 0; i < H; i++ {
		if h[(idx+i)%int(N)].item == key {
			h[(idx+i)%int(N)].item = 0
			h[idx].bitmap[i] = false
		}
	}

	return
}
