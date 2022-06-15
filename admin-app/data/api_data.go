package data

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	URL        string                  `json:"request" bson:"request"`
	Response   *map[string]interface{} `json:"response,omitempty" bson:"response,omitempty"`
	LastUpdate time.Time               `json:"lastUpdate,omitempty" bson:"lastUpdate,omitempty"`
}

type APIData struct {
	Id           primitive.ObjectID   `json:"_id" bson:"_id"`
	Title        string               `json:"title" bson:"title"`
	Description  string               `json:"description" bson:"description"`
	ExternalURL  string               `json:"externalURL" bson:"externalURL"`
	DailyCount   *int                 `json:"dailyCount" bson:"dailyCount"`
	WeeklyCount  *int                 `json:"weeklyCount" bson:"weeklyCount"`
	MonthlyCount *int                 `json:"monthlyCount" bson:"monthlyCount"`
	YearlyCount  *int                 `json:"yearlyCount" bson:"yearlyCount"`
	TotalCount   *int                 `json:"totalCount" bson:"totalCount"`
	Base         string               `json:"base" bson:"base"`
	Categories   []primitive.ObjectID `json:"categories" bson:"categories"`
	Requests     []*Request           `json:"requests" bson:"requests"`
}

type APIDataItem struct {
	*APIData
	App                                       fyne.App
	Window                                    fyne.Window
	Connection                                *MongoConnection
	CategoriesC                               *APIDataCategoryCollection
	RequestsC                                 *APIDataRequestCollection
	TitleE, DescriptionE, ExternalURLE, BaseE *widget.Entry
	UpBut, DelBut                             *widget.Button
}

type APIDataCategoryItem struct {
	Collection *APIDataCategoryCollection
	Entry      *widget.Entry
	Delete     *widget.Button
}

type APIDataCategoryCollection struct {
	*APIDataItem
	Categories []*APIDataCategoryItem
	Accordion  *widget.Accordion
	NewBut     *widget.Button
}

type APIDataRequestItem struct {
	Collection *APIDataRequestCollection
	Entry      *widget.Entry
	Delete     *widget.Button
}

type APIDataRequestCollection struct {
	*APIDataItem
	Requests  []*APIDataRequestItem
	Accordion *widget.Accordion
	NewBut    *widget.Button
}

func CreateAPIData() {
	fmt.Println("Creating a new APIData")
}

func InitEmptyAPIDataItem(app *APIData, connection *MongoConnection, a fyne.App, w fyne.Window) APIDataItem {
	var item APIDataItem
	item.APIData = app

	item.Connection = connection
	item.App = a

	item.CategoriesC = &APIDataCategoryCollection{}
	item.CategoriesC.APIDataItem = &item
	item.CategoriesC.NewBut = widget.NewButton("Add Category", item.CategoriesC.AddCategoryToAPI)
	item.CategoriesC.Categories = []*APIDataCategoryItem{}
	for _, c := range app.Categories {
		var category APIDataCategoryItem
		category.Collection = item.CategoriesC
		categoryEntry := widget.NewEntry()
		categoryEntry.Text = c.Hex()
		categoryDelete := widget.NewButton("Remove", category.RemoveCategoryFromAPI)
		category.Entry = categoryEntry
		category.Delete = categoryDelete
		item.CategoriesC.Categories = append(item.CategoriesC.Categories, &category)
	}
	item.CategoriesC.Accordion = widget.NewAccordion()
	categories := container.NewGridWithColumns(2)
	for _, c := range item.CategoriesC.Categories {
		categories.Add(c.Entry)
		categories.Add(c.Delete)
	}
	item.CategoriesC.Accordion = widget.NewAccordion(widget.NewAccordionItem("Categories", container.NewBorder(item.CategoriesC.NewBut, categories, nil, nil)))

	item.RequestsC = &APIDataRequestCollection{}
	item.RequestsC.APIDataItem = &item
	item.RequestsC.NewBut = widget.NewButton("Add Request", item.RequestsC.AddRequest)
	item.RequestsC.Requests = []*APIDataRequestItem{}
	for _, r := range app.Requests {
		var request APIDataRequestItem
		request.Collection = item.RequestsC
		requestEntry := widget.NewEntry()
		requestEntry.Text = r.URL
		requestDelete := widget.NewButton("Remove", request.RemoveRequest)
		request.Entry = requestEntry
		request.Delete = requestDelete
		item.RequestsC.Requests = append(item.RequestsC.Requests, &request)
	}
	item.RequestsC.Accordion = widget.NewAccordion()
	requests := container.NewGridWithColumns(2)
	for _, c := range item.RequestsC.Requests {
		requests.Add(c.Entry)
		requests.Add(c.Delete)
	}
	item.RequestsC.Accordion = widget.NewAccordion(widget.NewAccordionItem("Requests", container.NewBorder(item.RequestsC.NewBut, requests, nil, nil)))

	item.TitleE = widget.NewEntry()
	item.TitleE.Text = app.Title

	item.DescriptionE = widget.NewEntry()
	item.DescriptionE.Text = app.Description

	item.ExternalURLE = widget.NewEntry()
	item.ExternalURLE.Text = app.ExternalURL

	item.BaseE = widget.NewEntry()
	item.BaseE.Text = app.Base

	item.UpBut = widget.NewButton("Update", item.UpdateAPIData)

	item.DelBut = widget.NewButton("Delete", item.DeleteAPIData)

	return item
}

func (a *APIDataItem) UpdateAPIData() {
	fmt.Println("Update Item")
}

func (a *APIDataItem) DeleteAPIData() {
	fmt.Println("Deleted Item")
}

