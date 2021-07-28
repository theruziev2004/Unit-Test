package main

import "fmt"

func main() {
	fmt.Println(ArithmeticMean([]int{4, 2, 4, 3, 5}))

}

func ArithmeticMean(numbers []int) float64 {
	sum := 0
	for _, score := range numbers{
		sum += score
	}
	var middle float64 = float64(sum) / float64(len(numbers))
	return middle
}


