package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	dnd "github.com/zenayr/adventurecompanion/internal/struct"
)

func Parse() dnd.Races {
	xmlFile, err := os.Open("../../xml/races-phb.xml")
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	var races dnd.Races
	xml.Unmarshal(byteValue, &races)

	return races
}

func MakeXml(r dnd.Race) {
	//r := &dnd.Race{Name: "test", Size: "10", Speed: "10", Ability: " ", Proficiency: " ", SpellAbility: " "}
	var races [1]dnd.Race
	races[0] = r

	v := &dnd.Races{Races: races[:]}

	output, err := xml.MarshalIndent(v, " ", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
}
