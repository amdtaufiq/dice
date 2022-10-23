package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var game Game
	game.Game(3, 4)
	game.start()

}

type Dice struct {
	Value int
}

func (d *Dice) getValue() (value int) {
	value = d.Value
	return
}

// func (d *Dice) roll() (value int) {
// 	rand.Seed(time.Now().UTC().UnixNano())
// 	d.Value = rand.Intn(6) + 1
// 	value = d.Value
// 	fmt.Printf("value: %v\n", value)
// 	return
// }

type Player struct {
	Dices    []Dice
	Name     string
	Position int
	Point    int
}

func (p *Player) Player(numberOfDice, position int, name string) Player {
	p.Point = 0
	p.Position = position
	p.Name = name
	for i := 0; i < numberOfDice; i++ {
		p.Dices = append(p.Dices, Dice{})
	}

	return *p
}

func (p *Player) getDices() (dices []Dice) {
	dices = p.Dices
	return
}

func (p *Player) getName() (name string) {
	name = p.Name
	return
}

func (p *Player) getPosition() (position int) {
	position = p.Position
	return
}

func (p *Player) addPoint() {
	p.Point++
	fmt.Print("Player => ", p.Name)
	fmt.Println(" & Point => ", p.Point)
}

func (p *Player) getPoint() (point int) {
	// fmt.Println(p.Name)
	point = p.Point
	fmt.Print("Player => ", p.Name)
	fmt.Println(" & Point => ", p.Point)
	return
}

func (p *Player) play() {
	for i, _ := range p.Dices {
		rand.Seed(time.Now().UnixNano())
		value := rand.Intn(6) + 1
		fmt.Print(value, " ")
		p.Dices[i].Value = value
	}
}

func (p *Player) removeDice(index, player int) {
	// p.Dices = append(p.Dices[:index], p.Dices[index+1:]...)
	p.Dices[index].Value = 0
	// fmt.Printf("player => %d index => %d", player+1, index)
	// fmt.Println()
}

func (p *Player) addDice(dice Dice) {
	p.Dices = append(p.Dices, dice)
}

const (
	REMOVED_WHEN_DICE_TOP = 6
	MOVE_WHEN_DICE_TOP    = 1
)

type Game struct {
	Players               []Player
	Round                 int
	NumberOfPlayer        int
	NumberOfDicePerPlayer int
}

func (g *Game) Game(numberOfPlayer, numberOfDicePerPlayer int) Game {
	g.Round = 0
	g.NumberOfPlayer = numberOfPlayer
	g.NumberOfDicePerPlayer = numberOfDicePerPlayer

	for i := 0; i < numberOfPlayer; i++ {
		var player Player
		g.Players = append(g.Players, player.Player(numberOfDicePerPlayer, i, fmt.Sprint(i+1)))
	}
	return *g
}

func (g *Game) displayRound() Game {
	fmt.Printf("Giliran %d ", g.Round)
	return *g
}

func (g *Game) displayTopSideDice(title string) Game {

	// fmt.Printf("%s:\n", title)
	for _, player := range g.Players {
		// 	fmt.Printf("\tPemain #%s(%d):", player.getName(), player.getPoint())
		// 	diceToSide := ""
		// 	for _, dice := range player.getDices() {
		// 		diceToSide += fmt.Sprintf("%d, ", dice.getValue())
		// 	}

		// 	fmt.Printf("%s\n", strings.TrimRight(diceToSide, ", "))
		fmt.Println("Get Point")
		player.getPoint()
	}
	return *g
}

func (g *Game) displayWinner(player Player) Game {
	fmt.Println("Pemenang")
	fmt.Printf("Pemain %s", player.getName())
	return *g
}

func (g *Game) getWinner() (winner Player) {
	highScore := 0
	for _, player := range g.Players {
		if player.getPoint() > highScore {
			highScore = player.getPoint()
			winner = player
		}
	}
	return
}

func (g *Game) start() {
	fmt.Printf("Pemain = %d, Dadu = %d \n", g.NumberOfPlayer, g.NumberOfDicePerPlayer)
	index := 0
	for index < 2 {
		index++
		g.Round++
		// var diceCarryForward [][]Dice

		for _, player := range g.Players {
			player.play()
		}

		g.displayRound()
		g.displayTopSideDice("Lempar Dadu")

		for i, player := range g.Players {
			// fmt.Println(player)
			// var tempDiceArray []Dice
			for j, dice := range player.getDices() {
				// fmt.Print(dice)
				if dice.getValue() == REMOVED_WHEN_DICE_TOP {
					// fmt.Println(" => Masuk")
					fmt.Println("Add Point")
					player.removeDice(j, i)
				}

				// if dice.getValue() == MOVE_WHEN_DICE_TOP {
				// 	if player.getPosition() == g.NumberOfPlayer-1 {
				// 		g.Players[0].addDice(dice)
				// 		player.removeDice(j)
				// 	} else {
				// 		tempDiceArray = append(tempDiceArray, dice)
				// 		player.removeDice(j)
				// 	}
				// }
			}

			// diceCarryForward = append(diceCarryForward, tempDiceArray)

			// if diceCarryForward[i] != nil && len(diceCarryForward[i]) > 0 {
			// 	for _, dice := range diceCarryForward[i] {
			// 		player.addDice(dice)
			// 	}

			// 	diceCarryForward = [][]Dice{}
			// }
		}

		g.displayTopSideDice("Setelah Evaluasi")

		playerHasDice := g.NumberOfPlayer

		for _, player := range g.Players {
			if len(player.getDices()) <= 0 {
				playerHasDice--
			}
		}

		if playerHasDice == 1 {
			g.displayWinner(g.getWinner())
			break
		}
	}
}
