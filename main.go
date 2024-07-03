package main

import "fmt"

func Bl(sum int, arr []int) [][]int {
	// Создаем массив для хранения результатов
	result := make([][]int, 0)

	// Функция рекурсивного поиска вариантов размена
	var recursive func(currentSum int, currentArr []int, start int)
	recursive = func(currentSum int, currentArr []int, start int) {
		// Если текущая сумма равна заданной сумме, добавляем текущий массив в результат
		if currentSum == sum {
			result = append(result, append([]int(nil), currentArr...))
			return
		}

		// Если текущая сумма больше заданной суммы, выходим из функции
		if currentSum > sum {
			return
		}

		// Проходимся по остальным элементам массива
		for i := start; i < len(arr); i++ {
			// Создаем временный массив для хранения текущего результата
			temp := append([]int(nil), currentArr...)

			// Добавляем текущий элемент к текущей сумме
			temp = append(temp, arr[i])

			// Рекурсивно вызываем функцию для следующего элемента
			recursive(currentSum+arr[i], temp, i)
		}
	}

	// Вызываем функцию рекурсивного поиска с начальными значениями
	recursive(0, []int{}, 0)

	return result
}

func main() {
	arr := []int{5000, 2000, 1000, 500, 200, 100, 50}
	sum := 400

	result := Bl(sum, arr)
	for _, arr := range result {
		fmt.Println(arr)

	}

}
