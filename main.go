package main

import (
	"fmt"

	"./hopscotch"
)

func main() {
	h := hopscotch.NewHopscotch()

	for i := 0; i < 10; i++ {
		err := h.Insert(int64(i + 1))
		for err != nil {
			h = h.Reconstruct()
			err = h.Insert(int64(i + 1))
		}

		fmt.Println("-----------------------------")
		fmt.Println("No. \t| bucket")
		fmt.Println("-----------------------------")
		for j, b := range h {
			fmt.Println(j, "\t| ", b)
		}
		fmt.Println("")
	}
}
