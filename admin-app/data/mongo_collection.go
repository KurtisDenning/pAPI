package data

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MongoCollection struct {
	Name       string
	InfoButton *widget.Button
	Connection *MongoConnection
}

func NewMongoCollection(name string) *MongoCollection {
	mc := MongoCollection{
		Name: name}
	return &mc
}

func (c *MongoCollection) SetInfoButton(a fyne.App) {
	c.InfoButton = widget.NewButton("Info", func() { c.GetCollectionInfo(a) })
}

func (c *MongoCollection) GetCollectionInfo(a fyne.App) {
	if c.Name == "categories" {
		categories, err := c.Connection.GetCategories()
		if err != nil {
			fmt.Println(err)
		} else {
			c.NewCategoriesWindow(a, categories).Show()
		}
	} else if c.Name == "apiData" {
		apiDatas, err := c.Connection.GetAPIDatas()
		if err != nil {
			fmt.Println(err)
		} else {
			c.NewAPIDatasWindow(a, apiDatas).Show()

		}
	} else {
		fmt.Println(c.Name)
	}
}

func (c *MongoCollection) NewCategoriesWindow(a fyne.App, categories []*Category) fyne.Window {
	windows := a.Driver().AllWindows()
	for _, w := range windows {
		if w.Title() == c.Name {
			return w
		}
	}
	w := a.NewWindow(c.Name)
	content := container.NewVBox()
	header := canvas.NewText(c.Name, color.White)
	header.TextSize = 32
	header.Alignment = fyne.TextAlignCenter

	content.Add(header)
	w.SetContent(content)
	return w
}

func (c *MongoCollection) NewAPIDatasWindow(a fyne.App, apiDatas []*APIData) fyne.Window {
	windows := a.Driver().AllWindows()
	for _, w := range windows {
		if w.Title() == c.Name {
			return w
		}
	}
	w := a.NewWindow(c.Name)
	content := container.NewVBox()
	header := canvas.NewText(c.Name, color.White)
	header.TextSize = 32
	header.Alignment = fyne.TextAlignCenter
	content.Add(header)
	w.SetContent(content)
	return w
}
