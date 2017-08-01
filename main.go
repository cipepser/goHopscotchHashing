package main

import (
	"fmt"
	"log"

	"./hopscotch"
)

func main() {
	h := hopscotch.NewHopscotch()

	keys := make([]int64, hopscotch.N)
	for k := 1; k <= int(hopscotch.N); k++ {
		keys[k-1] = int64(k)
	}
	fmt.Println(keys)

	for _, k := range keys {
		err := h.Insert(k)
		for i, b := range h {
			fmt.Println(i, ":", b)
		}
		fmt.Println("~~~~~~~~~~~~")
		if err != nil {
			log.Fatal(err)
		}
	}

	// fmt.Println(len(h))
	// for i, b := range h {
	// 	fmt.Println(i, ":", b)
	// }

}
