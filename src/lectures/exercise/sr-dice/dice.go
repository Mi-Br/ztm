//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Dice struct {
	sides int
}

func (d *Dice) roll() (out int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(d.sides-1) + 1
}

func rollTheDice(times, diceCount, sides int) {

	var totalScore int = 0
	var d Dice = Dice{sides}
	for times > 0 {
		for diceCount > 0 {
			totalScore += d.roll()
			diceCount--
		}
		times--
	}

	switch {
	case totalScore == 2 && diceCount == 2:
		fmt.Print("Snake eyes")
	case totalScore == 7:
		fmt.Print("Lucky 7")
	case totalScore%2 == 0:
		fmt.Print("Even")
	default:
		fmt.Print("Odd")
	}
	fmt.Println("\t", totalScore)
	//  - "Snake eyes": when the total roll is 2, and total dice is 2
	//  - "Lucky 7": when the total roll is 7
	//  - "Even": when the total roll is even
	//  - "Odd": when the total roll is odd

}

func main() {

	rollTheDice(1, 3, 6)

}
