package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/zenayr/adventurecompanion/internal/apputil"
	xmlparser "github.com/zenayr/adventurecompanion/internal/util"
)

func main() {

	byteValue, err := xmlparser.GetBytesFromXml("xml/races-phb.xml")
	if err != nil {
		log.Fatal(err)
	}

	races, err := xmlparser.ParseXml(byteValue)
	if err != nil {
		log.Fatal(err)
	}

	a := app.New()
	w := a.NewWindow("Adventure Companion")

	content, err := apputil.MakeRacesList(races)
	if err != nil {
		fmt.Println(err)
	}
	// content, err := apputil.MakeRacesAccordion(races)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	if content != nil {
		w.SetContent(content)
	}
	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()

}
