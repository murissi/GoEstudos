package main

import "fmt"

func main() {
	// Arrays
	var sites [4]string
	sites[0] = "Site01"
	sites[1] = "Site02"
	sites[2] = "Site03"
	sites[3] = "Site04"

	fmt.Println(sites)

	// Slices
	sitez := []string{"Site01", "Site02", "Site03", "Site04"}
	fmt.Println(sitez)

	// Arrays e Slices
	var eArray [2]int
	eArray[0] = 0
	eArray[1] = 1

	eslice := []int{1, 2, 3}
	fmt.Println(eslice)
	// Toda vez que usamos append = slice duplica de capcidade

	eslice = append(eslice, 4)
	fmt.Println(len(eslice))
	fmt.Println(cap(eslice))

}
