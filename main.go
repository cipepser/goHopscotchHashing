package main

import (
	"fmt"

	"./hopscotch"
)

func main() {
	h := hopscotch.NewHopscotch()
	// h.Insert(1)
	// h.Insert(2)
	// h.Insert(3)
	// h.Insert(4)
	// h.Insert(5)
	//
	keys := make([]int64, hopscotch.N)
	for k := 1; k <= int(hopscotch.N); k++ {
		keys[k-1] = int64(k)
	}
	// fmt.Println(keys)

	for _, k := range keys {
		err := h.Insert(k)
		if err != nil {
			// log.Fatal(err)
			h = h.Reconstruct()
		}

		for i, b := range h {
			fmt.Println(i, ":", b)
		}
		fmt.Println("~~~~~~~~~~~~")
	}

	fmt.Println(len(h))
	for i, b := range h {
		fmt.Println(i, ":", b)
	}

}