func (a *APIDataCategoryCollection) AddCategoryToAPI() {
	win := a.AddCategoryToAPIWindow()
	win.Show()
}

func (a *APIDataCategoryCollection) AddCategoryToAPIWindow() fyne.Window {
	for _, w := range a.App.Driver().AllWindows() {
		if w.Title() == fmt.Sprintf("%v | Add Category", a.Title) {
			return w
		}
	}
	win := a.App.NewWindow(fmt.Sprintf("%v | Add Category", a.Title))

	catLabel := widget.NewLabel("Category ID")
	catEntry := widget.NewEntry()
	body := container.NewGridWithColumns(1, catLabel, catEntry)

	acceptButton := widget.NewButton("Add", func() {
		err := a.Connection.AddCategoryToAPI(catEntry.Text, a.Id)
		var d dialog.Dialog
		if err != nil {
			d = dialog.NewInformation("Error", err.Error(), win)
		} else {
			d = dialog.NewInformation("Success", fmt.Sprintf("Added %v", catEntry.Text), win)
			d.SetOnClosed(func() {
				win.Close()
			})
		}
		d.Show()
	})
	closeButton := widget.NewButton("Cancel", func() { win.Close() })
	footer := container.NewGridWithColumns(2, acceptButton, closeButton)

	content := container.NewBorder(body, footer, nil, nil)
	win.SetContent(content)
	return win
}

func (a *APIDataCategoryItem) RemoveCategoryFromAPI() {
	win := a.RemoveCategoryFromAPIWindow()
	win.Show()
}

func (a *APIDataCategoryItem) RemoveCategoryFromAPIWindow() fyne.Window {
	for _, w := range a.Collection.App.Driver().AllWindows() {
		if w.Title() == fmt.Sprintf("%v | Remove Category", a.Entry.Text) {
			return w
		}
	}
	win := a.Collection.App.NewWindow(fmt.Sprintf("%v | Remove Category", a.Entry.Text))

	checkLabel := widget.NewLabel(fmt.Sprintf("Are you sure you want to delete %v?", a.Entry.Text))
	body := container.NewGridWithColumns(1, checkLabel)

	acceptButton := widget.NewButton("Delete", func() {
		err := a.Collection.Connection.RemoveCategoryFromAPI(a.Entry.Text, a.Collection.Id)
		var d dialog.Dialog
		if err != nil {
			d = dialog.NewInformation("Error", err.Error(), win)
		} else {
			d = dialog.NewInformation("Success", fmt.Sprintf("Removed %v", a.Entry.Text), win)
			d.SetOnClosed(func() {
				win.Close()
			})
		}
		d.Show()
	})
	closeButton := widget.NewButton("Cancel", func() { win.Close() })
	footer := container.NewGridWithColumns(2, acceptButton, closeButton)

	content := container.NewBorder(body, footer, nil, nil)
	win.SetContent(content)
	return win
}

func (a *APIDataRequestCollection) AddRequest() {
	win := a.NewAddRequestWindow()
	win.Show()
}

func (a *APIDataRequestCollection) NewAddRequestWindow() fyne.Window {
	for _, w := range a.App.Driver().AllWindows() {
		if w.Title() == fmt.Sprintf("%v | Add Request", a.Title) {
			return w
		}
	}
	win := a.App.NewWindow(fmt.Sprintf("%v | Add Request", a.Title))

	reqLabel := widget.NewLabel("Request URL")
	reqEntry := widget.NewEntry()
	body := container.NewGridWithColumns(1, reqLabel, reqEntry)

	acceptButton := widget.NewButton("Add", func() {
		err := a.Connection.AddRequestToAPI(reqEntry.Text, a.Id)
		var d dialog.Dialog
		if err != nil {
			d = dialog.NewInformation("Error", err.Error(), win)
		} else {
			d = dialog.NewInformation("Success", fmt.Sprintf("Added %v", reqEntry.Text), win)
			d.SetOnClosed(func() {
				win.Close()
			})
		}
		d.Show()
	})
	closeButton := widget.NewButton("Cancel", func() { win.Close() })
	footer := container.NewGridWithColumns(2, acceptButton, closeButton)

	content := container.NewBorder(body, footer, nil, nil)
	win.SetContent(content)
	return win
}

func (a *APIDataRequestItem) RemoveRequest() {
	win := a.RemoveRequestWindow()
	win.Show()
}

func (a *APIDataRequestItem) RemoveRequestWindow() fyne.Window {
	for _, w := range a.Collection.App.Driver().AllWindows() {
		if w.Title() == fmt.Sprintf("%v | Remove Request", a.Entry.Text) {
			return w
		}
	}
	win := a.Collection.App.NewWindow(fmt.Sprintf("%v | Remove Request", a.Entry.Text))

	checkLabel := widget.NewLabel(fmt.Sprintf("Are you sure you want to delete %v?", a.Entry.Text))
	body := container.NewGridWithColumns(1, checkLabel)

	acceptButton := widget.NewButton("Delete", func() {
		err := a.Collection.Connection.RemoveRequestFromAPI(a.Entry.Text, a.Collection.Id)
		var d dialog.Dialog
		if err != nil {
			d = dialog.NewInformation("Error", err.Error(), win)
		} else {
			d = dialog.NewInformation("Success", fmt.Sprintf("Removed %v", a.Entry.Text), win)
			d.SetOnClosed(func() {
				win.Close()
			})
		}
		d.Show()
	})
	closeButton := widget.NewButton("Cancel", func() { win.Close() })
	footer := container.NewGridWithColumns(2, acceptButton, closeButton)

	content := container.NewBorder(body, footer, nil, nil)
	win.SetContent(content)
	return win
}
