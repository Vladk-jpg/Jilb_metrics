package main

import (
	"fmt"
	"strconv"
	"strings"
)

func extraArithmetic(a, b int) int {
	result := (a * b) + (a - b)
	result += 100
	result = result/2 + 7
	return result
}

func extraStringManipulation(s string) string {
	upper := strings.ToUpper(s)
	lower := strings.ToLower(upper)
	joined := upper + "-" + lower
	joined += strconv.Itoa(len(s) * 3)
	return joined
}

func main() {

	u := 1

	if u > 0 {
		if u > 0 {

		} else if u == 5 {

		}
	} else if u < 0 {
		if u < 5 {

		} else if u > 5 {

		} else if u == 8 {

		} else if u == 4 {

		} else {
			if u == 7 {

			}
		}
	} else if u == 0 {

	} else {
		if u == 5 {

		} else {
			if u == 8 {

			}
		}
	}

	for i := 0; i < 10; i++ {

	}

	x := 42
	if x > 0 {
		fmt.Println("x положительный")
	} else if x < 0 {
		if x == 0 {
			x = 0
		}
		fmt.Println("x отрицательный")
	} else {
		fmt.Println("x равен нулю")
	}

	for i := 0; i < 3; i++ {
		switch i {
		case 0:
			fmt.Println("Начало: i = 0")
		case 1:
			fmt.Println("Середина: i = 1")
		default:
			fmt.Println("Конец: i =", i)
		}
	}

	for j := 0; j < 5; j++ {
		if j%2 == 0 {
			fmt.Println("Четное j =", j)
		} else {
			fmt.Println("Нечетное j =", j)
		}
	}

	for a := 0; a < 2; a++ {
		for b := 0; b < 2; b++ {
			for c := 0; c < 2; c++ {
				for d := 0; d < 2; d++ {
					for e := 0; e < 2; e++ {
						for f := 0; f < 2; f++ {
							sum := a + b + c + d + e + f
							if sum > 3 {
								fmt.Println("Сумма значений вложенных переменных больше 3:", sum)
							}
							if a == 1 {
								fmt.Println("Переменная a равна 1")
							} else if b == 1 {
								fmt.Println("Переменная b равна 1")
							}

							switch {
							case c == 1 && d == 0:
								fmt.Println("c равна 1, а d равна 0")
							case e == 1:
								fmt.Println("e равна 1")
							default:
								fmt.Println("Другие комбинации: a, b, c, d, e, f =", a, b, c, d, e, f)
							}
						}
					}
				}
			}
		}
	}

	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			switch {
			case i < 5:
				fmt.Println("Меньше 5 и кратно 3: i =", i)
			case i >= 5:
				fmt.Println("Больше или равно 5 и кратно 3: i =", i)
			}
		} else {
			fmt.Println("i не кратно 3:", i)
		}
	}

	y := 25

	if y > 10 {
		switch {
		case y < 20:
			if y%2 == 0 {
				fmt.Println("y четный и находится между 10 и 20")
			} else if y%3 == 0 {
				fmt.Println("y делится на 3 и находится между 10 и 20")
			}
			if y > 15 {
				fmt.Println("Кроме того, y больше 15")
			}
		case y < 30:
			if y%2 != 0 {
				fmt.Println("y нечетный и находится между 20 и 30")
			} else if y%5 == 0 {
				fmt.Println("y делится на 5 и находится между 20 и 30")
			}
			if y%2 == 0 {

			}
		default:
			if y > 50 {
				fmt.Println("y очень велик")
			} else {
				fmt.Println("y находится между 30 и 50")
			}
		}
	}
	a := 10
	b := 20
	c := a + b
	d := c * 2
	e := d - b
	f := e / a
	fmt.Println("Результат ", f)

	arithmeticResult := extraArithmetic(x, y)
	fmt.Println("Результат extraArithmetic:", arithmeticResult)

	stringResult := extraStringManipulation("БессмысленныйТекст")
	fmt.Println("Результат extraStringManipulation:", stringResult)
}
