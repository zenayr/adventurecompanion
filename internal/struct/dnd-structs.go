package dnd

import "encoding/xml"

type Races struct {
	XMLName xml.Name `xml:"races"`
	Races   []Race   `xml:"race"`
}

type Race struct {
	XMLName      xml.Name `xml:"race"`
	Name         string   `xml:"name"`
	Size         string   `xml:"size"`
	Speed        string   `xml:"speed"`
	Ability      string   `xml:"ability"`
	Proficiency  string   `xml:"proficiency"`
	SpellAbility string   `xml:"spellability"`
}
