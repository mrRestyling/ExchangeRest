package logic

import "exchange/internal/models"

func (a Appeal) Business(request *models.Request) [][]int {
	// Создаем массив для хранения результатов
	result := make([][]int, 0)

	// Функция рекурсивного поиска вариантов размена
	var recursive func(currentSum int, currentArr []int, start int)
	recursive = func(currentSum int, currentArr []int, start int) {
		// Если текущая сумма равна заданной сумме, добавляем текущий массив в результат
		if currentSum == request.Amount {
			result = append(result, append([]int(nil), currentArr...))
			return
		}

		// Если текущая сумма больше заданной суммы, выходим из функции
		if currentSum > request.Amount {
			return
		}

		// Проходимся по остальным элементам массива
		for i := start; i < len(request.Banknotes); i++ {
			// Создаем временный массив для хранения текущего результата
			temp := append([]int(nil), currentArr...)

			// Добавляем текущий элемент к текущей сумме
			temp = append(temp, request.Banknotes[i])

			// Рекурсивно вызываем функцию для следующего элемента
			recursive(currentSum+request.Banknotes[i], temp, i)
		}
	}

	// Вызываем функцию рекурсивного поиска с начальными значениями
	recursive(0, []int{}, 0)

	return result
}
