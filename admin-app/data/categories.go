package data

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	ShortDesc string             `json:"shortDesc" bson:"shortDesc"`
	LongDesc  string             `json:"longDesc" bson:"longDesc"`
}
type CategoryItem struct {
	*Category
	Connection                                        *MongoConnection
	IDEntry, NameEntry, ShortDescEntry, LongDescEntry *widget.Entry
	UpdateButt, DeleteButt                            *widget.Button
}

func (c *CategoryItem) Update(a fyne.App, windowTitle string) {
	var w fyne.Window
	for _, window := range a.Driver().AllWindows() {
		if window.Title() == windowTitle {
			w = window
			break
		}
	}

	c.Name = c.NameEntry.Text
	c.ShortDesc = c.ShortDescEntry.Text
	c.LongDesc = c.LongDescEntry.Text
	byteData, err := bson.Marshal(c.Category)
	if err != nil {
		d := dialog.NewInformation("Error", err.Error(), w)
		d.Show()
	} else {
		var bsonData bson.D
		err = bson.Unmarshal(byteData, &bsonData)
		if err != nil {
			d := dialog.NewInformation("Error", err.Error(), w)
			d.Show()
		} else {
			err = c.Connection.UpdateCategory(c.Id, bsonData)
			if err != nil {
				d := dialog.NewInformation("Error", err.Error(), w)
				d.Show()
			} else {
				d := dialog.NewInformation("Success", fmt.Sprintf("Successfully updated %v", c.Id.Hex()), w)
				d.Show()
			}
		}
	}
}

func (c *CategoryItem) Create(w fyne.Window) {
	c.Category = &Category{
		Id:        primitive.NewObjectID(),
		Name:      c.NameEntry.Text,
		ShortDesc: c.ShortDescEntry.Text,
		LongDesc:  c.LongDescEntry.Text,
	}

	byteData, err := bson.Marshal(c.Category)
	if err != nil {
		d := dialog.NewInformation("Error!", err.Error(), w)
		d.Show()
		d.SetOnClosed(func() {
			w.Close()
		})
	} else {
		var bsonData bson.D
		err := bson.Unmarshal(byteData, &bsonData)
		if err != nil {
			d := dialog.NewInformation("Error!", err.Error(), w)
			d.Show()
			d.SetOnClosed(func() {
				w.Close()
			})
		} else {
			fmt.Println(bsonData)
			err := c.Connection.AddCategory(bsonData)
			if err != nil {
				d := dialog.NewInformation("Error!", err.Error(), w)
				d.Show()
				d.SetOnClosed(func() {
					w.Close()
				})
			} else {
				d := dialog.NewInformation("Success!", fmt.Sprintf("Created %v", c.ShortDescEntry.Text), w)
				d.Show()
				d.SetOnClosed(func() {
					w.Close()
				})
			}
		}
	}

	d := dialog.NewInformation("Success!", fmt.Sprintf("Created %v", c.NameEntry.Text), w)
	d.Show()
	d.SetOnClosed(func() {
		w.Close()
	})
}

func (c *CategoryItem) Delete(a fyne.App, windowTitle string, categories []*Category, coll *MongoCollection) {
	var w fyne.Window
	for _, window := range a.Driver().AllWindows() {
		if window.Title() == windowTitle {
			w = window
			break
		}
	}
	d := dialog.NewConfirm("Delete", fmt.Sprintf("Are you sure you want to delete %v?", c.Name), func(b bool) {
		if b {
			err := c.Connection.DeleteCategory(c.Id)
			if err != nil {
				d := dialog.NewInformation("Error", err.Error(), w)
				d.Show()
			} else {
				fmt.Println("Deleted something!")
				categories, err := coll.Connection.GetCategories()
				if err != nil {
					fmt.Println(err)
				} else {
					coll.NewCategoriesWindow(a, categories)
					w.Close()
				}
			}
		}
	}, w)
	d.Show()
}
