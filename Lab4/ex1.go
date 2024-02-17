package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"math"
	"strconv"
)

func initGUI() {
	window := ui.NewWindow("Завдання №1", 400, 150, false)
	window.SetMargined(true)

	wLabel, hLabel, matLabel := ui.NewLabel("Ширина (см):"), ui.NewLabel("Висота (см):"), ui.NewLabel("Матеріал:")

	wInput, hInput := ui.NewEntry(), ui.NewEntry()

	matCmbox, typeCombobox := ui.NewCombobox(), ui.NewCombobox()

	matCmbox.Append("Дерево")
	matCmbox.Append("Метал")
	matCmbox.Append("Металопластик")
	typeCombobox.Append("Однокамерний")
	typeCombobox.Append("Двокамерний")

	matCmbox.SetSelected(0)
	typeCombobox.SetSelected(0)

	wBox, hBox, matBox := ui.NewHorizontalBox(), ui.NewHorizontalBox(), ui.NewHorizontalBox()

	wBox.SetPadded(true)
	hBox.SetPadded(true)
	matBox.SetPadded(true)

	wBox.Append(wLabel, true)
	wBox.Append(wInput, false)
	hBox.Append(hLabel, true)
	hBox.Append(hInput, false)
	matBox.Append(matLabel, true)
	matBox.Append(matCmbox, false)

	paramsBox, typeBox := ui.NewVerticalBox(), ui.NewVerticalBox()
	paramsLabel, typeLabel := ui.NewLabel("Розміри вікна"), ui.NewLabel("Склопакет")
	winsillCheckbox := ui.NewCheckbox("Підвіконня")

	paramsBox.SetPadded(true)
	typeBox.SetPadded(true)

	paramsBox.Append(paramsLabel, false)
	paramsBox.Append(wBox, false)
	paramsBox.Append(hBox, false)
	paramsBox.Append(matBox, false)
	typeBox.Append(typeLabel, false)
	typeBox.Append(typeCombobox, false)
	typeBox.Append(winsillCheckbox, true)

	mainParamsBox := ui.NewHorizontalBox()
	mainParamsBox.SetPadded(true)
	mainParamsBox.Append(paramsBox, false)
	mainParamsBox.Append(typeBox, true)

	resBox := ui.NewHorizontalBox()
	resBtn := ui.NewButton("Розрахувати")
	resultLabel := ui.NewLabel("")

	resBox.SetPadded(true)
	resBox.Append(resultLabel, true)

	winBox := ui.NewVerticalBox()
	winBox.SetPadded(true)
	winBox.Append(mainParamsBox, false)
	winBox.Append(resBox, false)
	winBox.Append(resBtn, true)

	window.SetChild(winBox)

	resBtn.OnClicked(func(*ui.Button) {
		var width, height, result float64
		var err error

		width, err = strconv.ParseFloat(wInput.Text(), 64)
		if err != nil || width <= 0 {
			resultLabel.SetText("Помилка. Перевірте поле Ширина (см)!")
			return
		}

		height, err = strconv.ParseFloat(hInput.Text(), 64)
		if err != nil || height <= 0 {
			resultLabel.SetText("Помилка. Перевірте поле Висота (см)!")
			return
		}

		result = getPrice(width, height, matCmbox.Selected(), typeCombobox.Selected(), winsillCheckbox.Checked())

		resultStr := strconv.FormatFloat(result, 'f', -1, 64)
		resultLabel.SetText("Ціна: " + resultStr + " грн.")
	})

	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	window.Show()
}

func getPrice(w float64, h float64, material int, windows int, winsill bool) float64 {

	res := w * h
	win := true
	if windows == 1 {
		win = false
	}
	switch material {
	case 0:
		if win {
			res *= 0.25
		} else {
			res *= 0.3
		}
		break
	case 1:
		if win {
			res *= 0.05
		} else {
			res *= 0.1
		}
		break
	case 2:
		if win {
			res *= 0.15
		} else {
			res *= 0.2
		}
		break
	}
	if winsill {
		res += 35
	}
	return math.Round(res)
}

func main() {
	err := ui.Main(initGUI)
	if err != nil {
		panic(err)
	}
}
