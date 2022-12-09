package main

import (
	"image/color"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var CLIENT *http.Client

type RandomFact struct {
	Text string `json:"text"`
}

func main() {
	CLIENT = &http.Client{Timeout: 10 * time.Second}

	a := app.New()

	win := a.NewWindow("Get Useless Fact")
	win.Resize(fyne.NewSize(800, 300))

	title := canvas.NewText("Get Your Useless Facts", color.White)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	factText := widget.NewLabel("")
	factText.Wrapping = fyne.TextWrapWord

	button := widget.NewButton("Get Fact", func() {
		fact, err := GetRandomFact()
		if err != nil {
			dialog.ShowError(err, win)
		} else {
			factText.SetText(fact.Text)
		}
	})

	hBox := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		button,
		layout.NewSpacer(),
	)
	vBox := container.New(
		layout.NewVBoxLayout(),
		title,
		hBox,
		widget.NewSeparator(),
		factText,
	)

	win.SetContent(vBox)

	win.Show()
	a.Run()
}
