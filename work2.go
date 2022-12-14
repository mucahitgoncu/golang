package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var classStudents [3][10]string = [3][10]string{{"ali", "ahmet", "mehmet", "can", "cengiz", "polat", "memati", "abdulhey", "erhan", "süleyman"}, {"acun", "adem", "adnan", "arda", "alperen", "ata", "aykut", "ceyhun", "kılıç", "demir"}, {"batu", "bahadır", "mehmet", "barış", "savaş", "burak", "bülent", "yılmaz", "caner", "cüneyt"}}
	var classstudentNotes [3][10]int = [3][10]int{}
	for r := 0; r < len(classstudentNotes); r++ {
		for e := 0; e < len(classstudentNotes[e]); e++ {
			classstudentNotes[r][e] = rand.Intn(100)
		}
	}
	var err error
	var err2 error

	studentGradeEvaluation, err, err2 := studentGradeEvaluationFunc(classStudents, classstudentNotes)
	if err != nil {
		fmt.Println(err)
	} else {
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("Most successful class", studentGradeEvaluation)
		}
	}
}
func studentGradeEvaluationFunc(classStudents [3][10]string, studentNotes [3][10]int) (studentGradeEvaluation int, err error, err2 error) {
	var classNumberOfAchievements [3]int = [3]int{0, 0, 0}

	fmt.Println("Evaluation started")
	veryGoodStudents := 0
	veryGoodStudentsTotal := 0
	classNumber := 0

	for i := 0; i < len(classStudents); i++ {
		classNumber++
		veryGoodStudents = 0

		fmt.Println("class", classNumber)

		for j := 0; j < len(classStudents[0]); j++ {
			studentGrade := studentNotes[i][j]

			if 100 < studentGrade || studentGrade < 0 {
				fmt.Println(classStudents[i][j], studentGrade, "incorrect note")
			} else if 85 < studentGrade {
				veryGoodStudents++
				veryGoodStudentsTotal++
				fmt.Println(classStudents[i][j], studentGrade, "passed very good")

			} else if 85 >= studentGrade && studentGrade >= 50 {
				fmt.Println(classStudents[i][j], studentGrade, "passed the exam")

			} else {
				fmt.Println(classStudents[i][j], studentGrade, "failed the exam")

			}

		}
		fmt.Println("veryGoodStudents", veryGoodStudents)
		classNumberOfAchievements[i] = veryGoodStudents

	}
	fmt.Println("veryGoodStudentsTotal", veryGoodStudentsTotal)
	var successfulClass int
	successfulClass, err, err2 = mostSuccessfulClass(classNumberOfAchievements)
	fmt.Println("Evaluation is over")
	return successfulClass, err, err2
}
func mostSuccessfulClass(classNumberOfAchievements [3]int) (successfulClassno int, err error, err2 error) {
	var mostSuccessfulSlassNumberOfSuccessfulStudents int = 0
	var successfulClass = 0
	var successfulClass1 = 0
	var successfulClassNumberOfSuccess1 = 0
	var successfulClassNumberOfSuccess2 = 0
	for i := 0; i < len(classNumberOfAchievements); i++ {

		if mostSuccessfulSlassNumberOfSuccessfulStudents <= classNumberOfAchievements[i] {

			if mostSuccessfulSlassNumberOfSuccessfulStudents == classNumberOfAchievements[i] {

				successfulClassNumberOfSuccess1 = classNumberOfAchievements[i]
			}
			mostSuccessfulSlassNumberOfSuccessfulStudents = classNumberOfAchievements[i]

			successfulClassNumberOfSuccess2 = classNumberOfAchievements[i]
			successfulClass = i + 1
		}

	}
	if mostSuccessfulSlassNumberOfSuccessfulStudents == 0 {
		err = errors.New("No successful class")
	}
	successfulClass1, err2 = bestClassCalculationFunc(successfulClassNumberOfSuccess1, successfulClassNumberOfSuccess2)
	successfulClass1 = successfulClass
	return successfulClass1, err, err2
}

func bestClassCalculationFunc(successfulClassNumberOfSuccess1 int, successfulClassNumberOfSuccess2 int) (successfulClass1 int, err2 error) {
	successfulClass1 = 0
	var A int
	var B int

	A = successfulClassNumberOfSuccess2
	B = successfulClassNumberOfSuccess1
	if A > B {

		successfulClass1 = A
	} else {
		err2 = errors.New("There is no single most successful class")

	}

	return successfulClass1, err2

}

