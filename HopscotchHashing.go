package main

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
)

var (
	N int64 = 10 // テーブルの大きさ
)

const (
	H = 4 // bitmapのサイズ
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

// func (h *Hopscotch) Lookup(key int64) bool {
func (h Hopscotch) Lookup(key int64) bool {
	idx := int(hash(key))

	for i := 0; i < H; i++ {
		if h[idx+i].item == key {
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
	for h[i].item != 0 {
		i++
		if i >= int(N) {
			return errors.New("no empty bucket, you have to reconstruct backets with larger N.")
		}
	}

	// be able to insert the key within H-1
	if i < idx+H-1 {
		h[i].item = key
		h[idx].bitmap[i-idx] = true
		return nil
	}

	// back an empty backet util it has the index within H-1 from idx
	j := i - H + 1
	for i > idx+H-1 {
		k := 0
		for l, b := range h[j].bitmap {
			if b {
				k = l
				h[j].bitmap[k], h[j].bitmap[H-1] = h[j].bitmap[H-1], h[j].bitmap[k]
				h[j+k].item, h[j+H-1].item = h[j+H-1].item, h[j+k].item
				break
			}
		}

		i = j + k
	}

	h[idx].bitmap[0], h[idx].bitmap[i-idx] = h[idx].bitmap[i-idx], h[idx].bitmap[0]
	h[idx].item, h[i].item = h[i].item, h[idx].item

	return nil
}

func (h Hopscotch) Delete(key int64) {
	return
}

func main() {
	h := NewHopscotch()

	keys := []int64{1, 2, 3, 4, 5, 6, 7}

	for _, k := range keys {
		h.Insert(k)
	}

	fmt.Println(h.Lookup(0))
	fmt.Println(h.Lookup(1))

	h.Delete(1)

	fmt.Println(h.Lookup(1))

	fmt.Println("---------------")
	// for _, k := range keys {
	// 	fmt.Println(hash(k))
	// }
	for _, v := range h {
		fmt.Println(v)
	}

}
