/*	Name: BenjaminBlumch3n
	Date: 11.10.2024
	Reason: Creating Enemies
*/

package enemies

import "fmt"

type Basis struct {
	name string
	race string

	strength float32
	health   float32
	defence  float32
	speed    float32
}

const bStrength = 10
const bHealth = 100
const bDefence = 5
const bSpeed = 1

type Ork struct {
	Basis
}

func (b *Ork) SetStats(strength float32, health float32, defence float32, speed float32) {
	b.strength = strength
	b.health = health
	b.defence = defence
	b.speed = speed
}

func (b *Ork) SeeStats() [4]float32 {
	var stats [4]float32 = [4]float32{b.strength, b.health, b.defence, b.speed}
	fmt.Println(stats)
	return stats
}

func (b *Ork) InitBasisOrk() ([2]string, [4]float32) {
	b.name = "Günther"
	b.race = "Ork"
	b.strength = bStrength
	b.health = bHealth
	b.defence = bDefence
	b.speed = bSpeed

	var nameRace [2]string = [2]string{b.name, b.race}
	var stats [4]float32 = [4]float32{b.strength, b.health, b.defence, b.speed}
	return nameRace, stats
}

func NewOrk() *Ork {
	var ork *Ork = new(Ork)
	return ork
}

//////////UNTERRICHT////////////

//Ork vs Ork
func OrcFight(ork1, ork2 *Ork) {
	var decision string
	ork1.InitBasisOrk()
	ork2.InitBasisOrk()

	fmt.Println("Zwei Orcs fighten!")

	for ork1.health > 0 && ork2.health > 0 {
		ork1.health = ork1.health - ork2.strength
		if ork1.health > 0 {
			ork2.health = ork2.health - ork1.strength
		}
		fmt.Print("\nOrk1-HP: ", ork1.health, "\nOrk2-HP: ", ork2.health)
		fmt.Print("\n\nKampf abbrechen? Nein = x")
		fmt.Scanln(&decision)
		if decision == "x" {
			break
		}
	}
	if ork1.health <= 0 {
		fmt.Println("\nOrk 2 hat gewonnen!")
	} else {
		fmt.Println("\nOrk 1 hat gewonnen!")
	}
}
