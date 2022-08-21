package main

import (
	"errors"
	"fmt"
)

func main() {
	var housesM2 [8]float32 = [8]float32{200, 400, 150, 62, 98, 26.9, 700, 3000}
	var parquetWidth float32 = 0.15
	var parquetlength float32 = 1.2
	var parquetM2 float32 = 0
	var house = 0
	var err error
	var totalNumberOfParquetToBeMade float32
	var totalEarnings float32
	var missedJobPrice float32
	parquetM2func(parquetWidth, parquetlength, &parquetM2)

	for i := 0; i < len(housesM2); i++ {
		house++
		var numberOfFlooringToBeUsedForThehouse float32
		numberOfFlooringToBeUsedForThehouse = housesM2[i] / parquetM2
		housesM2[i] = numberOfFlooringToBeUsedForThehouse
		fmt.Println("house :", house, "Number of flooring to be used for the house :", numberOfFlooringToBeUsedForThehouse)
		switch {

		case housesM2[i] <= 5000 && housesM2[i] >= 150:
			totalNumberOfParquetToBeMade += housesM2[i]
			housesM2[i] = housesM2[i] * 10
			totalEarnings += housesM2[i]
			fmt.Println("house parquet fee", housesM2[i])
		case housesM2[i] <= 10000 && housesM2[i] > 5000:
			totalNumberOfParquetToBeMade += housesM2[i]
			housesM2[i] = (housesM2[i] * 10) * 0.9
			totalEarnings += housesM2[i]
			fmt.Println("house parquet fee", housesM2[i])
		case housesM2[i] <= 15000 && housesM2[i] > 10000:
			totalNumberOfParquetToBeMade += housesM2[i]
			housesM2[i] = (housesM2[i] * 10) * 0.8
			totalEarnings += housesM2[i]
			fmt.Println("house parquet fee", housesM2[i])
		case 15000 < housesM2[i]:
			missedJobPrice += housesM2[i]
			fallthrough

		case 150 > housesM2[i]:
			switch {
			case 150 > housesM2[i]:
				missedJobPrice += housesM2[i]
			}
			fallthrough

		default:
			err = errors.New("the number of parquet is not suitable for the workplace capacity")

			switch err {

			case nil:

			default:
				fmt.Println(err)
			}
		}
	}
	fmt.Println("totalNumberOfParquetToBeMade", totalNumberOfParquetToBeMade)
	m2ToBeDoneInTotalfunc(&totalNumberOfParquetToBeMade, parquetM2)
	fmt.Println("m2ToBeDoneInTotal", totalNumberOfParquetToBeMade, "M2")
	fmt.Println("missedJobPrice", missedJobPrice, "$")
	fmt.Println("totalEarnings", totalEarnings, "$")

}
func parquetM2func(parquetWidth float32, parquetlength float32, parquetM2 *float32) {

	*parquetM2 = parquetWidth * parquetlength

}
func m2ToBeDoneInTotalfunc(totalNumberOfParquetToBeMade *float32, parquetM2 float32) {

	*totalNumberOfParquetToBeMade = *totalNumberOfParquetToBeMade * parquetM2
}
