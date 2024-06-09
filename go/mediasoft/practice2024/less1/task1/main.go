package main

import "fmt"

// Условие задачи:
// Дан статический массив из 100 целых чисел.
// Написать одну из трёх сортировок статического массива (сортировку вставками, пузырьком, слиянием).

// Метод решения: Сортировка вставками

func main() {
	arr := [100]int{542, -565, 531, -294, -56, 14, 270, -51, -914, 605, -117, -768, 331, 708, -603, 84, -548, 579, 434, 751, 592, -349, 408, -602, 721, 909, 170, -432, -970, -171, -972, 316, 405, -676, -929, -795, -682, -646, 46, -609, -84, 180, -158, -662, -384, 854, -721, 39, 180, -197, -818, -946, -529, -555, -36, -853, -322, 540, -936, -919, 473, 978, 782, 586, 869, 333, -977, -548, -789, 988, -393, 807, -609, 997, 824, -480, -205, -576, 856, 494, 131, 40, -601, 467, 221, -640, 34, -220, 482, 948, 523, -27, -771, -914, 438, 957, 205, -411, -749, -723}

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}

	for i := 0; i < 100; i++ {
		fmt.Println(arr[i])
	}
}