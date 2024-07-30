package stdio

import (
	"fmt"
	"log"

	dndUtil "github.com/zenayr/adventurecompanion/internal"
	parser "github.com/zenayr/adventurecompanion/internal/util"
)

func ParseXmlFromCli() {
	path := GetPath()

	byteValue, err := parser.GetBytesFromXml(path)
	if err != nil {
		log.Fatal(err)
	}

	races, err := parser.ParseXml(byteValue)
	if err != nil {
		log.Fatal(err)
	}

	if len(races.Races) <= 0 {
		log.Fatal("Error: No Races")

	}

	dndUtil.PrintRaces(races)

	fmt.Print("any key to exit...")
	fmt.Scanf("%s")
}

func GetPath() string {
	fmt.Print("FilePath: ")

	var path string
	n, err := fmt.Scanf("%s", &path)
	if err != nil {
		fmt.Printf("Ivalid value at index %d\n", n)
		log.Fatal(err)
	}

	return path
}
