package apputil

import (
	"errors"
	"fmt"

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

	//icon := widget.NewIcon(nil)
	//label := widget.NewLabel("Select An Item From The List")
	raceCard := widget.NewCard("", "", widget.NewLabel("Select An Item From The List"))
	hbox := container.NewHBox(raceCard)

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
		raceCard.SetContent(MakeRaceCard(races.Races[id]))
	}

	list.OnUnselected = func(id widget.ListItemID) {
		raceCard.SetContent(widget.NewLabel("Select An Item From The List"))
	}
	list.Select(125)
	list.SetItemHeight(5, 50)
	list.SetItemHeight(6, 50)

	return container.NewHSplit(list, container.NewCenter(hbox)), nil
}

func MakeRaceCard(race dndstructs.Race) fyne.CanvasObject {
	size := widget.NewLabel("Size: " + race.Size)
	speed := widget.NewLabel("Speed: " + race.Speed)
	ability := widget.NewLabel("Ability: " + race.Ability)
	proficiency := widget.NewLabel("Proficiency: " + race.Proficiency)
	spellAbility := widget.NewLabel("SpellAbility: " + race.SpellAbility)

	content := container.NewGridWithRows(5, size, speed, ability, proficiency, spellAbility)
	raceCard := widget.NewCard(race.Name, "", content)
	raceCard.Image = canvas.NewImageFromResource(data.FyneLogo)

	return raceCard
}

func MakeRaceFormTab(_ fyne.Window) fyne.CanvasObject {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Dragonborn")
	abilityScoreEntry := widget.NewEntry()
	ageEntry := widget.NewMultiLineEntry()
	speedEntry := widget.NewEntry()

	sizeSelect := widget.NewSelectEntry([]string{
		"Tiny",
		"Small",
		"Medium",
		"Large",
		"Huge",
		"Gargantuan",
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: nameEntry, HintText: "Name of the race"},
			{Text: "Ability Score", Widget: abilityScoreEntry, HintText: "Ability score increases"},
			{Text: "Age", Widget: ageEntry, HintText: "How the race ages"},
			{Text: "Size", Widget: sizeSelect, HintText: "Size of the race"},
			{Text: "Speed", Widget: speedEntry, HintText: "Speed of the race"},
		},
		OnCancel: func() {
			fmt.Println("Cancelled")
		},
		OnSubmit: func() {
			fmt.Println("Form submitted")
		},
	}
	return form
}

func MakeTraitFormTab(_ fyne.Window) fyne.CanvasObject {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Dark Vision")
	descriptionEntry := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: nameEntry, HintText: "Trait name"},
			{Text: "Description", Widget: descriptionEntry, HintText: "Trait description"},
		},
		OnCancel: func() {
			fmt.Println("Cancelled")
		},
		OnSubmit: func() {
			fmt.Println("Form submitted")
		},
	}
	return form
}
