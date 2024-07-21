package dndUtil

import (
	"fmt"

	dnd "github.com/zenayr/adventurecompanion/internal/struct"
)

func PrintRaces(races dnd.Races) {
	for i := 0; i < len(races.Races); i++ {
		fmt.Println(races.Races[i].Name)
		fmt.Println(races.Races[i].Size)
		fmt.Println(races.Races[i].Speed)
		fmt.Println(races.Races[i].Ability)
		fmt.Println(races.Races[i].Proficiency)
		fmt.Println(races.Races[i].SpellAbility)
		fmt.Println("")
	}

}
