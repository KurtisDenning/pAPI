package localapp

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/oliverschweikert/pAPI/admin-app/data"
)

var (
	connection  data.MongoConnection
	collections []*data.MongoCollection
	a           fyne.App
)

func Start() {
	a = app.New()
	welcomeWindow().ShowAndRun()
}

func welcomeWindow() fyne.Window {
	w := a.NewWindow("pAPI Admin")
	w.Resize(fyne.NewSize(500, w.Canvas().Size().Height))
	content := container.NewVBox()
	title := widget.NewLabelWithStyle("Connect", fyne.TextAlignCenter, fyne.TextStyle{})
	loginButton := widget.NewButton("Connect", func() {
		checkPasswordWindow().Show()
		w.Close()
	})
	content.Add(title)
	content.Add(loginButton)
	w.SetContent(content)
	return w
}

func checkPasswordWindow() fyne.Window {
	c := a.NewWindow("Connect")
	c.Resize(fyne.NewSize(500, c.Content().MinSize().Height))
	content := container.NewVBox()
	errorMessage := canvas.NewText("Error connecting to the database", color.RGBA{255, 0, 0, 0})
	errorMessage.Alignment = fyne.TextAlignCenter
	errorMessage.Hidden = true
	content.Add(errorMessage)
	successMessage := canvas.NewText(fmt.Sprintf("Successfully connected to %v", data.DATABASE_NAME), color.RGBA{0, 255, 0, 0})
	successMessage.Alignment = fyne.TextAlignCenter
	successMessage.Hidden = true
	content.Add(successMessage)
	passwordField := widget.NewEntry()
	passwordField.Password = true
	form := widget.NewForm(widget.NewFormItem("Password", passwordField))
	content.Add(form)
	buttons := container.NewHBox()
	loginButton := widget.NewButton("Connect", func() {
		connection = data.NewConnection()
		err := connection.TestPassword(passwordField.Text)
		if err != nil {
			successMessage.Hidden = true
			errorMessage.Hidden = false
		} else {
			errorMessage.Hidden = true
			successMessage.Hidden = false
			content.Objects = content.Objects[:2]
			c.SetContent(content)
			time.Sleep(1 * time.Second)
			dbInfoWindow(a).Show()
			c.Close()
		}
	})
	cancelButton := widget.NewButton("Close", func() {
		welcomeWindow().Show()
		c.Close()
	})
	buttons.Add(loginButton)
	buttons.Add(cancelButton)
	content.Add(buttons)
	c.SetContent(content)
	return c
}

func dbInfoWindow(a fyne.App) fyne.Window {
	dbInfo := a.NewWindow("Database Information")
	content := container.NewVBox()
	header := canvas.NewText("Collections", color.White)
	header.TextSize = 32
	header.Alignment = fyne.TextAlignCenter
	content.Add(header)

	mCollections, err := connection.GetAllCollections()
	collections = mCollections
	if err != nil {
		errorLabel := widget.NewLabel("Error fetching collections - try restarting the app")
		errorLabel.Alignment = fyne.TextAlignCenter
		content.Add(errorLabel)
	} else {
		itemContainer := container.NewVBox()
		for _, c := range collections {
			c.Connection = &connection
			c.SetInfoButton(a)
			collectionRow := container.NewGridWithColumns(
				2,
				widget.NewLabel(c.Name), c.InfoButton)
			itemContainer.Add(collectionRow)
		}
		content.Add(itemContainer)
	}
	dbInfo.SetContent(content)
	return dbInfo
}
