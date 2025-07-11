package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"auth/internal/calculator"
)

func main() {
	for {
		fmt.Println("\n<---Калькулятор--->")
		fmt.Println("1. Арабские числа")
		fmt.Println("2. Римские числа")
		fmt.Println("3. Выход")
		fmt.Println("4. Конвертер (например, 3 + 5 или IV * II)")

		var mode int
		fmt.Print("Ваш выбор: ")
		fmt.Scan(&mode)

		switch mode {
		case 1:
			fmt.Println("Вы выбрали режим: Арабские числа")

			var num1, num2 float64
			var operator string

			fmt.Print("Введите первое число: ")
			fmt.Scan(&num1)

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Введите операцию (+, -, *, /): ")
			opInput, _ := reader.ReadString('\n')
			operator = strings.TrimSpace(opInput)

			if len(operator) != 1 || !strings.Contains("+-*/", operator) {
				fmt.Println("Ошибка: введите только один символ операции (+, -, *, /)")
				continue
			}

			fmt.Print("Введите второе число: ")
			fmt.Scan(&num2)

			switch operator {
			case "+":
				fmt.Println("Результат:", num1+num2)
			case "-":
				fmt.Println("Результат:", num1-num2)
			case "*":
				fmt.Println("Результат:", num1*num2)
			case "/":
				if num2 == 0 {
					fmt.Println("Ошибка: деление на ноль")
					continue
				}
				fmt.Println("Результат:", num1/num2)
			}

		case 2:
			fmt.Println("Вы выбрали режим: Римские числа")

			var roman1, roman2 string
			var operator string
			var num1, num2, result int
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Введите первое римское число: ")
			roman1, _ = reader.ReadString('\n')
			roman1 = strings.TrimSpace(roman1)

			if !calculator.IsValidRoman(roman1) {
				fmt.Println("Ошибка: введено не римское число")
				continue
			}
			num1 = calculator.RomanToInt(roman1)

			fmt.Print("Введите операцию (+, -, *, /): ")
			opInput, _ := reader.ReadString('\n')
			operator = strings.TrimSpace(opInput)

			if !strings.Contains("+-*/", operator) {
				fmt.Println("Ошибка: операция должна быть одной из +, -, *, /")
				continue
			}

			fmt.Print("Введите второе римское число: ")
			roman2, _ = reader.ReadString('\n')
			roman2 = strings.TrimSpace(roman2)

			if !calculator.IsValidRoman(roman2) {
				fmt.Println("Ошибка: введено не римское число")
				continue
			}
			num2 = calculator.RomanToInt(roman2)

			switch operator {
			case "+":
				result = num1 + num2
			case "-":
				result = num1 - num2
			case "*":
				result = num1 * num2
			case "/":
				if num2 == 0 {
					fmt.Println("Ошибка: деление на ноль")
					continue
				}
				result = num1 / num2
			}

			if result <= 0 {
				fmt.Println("Ошибка: римские числа не могут быть ≤ 0")
				continue
			}

			fmt.Println("Результат:", result)
			fmt.Println("Результат (римский):", calculator.IntToRoman(result))

		case 3:
			fmt.Println("Выход из программы...")
			return

		case 4:
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Введите выражение (например, 3 + 5 или IV * II):")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			parts := strings.Split(input, " ")
			if len(parts) != 3 {
				fmt.Println("Ошибка: введите выражение в формате: число операция число")
				continue
			}

			left := parts[0]
			op := parts[1]
			right := parts[2]

			if !strings.Contains("+-*/", op) {
				fmt.Println("Ошибка: допустимые операции: +, -, *, /")
				continue
			}

			leftIsRoman := calculator.IsValidRoman(left)
			rightIsRoman := calculator.IsValidRoman(right)

			if leftIsRoman != rightIsRoman {
				fmt.Println("Ошибка: нельзя смешивать римские и арабские числа")
				continue
			}

			var num1, num2, result int
			if leftIsRoman {
				num1 = calculator.RomanToInt(left)
				num2 = calculator.RomanToInt(right)
			} else {
				n1, err1 := strconv.Atoi(left)
				n2, err2 := strconv.Atoi(right)
				if err1 != nil || err2 != nil {
					fmt.Println("Ошибка: не удалось распознать числа")
					continue
				}
				num1 = n1
				num2 = n2
			}

			switch op {
			case "+":
				result = num1 + num2
			case "-":
				result = num1 - num2
			case "*":
				result = num1 * num2
			case "/":
				if num2 == 0 {
					fmt.Println("Ошибка: деление на ноль")
					continue
				}
				result = num1 / num2
			}

			if leftIsRoman {
				if result <= 0 {
					fmt.Println("Ошибка: римские числа не могут быть ≤ 0")
					continue
				}
				fmt.Println("Результат:", calculator.IntToRoman(result))
			} else {
				fmt.Println("Результат:", result)
			}

		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}
