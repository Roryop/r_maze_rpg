package random

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomNumber(max int) int {
	var finalRandom int
	var randomCalc int = time.Now().Nanosecond()

	for randomCalc < max {
		randomCalc = randomCalc * 9 / 4
	}

	for randomCalc > max {
		randomCalc = randomCalc * 4 / 9
	}

	finalRandom = randomCalc
	fmt.Println(finalRandom)
	fmt.Println(rand.Intn(100))
	return finalRandom
}
