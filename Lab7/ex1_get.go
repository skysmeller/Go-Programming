package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

const (
	pageHeader = `<!DOCTYPE HTML>
 <html>
 <head>
 <title>Завдання №1</title>
 <style>
 .error{
 color:#FF0000;
 }
body {
 background: linear-gradient(to right, #70e1f5, #ffd194);
	font-size: 20px;
}
input {
	border: 0;
	border-radius: 5px;
	min-height: 15px;
	padding: 5px;
}
 </style>
 </head>`
	pageBody = `<body>
 <h1>Завдання 1</h1>
 <h2>Квадратне рівняння, GET</h2>`
	form = `<form action="/" method="GET">
 <label for="val_a">Введіть a:</label><br />
 <input type="text" name="val_a" size="20" required><br />

 <label for="val_b">Введіть b:</label><br />
 <input type="text" name="val_b" size="20" required><br />

 <label for="val_c">Введіть c:</label><br />
 <input type="text" name="val_c" size="20" required><br />
<br/>
 <input type="submit" value="Розрахувати">
 </form><br>`
	pageFooter = `</body></html>`

	anError = `<p class="error">%s</p>`
)

type Solution string

var HttpSolution1 Solution = "Завдання 1_Post"

func (s Solution) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, pageHeader, pageBody, form)
	if r.Method == "GET" {

		flgError := false
		a, errA := strconv.ParseFloat(r.FormValue("val_a"), 32) //отримуємо GET-параметри та перетворюємо у числа
		if errA != nil {
			fmt.Fprintf(w, "%v", "Помилка в значенні а<br>")
			flgError = true
		}
		b, errB := strconv.ParseFloat(r.FormValue("val_b"), 32)
		if errB != nil {
			fmt.Fprintf(w, "%v", "Помилка в значенні b<br>")
			flgError = true
		}
		c, errC := strconv.ParseFloat(r.FormValue("val_c"), 32)
		if errC != nil {
			fmt.Fprintf(w, "%v", "Помилка в значенні c<br>")
			flgError = true
		}
		if !flgError {
			D := b*b - 4*a*c
			var res string
			if D < 0 {
				res = "Рівняння немає дійсних коренів"
			}
			if D == 0 {
				x1 := -b / 2 * a
				res = "x1 = x2" + strconv.FormatFloat(x1, 'f', 2, 64)
			}
			if D > 0 {
				x1 := -b + math.Sqrt(D)/(2*a)
				x2 := -b - math.Sqrt(D)/(2*a)
				res = "x1 = " + strconv.FormatFloat(x1, 'f', 2, 64) + "<br>" + "x2 = " + strconv.FormatFloat(x2, 'f', 2, 64)
			}

			equ := "Рівняння: "
			equ += r.FormValue("val_a") + "x<sup>2</sup>"
			if b > 0 {
				equ += "+" + r.FormValue("val_b") + "x"

			} else {
				equ += r.FormValue("val_b") + "x"
			}
			if b > 0 {
				equ += "+" + r.FormValue("val_c")
			} else {
				equ += r.FormValue("val_c")
			}
			fmt.Fprintf(w, "%v", equ+" = 0"+"<br>")
			fmt.Fprintf(w, "%v", res)
			fmt.Fprintf(w, "\n\t<br /><br />")
		}
	}
	fmt.Fprint(w, "\n", pageFooter)
}
func main() {
	http.ListenAndServe("localhost:80", HttpSolution1)
}
