package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Приветствую! Я — программа, помогающая решить приведённые квадратные уравнения вида: ax^2+bx+c=0")
	fmt.Println(" ")
	for {
		var a float64
		var b float64
		var c float64
		var close string

		fmt.Print("Введи значение a: ")
		fmt.Scan(&a)

		if a == 0 {
			fmt.Println("Значение a не может быть равно нулю, это уже не квадратное уравнение! Попробуйте ещё раз:")
			continue
		}

		fmt.Print("Введи значение b: ")
		fmt.Scan(&b)

		fmt.Print("Введи значение c: ")
		fmt.Scan(&c)

		fmt.Println(" ")
		D := (b * b) - 4*(a*c)
		fmt.Println("Таким образом, дискриминант равен:", D)

		if D > 0 {

			var x1 float64
			var x2 float64

			x1 = (-b + math.Sqrt(D)) / (2 * a)
			x2 = (-b - math.Sqrt(D)) / (2 * a)

			fmt.Println("Ваше уравнение имеет два корня:", x1, "и", x2)
		}

		if D == 0 {

			var x float64

			x = (-b + math.Sqrt(D)) / (2 * a)

			fmt.Println("Ваше уравнение имеет один корень:", x)
		}

		if D < 0 {

			fmt.Println("Ваше уравнение не имеет корней, так как дискриминант меньше нуля")
		}

		fmt.Println(" ")
		fmt.Print("Хотите продолжить? Введите 'да' для продолжения или 'нет' для выхода: ")
		fmt.Scan(&close)

		if close != "да" && close != "Да" {
			fmt.Println("До новых встреч!")
			break
		}

		fmt.Println()
	}
}
