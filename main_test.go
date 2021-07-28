package main

import "testing"

func TestArithmeticMean(t *testing.T) {
	// Настройка тестовых данных
	numbers := []int {1, 2, 3, 4, 5}
	expected := 3.0

	// Вызов тестируемого кода
	result := ArithmeticMean(numbers)

	// Проверка возвращаемых результатов
	if result != expected {
		t.Errorf("Incorrect result")
	}
}






