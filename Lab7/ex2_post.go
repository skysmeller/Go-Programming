package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
 <h1>Завдання 2, POST</h1>
 <h2>Запишіть числа через кому</h2>`
	form = `<form action="/" method="POST">
 <label for="numbers">Введіть числа</label><br />
 <input type="text" name="numbers" size="20" required><br />

 
<br/>
 <input type="submit" value="Розрахувати">
 </form><br>`
	pageFooter = `</body></html>`

	anError = `<p class="error">%s</p>`
)

type Solution string

var HttpSolution1 Solution = "Завдання 2_Post"

func (s Solution) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, pageHeader, pageBody, form)
	if r.Method == "POST" {
		err := r.ParseForm()
		post := r.PostForm
		if err != nil {
			fmt.Fprintf(w, anError, err)
			return
		}
		flgError := false
		sum_neg := 0.0
		dobutok := 1.0
		numbers := strings.Split(strings.Trim(post.Get("numbers"), " "), ",") //отримуємо string, обрізаємо пробіли і перетворюємо в масив
		for i := 0; i < len(numbers); i++ {
			num, err := strconv.ParseFloat(strings.Trim(numbers[i], " "), 64)
			if err != nil {
				flgError = true
			} else {
				dobutok *= num
				if num < 0 {
					sum_neg += num
				}

			}
		}
		if flgError {
			fmt.Fprintf(w, "%v", "Помилка в значенні, спробуйте ще раз. Наприклад 1,2,3,4,-5")
		} else {

			fmt.Fprintf(w, "%v%v", "Масив: "+post.Get("numbers"), "<br>")
			fmt.Fprintf(w, "%v%v%v", "Сума від'ємних елементів: ", sum_neg, "<br>")
			fmt.Fprintf(w, "%v%v%v", "Добуток елементів: ", dobutok, "<br>")
		}
		fmt.Fprintf(w, "\n\t<br /><br />")
	}

	fmt.Fprint(w, "\n", pageFooter)
}
func main() {
	http.ListenAndServe("localhost:80", HttpSolution1)
}
