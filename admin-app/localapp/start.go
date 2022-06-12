package localapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Start() {
	a := app.New()
	w := a.NewWindow("pAPI Admin")
	content := container.NewVBox()
	title := widget.NewLabel("Connect")
	loginButton := widget.NewButton("Connect", func() {
		checkPassword(a)
	})
	content.Add(title)
	content.Add(loginButton)
	w.SetContent(content)
	w.ShowAndRun()
}

func checkPassword(a fyne.App) {
	c := a.NewWindow("Connect")
	c.Resize(fyne.NewSize(500, c.Content().MinSize().Height))
	passwordField := widget.NewEntry()
	passwordField.Password = true
	content := container.NewVBox()
	form := widget.NewForm(widget.NewFormItem("Password", passwordField))
	content.Add(form)
	buttons := container.NewHBox()
	loginButton := widget.NewButton("Connect", func() {

	})
	cancelButton := widget.NewButton("Cancel", c.Close)
	buttons.Add(loginButton)
	buttons.Add(cancelButton)
	content.Add(buttons)
	c.SetContent(content)
	c.Show()
}
