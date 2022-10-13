package main

import "fmt"

func main() {
	var countS []int = []int{54, 44, 541, 560, 2, 65, 23}
	fmt.Println(countS)
	var countNew []int

	countNew = arrangement(countS)

	copy(countS, countNew)
	fmt.Println(countS)

}

func arrangement(slice []int) (slice2 []int) {
	var countBig int
	var smallCount int
	var index int
	for j := 0; j < len(slice); j++ {
		smallCount = slice[j]
		for i := 0; i < len(slice); i++ {

			if slice[i] < smallCount {

				smallCount = slice[i]
				index = i
			}

		}

		slice2 = append(slice2, smallCount)
		countBig = findLargeNumbers(slice)
		slice[index] = countBig
	}

	return slice2
}
func findLargeNumbers(slice []int) (countBig int) {
	countBig = slice[0]
	for i := 0; i < len(slice); i++ {
		if slice[i] > countBig {

			countBig = slice[i]

		}

	}

	return countBig
}

