package xmlparser

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	dndStruct "github.com/zenayr/adventurecompanion/internal/struct"
)

func GetBytesFromXml(FilePath string) ([]byte, error) {
	splitString := strings.Split(FilePath, ".")
	fileFormat := splitString[len(splitString)-1]

	if fileFormat != "xml" {
		return nil, errors.New("wrong format: file format is not xml")
	}

	xmlFile, err := os.Open(FilePath)
	if err != nil {
		return nil, err
	}

	defer xmlFile.Close()

	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	return byteValue, nil
}

func ParseXml(byteValue []byte) (dndStruct.Races, error) {
	var races dndStruct.Races

	if len(byteValue) <= 0 {
		return races, errors.New("no data provided: byteValue length <= 0")
	}

	err := xml.Unmarshal(byteValue, &races)
	if err != nil {
		return races, err
	}

	return races, nil
}

func MakeXml(r dndStruct.Race) {
	var races [1]dndStruct.Race
	races[0] = r

	v := &dndStruct.Races{Races: races[:]}

	output, err := xml.MarshalIndent(v, " ", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
}
