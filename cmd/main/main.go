package main

import (
	//"fmt"
	//"os"
	//"path/filepath"
	//"fmt"
	//"log"

	"errors"
	"fmt"
	"log"

	internal "github.com/zenayr/adventurecompanion/internal"
	dnd "github.com/zenayr/adventurecompanion/internal/struct"
	parser "github.com/zenayr/adventurecompanion/internal/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	//"fyne.io/fyne/v2/canvas"
	//"fyne.io/fyne/v2/driver/desktop"
	//"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	races := parser.Parse()

	internal.PrintRaces(races)

	if len(races.Races) <= 0 {
		log.Fatal("Error: No Races")

	}

	a := app.New()
	w := a.NewWindow("Adventure Companion")

	content, err := makeAccordion(races)
	if err != nil {
		fmt.Println(err)
	}

	//l1 := widget.NewLabel("Hello World!")
	//l2 := widget.NewLabel("Hello Again!")
	//content := container.New(layout.NewGridLayout(1), l1, l2)

	if content != nil {
		w.SetContent(content)
	}
	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()

}

func inputRaces() {

	r := dnd.Race{}

	fmt.Print("Name: ")
	n, err := fmt.Scanf("%s", &r.Name)
	if err != nil || n < 0 {
		log.Fatal(err)
	}
	fmt.Println()

	fmt.Print("Size: ")
	n1, err := fmt.Scanf("%s", &r.Size)
	if err != nil || n1 < 0 {
		log.Fatal(err)
	}
	fmt.Println()

	fmt.Print("Speed: ")
	n2, err := fmt.Scanf("%s", &r.Speed)
	if err != nil || n2 < 0 {
		log.Fatal(err)
	}
	fmt.Println()

	fmt.Print("Ability: ")
	n3, err := fmt.Scanf("%s", &r.Ability)
	if err != nil || n3 < 0 {
		log.Fatal(err)
	}
	fmt.Println()

	fmt.Print("Proficiency: ")
	n4, err := fmt.Scanf("%s", &r.Proficiency)
	if err != nil || n4 < 0 {
		log.Fatal(err)
	}
	fmt.Println()

	fmt.Print("SpellAbility: ")
	n5, err := fmt.Scanf("%s", &r.SpellAbility)
	if err != nil || n5 < 0 {
		log.Fatal(err)
	}
	fmt.Println()

	parser.MakeXml(r)

}

func makeAccordion(races dnd.Races) (fyne.CanvasObject, error) {
	if len(races.Races) <= 0 {
		return nil, errors.New(("Error: No Races"))
	}
	ac := widget.NewAccordion()

	for i := 0; i < len(races.Races); i++ {
		size := widget.NewLabel(races.Races[i].Size)
		speed := widget.NewLabel(races.Races[i].Speed)
		ability := widget.NewLabel(races.Races[i].Ability)
		procifiency := widget.NewLabel(races.Races[i].Proficiency)
		spellAbility := widget.NewLabel(races.Races[i].SpellAbility)

		detail := container.New(layout.NewGridLayout(5), size, speed, ability, procifiency, spellAbility)

		newItem := widget.NewAccordionItem(races.Races[i].Name,
			detail,
		)
		ac.Append(newItem)

	}

	if len(ac.Items) <= 0 {
		return nil, errors.New(("Error: No Items"))
	}
	return ac, nil

}
