package main

import "fmt"

func main() {

	data := make([]int, 0, 10000)
	lastCap := cap(data)
	count := 0

	for r := 0; r < 1000000; r++ {
		data = append(data, r)

		if lastCap != cap(data) {
			count++
			capCh := float64(cap(data)-lastCap) / float64(lastCap) * 100
			lastCap = cap(data)

			fmt.Printf("Add [%p] Cap[%d - %v]\n", data, cap(data), capCh)

		}

	}
	fmt.Println(count)
}
