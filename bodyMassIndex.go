package main

import (
	"errors"
	"fmt"
)

func main() {

	var classStudentNames [5]string = [5]string{"babür", "yusuf", "mücahit", "ahmet", "mehmet"}
	var classStudentWeights [5]float32 = [5]float32{85, 77, 76, 90, 110}
	var gradeStudentHeights [5]float32 = [5]float32{1.77, 1.8, 1.95, 1.5, 2.1}
	var bodyMassİndex float32
	var err error

	for i := 0; i < len(classStudentNames); i++ {

		var amountOfKilosToBeGained float32
		var amountOfWeightToLose float32
		var healthyWeightLevel float32
		bodyMassİndex, err = bodyMassİndexdeğerifunc(classStudentNames[i], classStudentWeights[i], gradeStudentHeights[i])

		if err != nil {
			fmt.Println(classStudentNames[i], err)
		}
		if bodyMassİndex <= 0 {
			continue
		} else if bodyMassİndex >= 18 && bodyMassİndex <= 25 {

			fmt.Println(classStudentNames[i], "body mass index normal", bodyMassİndex)

		} else if bodyMassİndex < 18 {

			healthyWeightLevel = 18 * (gradeStudentHeights[i] * gradeStudentHeights[i])

			amountOfKilosToBeGained = healthyWeightLevel - classStudentWeights[i]
			fmt.Println(classStudentNames[i], " weight to be gained : +", amountOfKilosToBeGained)
			fmt.Println("bodyMassİndex", bodyMassİndex)
		} else if bodyMassİndex > 25 {
			healthyWeightLevel := 25 * (gradeStudentHeights[i] * gradeStudentHeights[i])

			amountOfWeightToLose = classStudentWeights[i] - healthyWeightLevel
			fmt.Println(classStudentNames[i], "weight to be lost : -", amountOfWeightToLose)
			fmt.Println("bodyMassİndex", bodyMassİndex)
		}
		fmt.Println("")
	}

}

func bodyMassİndexdeğerifunc(classStudentNames string, classStudentWeights float32, gradeStudentHeights float32) (bodyMassİndexdeğeri float32, err error) {

	bodyMassİndex := classStudentWeights / (gradeStudentHeights * gradeStudentHeights)
	if bodyMassİndex <= 0 {
		err = errors.New("body mass index cannot be less than or equal to 0")
	}
	return bodyMassİndex, err
}
