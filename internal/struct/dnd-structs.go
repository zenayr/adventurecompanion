package dnd

type Races struct {
	Races []Race `xml:"race"`
}

type Race struct {
	Name         string `xml:"name"`
	Size         string `xml:"size"`
	Speed        string `xml:"speed"`
	Ability      string `xml:"ability"`
	Proficiency  string `xml:"proficiency"`
	SpellAbility string `xml:"spellability"`
}
