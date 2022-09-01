//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name

type Player struct {
	Name  string
	Hp    int
	MaxHp int
	Ep    int
	MaxEp int
}

func (p *Player) ModifyHP(hp, maxhp int) {
	p.Hp = hp
	p.MaxHp = maxhp
}
func (p *Player) ModifyEP(ep, maxep int) {
	p.Ep = ep
	p.MaxEp = maxep
}

func (p *Player) ChangeName(new string) {
	p.Name = new
}

func (p *Player) Display() {
	fmt.Println("_____________________________________________________________________________")
	fmt.Printf("Player:  %v\t | HP %v of %v | EP %v of %v", p.Name, p.Hp, p.MaxHp, p.Ep, p.MaxEp)
	fmt.Println()
}

//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.

//  - Print out the statistic change within each function
//  - Execute each function at least once

func main() {

	player := Player{Name: "Mike", Hp: 50, MaxHp: 100, Ep: 30, MaxEp: 60}
	player.Display()

	player.ChangeName("Michail")
	player.ModifyHP(100, 100)
	player.ModifyEP(100, 100)

	player.Display()

}
