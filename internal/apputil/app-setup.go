package apputil

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/cmd/fyne_demo/data"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	dndstructs "github.com/zenayr/adventurecompanion/internal/struct"
)

func MakeRacesAccordion(races dndstructs.Races) (fyne.CanvasObject, error) {
	if len(races.Races) <= 0 {
		return nil, errors.New(("error: no races"))
	}
	ac := widget.NewAccordion()

	for i := 0; i < len(races.Races); i++ {
		size := widget.NewLabel(races.Races[i].Size)
		speed := widget.NewLabel(races.Races[i].Speed)
		ability := widget.NewLabel(races.Races[i].Ability)
		procifiency := widget.NewLabel(races.Races[i].Proficiency)
		spellAbility := widget.NewLabel(races.Races[i].SpellAbility)

		detail := container.New(layout.NewGridLayout(5), size, speed, ability, procifiency, spellAbility)

		newItem := widget.NewAccordionItem(races.Races[i].Name, detail)
		ac.Append(newItem)

	}

	if len(ac.Items) <= 0 {
		return nil, errors.New(("error: eo items"))
	}
	return ac, nil

}

func MakeRacesList(races dndstructs.Races) (fyne.CanvasObject, error) {
	if len(races.Races) <= 0 {
		return nil, errors.New(("error: no races"))
	}

	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Select An Item From The List")
	card := widget.NewCard("", "", nil)
	hbox := container.NewHBox(card)

	list := widget.NewList(
		func() int {
			return len(races.Races)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(races.Races[id].Name)
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		label.SetText(races.Races[id].Name)
		icon.SetResource(theme.DocumentIcon())
		MakeRaceCard(races.Races[id])
	}
	list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select An Item From The List")
		icon.SetResource(nil)
		card = widget.NewCard("", "", nil)
	}
	list.Select(125)
	list.SetItemHeight(5, 50)
	list.SetItemHeight(6, 50)

	return container.NewHSplit(list, container.NewCenter(hbox)), nil
}

func MakeRaceCard(race dndstructs.Race) fyne.Widget {
	card := widget.NewCard(race.Name, "Race", nil)
	card.Image = canvas.NewImageFromResource(data.FyneLogo)

	return card
}
