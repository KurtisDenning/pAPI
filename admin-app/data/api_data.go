package data

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	URL        string                  `json:"request" bson:"request"`
	Response   *map[string]interface{} `json:"response,omitempty" bson:"response,omitempty"`
	LastUpdate *time.Time              `json:"lastUpdate,omitempty" bson:"lastUpdate,omitempty"`
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
	CategoriesC                               *APIDataCategoryCollection
	RequestsC                                 *APIDataRequestCollection
	TitleE, DescriptionE, ExternalURLE, BaseE *widget.Entry
	UpBut, DelBut                             *widget.Button
}

type APIDataCategoryItem struct {
	Entry  *widget.Entry
	Delete *widget.Button
}

type APIDataCategoryCollection struct {
	*APIDataItem
	Categories []*APIDataCategoryItem
	Accordion  *widget.Accordion
	NewBut     *widget.Button
}

type APIDataRequestItem struct {
	Entry  *widget.Entry
	Delete *widget.Button
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

func InitEmptyAPIDataItem(a *APIData) APIDataItem {
	var item APIDataItem
	item.APIData = a

	item.CategoriesC = &APIDataCategoryCollection{}
	item.CategoriesC.APIDataItem = &item
	item.CategoriesC.NewBut = widget.NewButton("Add Category", item.CategoriesC.AddCategory)
	item.CategoriesC.Categories = []*APIDataCategoryItem{}
	for _, c := range a.Categories {
		var category APIDataCategoryItem
		categoryEntry := widget.NewEntry()
		categoryEntry.Text = c.Hex()
		categoryDelete := widget.NewButton("Remove", category.RemoveCategory)
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
	for _, r := range a.Requests {
		var request APIDataRequestItem
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
	item.TitleE.Text = a.Title

	item.DescriptionE = widget.NewEntry()
	item.DescriptionE.Text = a.Description

	item.ExternalURLE = widget.NewEntry()
	item.ExternalURLE.Text = a.ExternalURL

	item.BaseE = widget.NewEntry()
	item.BaseE.Text = a.Base

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

func (a *APIDataCategoryCollection) AddCategory() {
	fmt.Println("Added a new category to the API")
}

func (a *APIDataCategoryItem) RemoveCategory() {
	fmt.Printf("Removed the category %v\n", a.Entry.Text)
}

func (a *APIDataRequestCollection) AddRequest() {
	fmt.Println("Added a new request to the API")
}

func (a *APIDataRequestItem) RemoveRequest() {
	fmt.Printf("Removed the request %v\n", a.Entry.Text)
}
