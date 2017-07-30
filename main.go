package main

import (
	"fmt"
	"log"

	"./hopscotch"
)

func main() {
	h := hopscotch.NewHopscotch()

	keys := make([]int64, 10)
	for k := 1; k <= int(10); k++ {
		keys[k-1] = int64(k)
	}
	fmt.Println(keys)

	for _, k := range keys {
		err := h.Insert(k)
		for _, b := range h {
			fmt.Println(b)
		}
		fmt.Println("-----------")
		if err != nil {
			log.Fatal(err)
		}
	}

	// fmt.Println(len(h))
	for _, b := range h {
		fmt.Println(b)
	}

}
