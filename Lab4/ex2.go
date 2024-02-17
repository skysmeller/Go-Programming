package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"math"
	"strconv"
)

func initGUI() {
	window := ui.NewWindow("Завдання №2", 300, 150, false)
	window.SetMargined(true)

	daysLabel, countryLabel, timeLabel:= ui.NewLabel("К-сть днів:"),  ui.NewLabel("Країна:"), ui.NewLabel("Пора року:")

	daysInput := ui.NewEntry()

	countryComboBox, timeCombobox := ui.NewCombobox(), ui.NewCombobox()

	countryComboBox.Append("Болгарія")
	countryComboBox.Append("Німеччина")
	countryComboBox.Append("Польща")

	timeCombobox.Append("Літо")
	timeCombobox.Append("Зима")

	countryComboBox.SetSelected(0)
	timeCombobox.SetSelected(0)

	dayBox, countryBox, timeBox := ui.NewHorizontalBox(), ui.NewHorizontalBox(),ui.NewHorizontalBox()

	dayBox.SetPadded(true)

	countryBox.SetPadded(true)

	dayBox.Append(daysLabel, true)
	dayBox.Append(daysInput, false)

	countryBox.Append(countryLabel, true)
	countryBox.Append(countryComboBox, false)

	timeBox.Append(timeLabel, true)
	timeBox.Append(timeCombobox, false)


	paramsBox, typeBox := ui.NewVerticalBox(), ui.NewVerticalBox()

	gidCheckBox := ui.NewCheckbox("Індивідуальний гід")
	luxCheckbox := ui.NewCheckbox("Люкс номер")

	paramsBox.SetPadded(true)
	typeBox.SetPadded(true)


	paramsBox.Append(dayBox, false)
	paramsBox.Append(countryBox, false)

	typeBox.Append(timeBox, false)

	typeBox.Append(gidCheckBox, true)
	typeBox.Append(luxCheckbox, true)

	mainParamsBox := ui.NewVerticalBox()
	mainParamsBox.SetPadded(true)
	mainParamsBox.Append(paramsBox, false)
	mainParamsBox.Append(typeBox, true)


	resBtn := ui.NewButton("Розрахувати")

	resultLabel := ui.NewLabel("")

	winBox := ui.NewVerticalBox()
	winBox.SetPadded(true)
	winBox.Append(mainParamsBox, false)
	winBox.Append(resultLabel, true)
	winBox.Append(resBtn, true)

	window.SetChild(winBox)

	resBtn.OnClicked(func(*ui.Button) {
		var days, result float64
		var err error

		days, err = strconv.ParseFloat(daysInput.Text(), 64)
		if err != nil || days <= 0 {
			resultLabel.SetText("Помилка. Перевірте поле Кількість днів!")
			return
		}

		result = getPrice(days, countryComboBox.Selected(), timeCombobox.Selected(), gidCheckBox.Checked(), luxCheckbox.Checked())

		resultStr := strconv.FormatFloat(result, 'f', -1, 64)
		resultLabel.SetText("Ціна: " + resultStr + " $.")
	})

	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	window.Show()
}

func getPrice(days float64,  country int, season int, gid bool, lux bool) float64 {

	res := days

	time := true
	if season == 1 {
		time = false
	}
	switch country {
	case 0:
		if time {
			res *= 100
		} else {
			res *= 150
		}
		break
	case 1:
		if time {
			res *=160
		} else {
			res *= 200
		}
		break
	case 2:
		if time {
			res *= 120
		} else {
			res *= 180
		}
		break
	}

	if lux {
		res += res*0.2
	}
	if gid {
		res += 50
	}

	return math.Round(res)
}

func main() {
	err := ui.Main(initGUI)
	if err != nil {
		panic(err)
	}
}
