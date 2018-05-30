package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

// Player represents a player with an ID
// and a terminal colour.
type Player struct {
	ID    int
	Color func(format string, a ...interface{})
}

var (
	board   = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	player1 = Player{1, color.Cyan}
	player2 = Player{2, color.Yellow}
)

func main() {
	var finished bool
	var player Player
	turnNumber := 1

	for !finished {
		displayBoard(board)
		if turnNumber%2 == 1 {
			player1.Color("ready player one")
			player = player1
		} else {
			player2.Color("ready player two")
			player = player2
		}

		currentMove, err := getMove(player)

		if err != nil {
			color.Red(err.Error())
			time.Sleep(1 * time.Second)
		} else {
			if board, err = executeMove(currentMove, player, board); err != nil {
				color.Red(err.Error())
				time.Sleep(1 * time.Second)
			}
		}

		result := checkWinner(board)
		if result > 0 {
			fmt.Printf("Player %d wins!\n\n", result)
			finished = true
		} else {
			turnNumber++
		}
	}

}

func getMove(player Player) (int, error) {
	var move int
	player.Color("make your move")
	_, err := fmt.Scan(&move)
	if err != nil {
		return move, fmt.Errorf("please select a numbered position, miss a turn")
	}
	return move, nil
}

func executeMove(move int, player Player, b [9]int) ([9]int, error) {
	if player.ID == 1 {
		if b[move] == 10 {
			return b, fmt.Errorf("invalid move, miss a turn")
		}
		b[move] = 1
		return b, nil
	} else if player.ID == 2 {
		if b[move] == 1 {
			return b, fmt.Errorf("invalid move, miss a turn")
		}
		b[move] = 10
		return b, nil
	}
	return b, nil
}

func displayBoard(b [9]int) {
	for i, v := range b {
		if v == 0 {
			// empty space. Display number
			fmt.Printf("%d", i)
		} else if v == 1 {
			fmt.Printf("X")
		} else if v == 10 {
			fmt.Printf("O")
		}
		// And now the decorations
		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(" | ")
		}
	}
}

func checkWinner(b [9]int) int {
	sums := [9]int{0, 0, 0, 0, 0, 0, 0, 0}
	for _, v := range b[0:2] {
		sums[7] += v
	}
	for _, v := range b[3:5] {
		sums[6] += v
	}
	for _, v := range b[6:8] {
		sums[5] += v
	}

	sums[0] = b[2] + b[4] + b[6]
	sums[1] = b[0] + b[3] + b[6]
	sums[2] = b[1] + b[4] + b[7]
	sums[3] = b[2] + b[5] + b[8]
	sums[4] = b[0] + b[4] + b[8]

	for _, v := range sums {
		if v == 3 {
			return 1
		} else if v == 30 {
			return 2
		}
	}
	return 0
}
