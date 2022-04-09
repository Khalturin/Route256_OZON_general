package main

import "fmt"

/*
Маршруты

Условия
	Кирилл работает аналитиком в Ozon, и, недавно, ему в руки попал отчет, из которого он понял, что время доставки
	товаров в пункты выдачи можно значительно сократить.  Он заметил, что пункты выдачи в городе образуют выпуклый
	многоугольник, с количеством вершин, равным n, и располагаются на его вершинах, где одна вершина = один пункт
	выдачи.

	Кирилл решил воспользоваться всеми прелестями современных технологий, он хочет проложить воздушные пути между
	пунктами выдачи, по которым будут перемещаться курьеры на грузоподъемных реактивных ранцах.

	Можно выбрать какое-то число k, что каждый пункт выдачи будет соединен с k соседними пунктами выдачи слева и справа.
	Нужно найти минимальное k, чтобы кратчайшее расстояние между любыми двумя пунктами выдачи было меньше или равно r.
	Расстояние между пунктами выдачи примерно одинаковое, поэтому расстояние будет измеряться в количестве переходов,
	которые нужно сделать, чтобы попасть из одного пункта выдачи в другой.

Формат входных данных
	В первую строку вводится одно целое число P(1<=R<=100) - количество наборов входных данных.
	Для каждого набора входных данных в строку через пробел вводится два целых числа n и r.
	Ограничения: 2 < n < 10^18, 1 <= r < 10^7

Формат выходных данных
	Для каждого набора чисел выведите минимальное k, удовлетворяющее условию.
*/

// 3 1
// 4 1
// 5 2
// 6 2
// 7 3
// 8 3
// 9 3

func two(n, r []int, p int) (res []int) {
	for i := 0; i < p; i++ {
		val := (n[i] - 1) / 2
		if r[i] < val {
			val = r[i]
		}
		res = append(res, val)
	}
	return
}

func main() {
	p := 0
	fmt.Scan(&p)

	n := make([]int, p, p)
	r := make([]int, p, p)
	for i := 0; i < p; i++ {
		fmt.Scan(&n[i])
		fmt.Scan(&r[i])
	}

	for _, val := range two(n, r, p) {
		fmt.Println(val)
	}
}
