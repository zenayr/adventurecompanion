package main

import (
	//"fmt"
	//"os"
	//"path/filepath"
	"fmt"
	"log"

	dnd "github.com/zenayr/adventurecompanion/internal/struct"
	parser "github.com/zenayr/adventurecompanion/internal/util"
)

func main() {
	// ex, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }

	// exePath := filepath.Dir(ex)
	// os.Mkdir(exePath+"/"+"xml", 0644)
	// fmt.Println("Directory xml created in " + exePath)

	// races := parser.Parse()

	// for i := 0; i < len(races.Races); i++ {
	// 	fmt.Println(races.Races[i].Name)
	// 	fmt.Println(races.Races[i].Size)
	// 	fmt.Println(races.Races[i].Speed)
	// 	fmt.Println(races.Races[i].Ability)
	// 	fmt.Println(races.Races[i].Proficiency)
	// 	fmt.Println(races.Races[i].SpellAbility)
	// 	fmt.Println("")
	// }
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
