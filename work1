package main

import "fmt"

func main() {
	var isBigFarm bool
	isBigFarm = true
	cowsCount := 22
	sheepsCount := 10
	chickensCount := 5
	var isCowHeavy bool
	var isSheepHeavy bool
	var isChickenHeavy bool
	var x string
	var y string
	x = "Hi, I'm Tosuncuk"
	y = "from Farm Bank."
	var barn int
	var cowAverageKg int
	var sheepAverageKg int
	var chickenAverageKg int
	var cowsAverageTotalKg int
	var sheepsAverageTotalKg int
	var chickensAverageTotalKg int
	var animalsTotalKg int
	var cowsCountTotal int
	var sheepsCountTotal int
	var chickensCountTotal int
	isCowHeavy = true
	isSheepHeavy = false
	isChickenHeavy = false

	if isBigFarm {
		barn = 5
	} else {
		barn = 3
	}
	if isCowHeavy {
		cowAverageKg = 500
	} else {
		cowAverageKg = 300
	}
	if isSheepHeavy {
		sheepAverageKg = 35
	} else {
		sheepAverageKg = 18
	}
	if isChickenHeavy {
		chickenAverageKg = 2
	} else {
		chickenAverageKg = 1
	}

	cowsAverageTotalKg = animalTotalKgFunc(cowsCount, cowAverageKg)
	sheepsAverageTotalKg = animalTotalKgFunc(sheepsCount, sheepAverageKg)
	chickensAverageTotalKg = animalTotalKgFunc(chickensCount, chickenAverageKg)
	animalsTotalKg = animalsTotalKgFunc(cowsAverageTotalKg, sheepsAverageTotalKg, chickensAverageTotalKg, barn)
	cowsCountTotal = animalsCountTotalFunc(cowsCount, barn)
	sheepsCountTotal = animalsCountTotalFunc(sheepsCount, barn)
	chickensCountTotal = animalsCountTotalFunc(chickensCount, barn)
	fmt.Println(x, y)
	fmt.Println(barn, "barn")
	fmt.Println(cowsCountTotal, "cowsCountTotal")
	fmt.Println(cowAverageKg, "cowAverageKg")
	fmt.Println(cowsAverageTotalKg*barn, "cowsAverageTotalKg")
	fmt.Println(sheepsCountTotal, "sheepsCountTotal")
	fmt.Println(sheepAverageKg, "sheepAverageKg")
	fmt.Println(sheepsAverageTotalKg*barn, "sheepsAverageTotalKg")
	fmt.Println(chickensCountTotal, "chickensCountTotal")
	fmt.Println(chickenAverageKg, "chickenAverageKg")
	fmt.Println(chickensAverageTotalKg*barn, "chickensAverageTotalKg")
	fmt.Println(animalsTotalKg, "animalsTotalKg")
//
}
func animalTotalKgFunc(animals int, animalsAverageKg int) int {
	var animalTotalKg int
	animalTotalKg = animals * animalsAverageKg

	return animalTotalKg
}

func animalsTotalKgFunc(cowsAverageTotalKg int, sheepsAverageTotalKg int, chickensAverageTotalKg int, barn int) int {
	var animalsTotalKg int

	animalsTotalKg = (cowsAverageTotalKg + sheepsAverageTotalKg + chickensAverageTotalKg) * (barn)
	return animalsTotalKg

}
func animalsCountTotalFunc(animalsCount int, barn int) int {
	var animalsCountTotal int
	animalsCountTotal = animalsCount * barn

	return animalsCountTotal
}
