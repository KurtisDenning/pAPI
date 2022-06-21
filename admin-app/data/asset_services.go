package data

import "fyne.io/fyne/v2"

func GetTrashImage() (fyne.Resource, error) {
	resource, err := fyne.LoadResourceFromPath("./assets/images/trash.svg")
	if err != nil {
		return nil, err
	}
	return resource, nil
}
