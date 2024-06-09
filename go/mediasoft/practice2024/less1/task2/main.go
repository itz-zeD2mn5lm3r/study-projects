package main

import "fmt"

/*
Программа хранит данные об офисных сотрудниках. Поддерживаются следующие операции:
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
Максимальное количество сотрудников - 512

Написать программу которая хранит данные об офисных сотрудниках,
сделать операцию добавления и удаления

После каждого вызова fmt.Scanf, результат ошибки (err) проверяется.
Если ошибка не равна nil, программа выводит сообщение об ошибке и продолжает следующий шаг цикла.
Добавление обработки ошибок после каждого вызова fmt.Scanf сделано для исправления варнингов
(предупреждений) о необработанных ошибках.
*/

type Employee struct {
	Name     string // имя
	Age      int    // возраст
	Position string // позиция
	Salary   int    // зарплата
}

var commands = `
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
`

func main() {
	const size = 512
	empls := [size]*Employee{}
	var cmd int
	for {
		fmt.Print(commands)
		_, err := fmt.Scanf("%d\n", &cmd) // \n чтобы игнорировать оставшийся символ новой строки
		if err != nil {
			fmt.Println("Ошибка при вводе команды:", err)
			continue
		}

		switch cmd {
		case 1:
			// Добавляем нового сотрудника
			empl := new(Employee)
			fmt.Println("Имя:")
			_, err := fmt.Scanf("%s\n", &empl.Name)
			if err != nil {
				fmt.Println("Ошибка при вводе имени:", err)
				continue
			}
			fmt.Println("Возраст:")
			_, err = fmt.Scanf("%d\n", &empl.Age)
			if err != nil {
				fmt.Println("Ошибка при вводе возраста:", err)
				continue
			}
			fmt.Println("Позиция:")
			_, err = fmt.Scanf("%s\n", &empl.Position)
			if err != nil {
				fmt.Println("Ошибка при вводе позиции:", err)
				continue
			}
			fmt.Println("Зарплата:")
			_, err = fmt.Scanf("%d\n", &empl.Salary)
			if err != nil {
				fmt.Println("Ошибка при вводе зарплаты:", err)
				continue
			}

			added := false
			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					added = true
					fmt.Println("Сотрудник добавлен.")
					break
				}
			}
			if !added {
				fmt.Println("Ошибка: достигнут лимит сотрудников.")
			}

		case 2:
			// Удаляем сотрудника
			fmt.Println("Введите индекс сотрудника для удаления (0-511):")
			var index int
			_, err := fmt.Scanf("%d\n", &index)
			if err != nil {
				fmt.Println("Ошибка при вводе индекса:", err)
				continue
			}

			if index >= 0 && index < size && empls[index] != nil {
				empls[index] = nil
				fmt.Println("Сотрудник удален.")
			} else {
				fmt.Println("Ошибка: некорректный индекс или сотрудник не существует.")
			}

		case 3:
			// Выводим список сотрудников
			fmt.Println("Список сотрудников:")
			for i, empl := range empls {
				if empl != nil {
					fmt.Printf("Индекс: %d, Имя: %s, Возраст: %d, Позиция: %s, Зарплата: %d\n", i, empl.Name, empl.Age, empl.Position, empl.Salary)
				}
			}

		case 4:
			// Выход из программы
			fmt.Println("Выход из программы.")
			return

		default:
			fmt.Println("Ошибка: неизвестная команда.")
		}
	}
}
