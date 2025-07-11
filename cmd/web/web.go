// web.go — простой веб-сервер с HTML-интерфейсом калькулятора
package main

import (
	"auth/internal/calculator"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type PageData struct {
	Expression string
	Result     string
	Error      string
}

func main() {
	http.HandleFunc("/", handleForm)
	http.ListenAndServe(":8080", nil)
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	data := PageData{}

	if r.Method == http.MethodPost {
		expr := r.FormValue("expr")
		data.Expression = expr

		parts := strings.Split(expr, " ")
		if len(parts) != 3 {
			data.Error = "Ошибка: введите выражение вида '2 + 2' или 'X * IV'"
			render(w, data)
			return
		}

		left, op, right := parts[0], parts[1], parts[2]

		if !strings.Contains("+-*/", op) {
			data.Error = "Ошибка: допустимы операции + - * /"
			render(w, data)
			return
		}

		leftIsRoman := calculator.IsValidRoman(left)
		rightIsRoman := calculator.IsValidRoman(right)

		if leftIsRoman != rightIsRoman {
			data.Error = "Ошибка: нельзя смешивать римские и арабские числа"
			render(w, data)
			return
		}

		var num1, num2, result int

		if leftIsRoman {
			num1 = calculator.RomanToInt(left)
			num2 = calculator.RomanToInt(right)
		} else {
			n1, err1 := strconv.Atoi(left)
			n2, err2 := strconv.Atoi(right)
			if err1 != nil || err2 != nil {
				data.Error = "Ошибка: не удалось распознать числа"
				render(w, data)
				return
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
				data.Error = "Ошибка: деление на ноль"
				render(w, data)
				return
			}
			result = num1 / num2
		}

		if leftIsRoman {
			if result <= 0 {
				data.Error = "Ошибка: римские числа не могут быть ≤ 0"
				render(w, data)
				return
			}
			data.Result = calculator.IntToRoman(result)
		} else {
			data.Result = strconv.Itoa(result)
		}
	}

	render(w, data)
}

func render(w http.ResponseWriter, data PageData) {
	tmpl := `
<!DOCTYPE html>
<html>
<head><meta charset="UTF-8"><title>Калькулятор</title></head>
<body>
	<h1>Калькулятор (римские / арабские)</h1>
	<form method="POST">
		<input type="text" name="expr" placeholder="например: X + V или 1 + 5 " value="{{.Expression}}" size="30" />
		<input type="submit" value="Вычислить">
	</form>
	{{if .Result}}<p><strong>Результат:</strong> {{.Result}}</p>{{end}}
	{{if .Error}}<p style="color:red;"><strong>{{.Error}}</strong></p>{{end}}
</body>
</html>
`
	t := template.Must(template.New("page").Parse(tmpl))
	t.Execute(w, data)
}
