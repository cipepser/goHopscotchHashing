package main

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
)

const (
	N      int64 = 100000 // テーブルの大きさ
	bmSize       = 32     // bitmapのサイズ
)

func hash(key int64) (h1, h2 int64) {
	hasher := md5.New()

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(key))
	hasher.Write(b)

	h := hex.EncodeToString(hasher.Sum(nil))

	t1, _ := new(big.Int).SetString(h[:int(len(h)/2)], 16)
	t2, _ := new(big.Int).SetString(h[int(len(h)/2):], 16)

	h1 = t1.Rem(t1, big.NewInt(N)).Int64()
	h2 = t2.Rem(t2, big.NewInt(N)).Int64()

	return h1, h2
}

type Hopscotch struct {
	item   int64
	bitmap [bmSize]bool
}

func NewHopscotch() *Hopscotch {
	return new(Hopscotch)
}

func (h *Hopscotch) Lookup(key int64) bool {

	return false
}

func (h *Hopscotch) Insert(key int64) {
}

func (h *Hopscotch) Delete(key int64) {
	return
}

func main() {
	h := NewHopscotch()

	keys := []int64{1, 2, 3, 4, 5}

	for _, k := range keys {
		h.Insert(k)
	}

	fmt.Println(h.Lookup(0))
	fmt.Println(h.Lookup(1))

	h.Delete(1)

	fmt.Println(h.Lookup(1))
}
