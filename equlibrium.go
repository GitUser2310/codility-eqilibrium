package main

import "fmt"

var numbers = []int{-7, 1, 5, 2, -4, 3, 0}

func calcSum(slice []int) (sum int) {

	for _, val := range slice {
		sum = sum + val
	}

	fmt.Print(" calSum=", sum)
	return
}

func main() {

	for pos := 0; pos < len(numbers); pos++ {
		if (pos != 0) && (pos != len(numbers)) {
			fmt.Print("pos=", pos, " ")
			if calcSum(numbers[0:pos]) == calcSum(numbers[(pos+1):len(numbers)]) {
				fmt.Println(" Equlibrium found!!!")
			} else {
				fmt.Println()
			}
		}
	}

}
