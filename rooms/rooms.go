/*	Name: BenjaminBlumch3n
	Date: 11.10.2024
	Reason: Initializing Rooms
*/

package rooms

import (
	"fmt"
	"math/rand"
)

func RandomMatrix(difficulty string) {

	var randomMatrix = [][]int{}
	var matrixLength int
	var cellValue int

	//adjusting matrix for difficulty
	//requirement: set difficulty
	//result: matrix with new length (square)

	switch difficulty {
	case "easy":
		matrixLength = 3
	case "normal":
		matrixLength = 5
	case "hard":
		matrixLength = 10
	case "impossible":
		matrixLength = 100
	}

	//var thisMatrix = [matrixLength][matrixLength]int
	for i := 0; i < matrixLength; i++ {
		for j := 0; j < matrixLength; j++ {

			cellProbability := rand.Intn(100)

			switch {
			case cellProbability <= 40:
				cellValue = 0
			case cellProbability <= 50 && cellProbability > 40:
				cellValue = 1
			case cellProbability <= 70 && cellProbability > 50:
				cellValue = 2
			case cellProbability <= 90 && cellProbability > 70:
				cellValue = 3
			case cellProbability <= 99 && cellProbability > 90:
				cellValue = 4
			case cellProbability == 100:
				cellValue = 5
			}
			//thisMatrix[i][j] = cellValue
			var s []int
			s = append(s, cellValue)
			randomMatrix[i] = s
		}
	}

	fmt.Println(randomMatrix, cellValue)
}
