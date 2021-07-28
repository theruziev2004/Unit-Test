package main

import "testing"

func TestArithmeticMean(t *testing.T) {
	// Настройка тестовых данных
	testTable := []struct{
		numbers []int
		expected float64
	}{
	{
		numbers: []int {1, 2, 3, 4, 5},
		expected: 3.0,
	},
	{
		numbers: []int {6, 7, 8, 9, 10},
		expected: 8.0,
	},
	{
		numbers: []int {54, 45, 85, 48, 78},
		expected: 62.0,
	},
	}


	// Вызов тестируемого кода
	for _, testCase := range testTable{
		result := ArithmeticMean(testCase.numbers)

		// Проверка возвращаемых результатов
		if result != testCase.expected {
			t.Errorf("Incorrect result")
		}
	}



}
