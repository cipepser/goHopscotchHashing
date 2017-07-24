package main

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
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

// func (h *Hopscotch) Lookup(key int64) bool {
func (h Hopscotch) Lookup(key int64) bool {

	return false
}

func (h Hopscotch) Insert(key int64) {
	idx := int(hash(key))

	fmt.Println(idx)

	for i := idx; ; i++ {
		// if h[i%int(N)].item == 0 {
		if h[i].item == 0 {
			if i > idx+H-1 {
				for j := i - H + 1; ; j++ {
					for k, b := range h[j].bitmap {
						if b {
							h[i].item = h[j+k].item
						}
					}
				}
			}

			h[i%int(N)].item = key
			break
		}

	}

	// for i := 0; i < H; i++ {
	// 	if h[(idx+i)%int(N)].item == 0 {
	// 		h[(idx+i)%int(N)].item = key
	// 		break
	// 	}
	//
	//   if i == H - 1 {
	//
	//   }
	// }

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
