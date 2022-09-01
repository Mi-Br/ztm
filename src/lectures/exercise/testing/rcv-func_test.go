//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func TestHealth(t *testing.T) {

	player := Player{
		name:      "John87",
		health:    50,
		maxHealth: 100,
		energy:    50,
		maxEnergy: 100,
	}

	player.addHealth(100)
	if player.health <= player.maxHealth {
		t.Errorf("Cannot exceed maximum health. Got %v, expect %v", player.health, player.maxHealth)
	}
	player.applyDamage(player.maxHealth + 1)
	if player.health == 0 {
		t.Errorf("Health cannot be negative")
	}

}
