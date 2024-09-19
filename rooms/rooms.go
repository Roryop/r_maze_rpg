package rooms

import (
	"fmt"
	"math/rand"
)

func RandomMatrix(difficulty string) {

	var randomMatrix = [][]int{{}, {}}
	var s []int
	s = append(s, 1, 2)
	randomMatrix[1] = s
	var matrixLength int
	var cellValue int

	if difficulty == "easy" {
		matrixLength = 3
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
		}
	}

	fmt.Println(randomMatrix, cellValue)
}
