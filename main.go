package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFloat(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.Contains(input, ",") {
			fmt.Println("Ошибка: используйте точку для разделения десятичной части. Попробуйте ещё раз.")
			continue
		}

		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Некорректный ввод. Попробуйте ещё раз.")
			continue
		}
		return value
	}
}

func main() {
	fmt.Println("Приветствую! Я — программа, решающая квадратные уравнения ax^2 + bx + c = 0")
	fmt.Println()

	for {
		a := readFloat("Введи значение a: ")
		if a == 0 {
			fmt.Println("Значение a не может быть равно нулю, это уже не квадратное уравнение!")
			continue
		}

		b := readFloat("Введи значение b: ")
		c := readFloat("Введи значение c: ")

		fmt.Println()
		D := (b * b) - 4*(a*c)
		fmt.Println("Дискриминант равен:", D)

		if D > 0 {
			x1 := (-b + math.Sqrt(D)) / (2 * a)
			x2 := (-b - math.Sqrt(D)) / (2 * a)
			fmt.Println("Уравнение имеет два корня:", x1, "и", x2)
		} else if D == 0 {
			x := -b / (2 * a)
			fmt.Println("Уравнение имеет один корень:", x)
		} else {
			fmt.Println("Уравнение не имеет действительных корней")
		}

		fmt.Println()
		fmt.Print("Хотите продолжить? Введите 'да' или 'нет': ")
		var close string
		fmt.Scan(&close)

		if strings.ToLower(close) != "да" {
			fmt.Println("До новых встреч!")
			break
		}
		fmt.Println()
	}
}
