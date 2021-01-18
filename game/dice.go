package game

import (
	"math/rand"
)

func rollDice() int {
	dice1 := rand.Intn(4)
	dice2 := rand.Intn(4)
	o := dice1 + dice2
	if o == 0 {
		o = 12
	}
	return o
}