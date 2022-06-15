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

	header := canvas.NewText(c.Name, color.White)
	header.TextSize = 32
	header.Alignment = fyne.TextAlignCenter

	categoryAccordion := widget.NewAccordion()
	for _, cat := range categories {

		nameEntry := widget.NewEntry()
		nameEntry.Text = cat.Name
		shortDescEntry := widget.NewEntry()
		shortDescEntry.Text = cat.ShortDesc
		longDescEntry := widget.NewEntry()
		longDescEntry.Text = cat.LongDesc

		upButt := widget.NewButton("Update", func() {})
		delButt := widget.NewButton("Delete", func() {})

		cItem := CategoryItem{cat, c.Connection, nameEntry, shortDescEntry, longDescEntry, upButt, delButt}
		cItem.UpdateButt.OnTapped = func() { cItem.Update(a, w.Title()) }
		cItem.DeleteButt.OnTapped = func() {
			cItem.Delete(a, w.Title(), categories, c)
		}

		category := widget.NewAccordionItem(cat.ShortDesc, container.NewGridWithColumns(
			2,
			widget.NewLabel("ID"),
			widget.NewLabel(cat.Id.Hex()),
			widget.NewLabel("Name"),
			cItem.NameEntry,
			widget.NewLabel("Short Description"),
			cItem.ShortDescEntry,
			widget.NewLabel("Long Description"),
			cItem.LongDescEntry,
			cItem.UpdateButt,
			cItem.DeleteButt,
		))
		categoryAccordion.Append(category)
	}

	content := container.NewBorder(header, widget.NewButton("New Category", func() { c.NewCategoryWindow(a) }), nil, nil, categoryAccordion)

	w.SetContent(content)
	return w
}

func (c *MongoCollection) NewCategoryWindow(a fyne.App) {
	w := a.NewWindow("New Category")
	for _, window := range a.Driver().AllWindows() {
		if window.Title() == "New Category" {
			w = window
			break
		}
	}

	newCategoryItem := CategoryItem{
		&Category{},
		c.Connection,
		widget.NewEntry(),
		widget.NewEntry(),
		widget.NewEntry(),
		&widget.Button{},
		&widget.Button{},
	}

	header := widget.NewLabelWithStyle("New Category", fyne.TextAlignCenter, fyne.TextStyle{})

	footer := container.NewGridWithColumns(2)
	cancelButton := widget.NewButton("Cancel", w.Close)
	submitButton := widget.NewButton("Submit", func() {
		newCategoryItem.Create(w)
	})
	footer.Add(submitButton)
	footer.Add(cancelButton)

	formFields := container.NewGridWithColumns(2)
	formFields.Add(widget.NewLabel("Name"))
	formFields.Add(newCategoryItem.NameEntry)
	formFields.Add(widget.NewLabel("Short Description"))
	formFields.Add(newCategoryItem.ShortDescEntry)
	formFields.Add(widget.NewLabel("Long Description"))
	formFields.Add(newCategoryItem.LongDescEntry)

	content := container.NewBorder(header, footer, nil, nil, formFields)
	w.SetContent(content)
	w.Show()
}

func (c *MongoCollection) NewAPIDatasWindow(a fyne.App, apiDatas []*APIData) fyne.Window {
	windows := a.Driver().AllWindows()
	for _, w := range windows {
		if w.Title() == c.Name {
			return w
		}
	}
	w := a.NewWindow(c.Name)

	header := canvas.NewText(c.Name, color.White)
	header.TextSize = 32
	header.Alignment = fyne.TextAlignCenter

	acc := widget.NewAccordion()
	for _, a := range apiDatas {
		apiItem := InitEmptyAPIDataItem(a)

		accDetail := container.NewGridWithColumns(2)
		accDetail.Add(widget.NewLabel("ID"))
		idWidget := widget.NewEntry()
		idWidget.Text = apiItem.Id.Hex()
		accDetail.Add(idWidget)
		accDetail.Add(widget.NewLabel("Title"))
		accDetail.Add(apiItem.TitleE)
		accDetail.Add(widget.NewLabel("Description"))
		accDetail.Add(apiItem.DescriptionE)
		accDetail.Add(widget.NewLabel("ExternalURL"))
		accDetail.Add(apiItem.ExternalURLE)
		accDetail.Add(widget.NewLabel("BaseURL"))
		accDetail.Add(apiItem.BaseE)

		accFunc := container.NewGridWithColumns(2)
		accFunc.Add(apiItem.UpBut)
		accFunc.Add(apiItem.DelBut)

		accCollections := container.NewBorder(apiItem.CategoriesC.Accordion, apiItem.RequestsC.Accordion, nil, nil)

		accAll := container.NewBorder(accDetail, accFunc, nil, nil, accCollections)
		accItem := widget.NewAccordionItem(
			apiItem.Title,
			accAll,
		)
		acc.Items = append(acc.Items, accItem)
	}
	apiItems := container.NewScroll(acc)

	footer := widget.NewButton(
		"New API Data",
		CreateAPIData,
	)

	content := container.NewBorder(header, footer, nil, nil, apiItems)
	w.SetContent(content)
	w.Resize(w.Content().Size().AddWidthHeight(acc.Size().Width, acc.Size().Height))
	return w
}
