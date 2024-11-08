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

	//var randomMatrix = [][]int{}
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

	//creating a slice made up of slices (2-dimensional)
	var randomMatrix = make([][]int, matrixLength)
	for i := 0; i < matrixLength; i++ {
		randomMatrix[i] = make([]int, matrixLength)
	}

	//var thisMatrix = [matrixLength][matrixLength]int
	for i := 0; i < matrixLength; i++ {

		var s []int //creating slice to later put into matrix

		//adding value to cell reliant on probability
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
			case cellProbability <= 100 && cellProbability > 90:
				if i == 0 || j == 0 || i == matrixLength-1 || j == matrixLength-1 { //room on edge can´t have 4 doors
					cellValue = 3
				} else {
					cellValue = 4
				}
			}
			//thisMatrix[i][j] = cellValue
			s = append(s, cellValue) //making slice bigger
		}
		randomMatrix[i] = s //inputting slice
	}

	fmt.Println(randomMatrix)
}
