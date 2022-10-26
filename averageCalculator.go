package main

import "fmt"

func main() {
	var conclusion float32
	var lectureNotes float32
	var numberOfLessons int8
	var notes []float32
	fmt.Println("how many lesson  notes do you have")
	fmt.Scan(&numberOfLessons)
	for i := 0; int8(i) < numberOfLessons; i++ {
		fmt.Println("please enter your lesson  grade in order")
		fmt.Scan(&lectureNotes)
		notes = append(notes, lectureNotes)

	}
	conclusion = lessonsAverageCalculator(notes)
	fmt.Printf("your lesson  average %f", conclusion)
}

func lessonsAverageCalculator(array []float32) float32 {
	var total float32
	var i int
	for i = 0; i < len(array); i++ {
		total = total + array[i]

	}
	total = total / float32(i)
	return total

}
