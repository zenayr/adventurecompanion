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
