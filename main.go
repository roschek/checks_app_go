package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Приложение для чеков")
	w.Resize(fyne.NewSize(300, 500))

	title := widget.NewLabel("Оформление заказа")
	nameLabel := widget.NewLabel("Введите ваше имя")
	name := widget.NewEntry()
	foodLabel := widget.NewLabel("Выберите еду из меню")
	foods := widget.NewCheckGroup([]string{"Пицца", "Наггетсы", "Бургер", "Кола"}, func(s []string) {})
	placeLabel := widget.NewLabel("Заказ в ресторане или на вынос?")
	placeRadio := widget.NewRadioGroup([]string{"В ресторане", "На вынос"}, func(s string) {})
	result := widget.NewLabel("")

	btn := widget.NewButton("Оформить", func() {
		userName := name.Text
		foodArray := foods.Selected
		currentPlace := placeRadio.Selected
		result.SetText("Имя клиента: " + userName + "\nМесто: " + currentPlace + "\nЗаказ: ")
		for _, i := range foodArray {
			result.SetText(result.Text + i + ",")
		}
	})
	w.SetContent(container.NewVBox(
		title,
		nameLabel,
		name,
		foodLabel,
		foods,
		placeLabel,
		placeRadio,
		btn,
		result,
	))
	w.ShowAndRun()
}
