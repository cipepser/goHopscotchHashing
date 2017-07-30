package main

import (
	"fmt"

	"./hopscotch"
)

func main() {
	h := hopscotch.NewHopscotch()

	h.Insert(1)
	h.Insert(1)
	// fmt.Println(len(h))
	for _, b := range h {
		fmt.Println(b)
	}

}
