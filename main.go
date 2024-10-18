package main

import (
	"fmt"
	"start/enemies"
)

func main() {
	fmt.Println("hi")
	//rooms.RandomMatrix("normal")

	var ork1 = enemies.NewOrk()
	fmt.Println(ork1.InitBasisOrk())
	ork1.SetStats(10, 10, 10, 10)
	ork1.SeeStats()
}
