//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:ch
//* Implement a player having the following statistics:
//  - health, Max health
//  - energy, Max energy
//  - name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - name

type Player struct {
	health, maxHealth uint
	energy, maxEnergy uint
	name              string
}

func printPlayerStats(p *Player) {
	var state string
	fmt.Println()

	if p.health == 0 {
		state = state + fmt.Sprint(" ☠️ ")
	}
	if p.energy == 0 {
		state = state + fmt.Sprint(" 🪫 ")
	}
	fmt.Printf("%v\t |HP:%v\t|EN: %v \t %s", p.name, p.health, p.energy, state)
	fmt.Println()
}

func MakePlayer(name string, hp, maxHP, ep, maxEP uint) *Player {

	pl := Player{
		name:      name,
		health:    hp,
		maxHealth: maxHP,
		energy:    ep,
		maxEnergy: maxEP,
	}
	return &pl
}

func (p *Player) addHealth(points uint) {
	p.health += points
	if p.health >= p.maxHealth {
		p.health = p.maxHealth
	}
	printPlayerStats(p)
}

func (p *Player) applyDamage(points uint) (dead bool) {
	dead = false
	if int(p.health)-int(points) <= 0 {
		p.health = 0
		dead = true
		fmt.Print(" ☠️ exhausted \n")
	} else {
		p.health -= points
	}
	printPlayerStats(p)
	return dead
}

func (p *Player) useEnergy(points uint) (exhausted bool) {
	exhausted = false
	if int(p.energy)-int(points) <= 0 {
		p.energy = 0
		exhausted = true
		fmt.Print(" ⚡️ exhausted \n")
	} else {
		p.energy -= points
	}
	printPlayerStats(p)
	return exhausted
}

func (p *Player) addEnergy(points uint) {
	p.energy += points
	if p.energy >= p.maxEnergy {
		p.energy = p.maxEnergy
	}
	printPlayerStats(p)
}

func main() {

	me := MakePlayer("Michail", 50, 100, 50, 100)
	printPlayerStats(me)
	me.addEnergy(10)
	me.addHealth(30)
	me.applyDamage(100)
	me.useEnergy(100)
	me.addEnergy(10)
	me.addHealth(30)
}
