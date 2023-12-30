package challenges

import (
	"aoc-2023/pkg/inpututils"
	"errors"
	"fmt"
	"strings"
)

func GetAnswerTwoB() {
	games := inpututils.GetFileInput("2-input.txt")

	// Sum of Game IDs.
	totalScore := 0

	// Update score with Game ID when it is a winning game.
	for _, gameLine := range strings.Split(games, "\n") {
		game, err := parseGame(gameLine)
		if err != nil {
			fmt.Println("Failed to get game ID. Exiting early.", err)
			return
		} else {
			score, scoreErr := game.calculateMaxDiceScore()
			if scoreErr != nil {
				fmt.Println("Failed to get score for game.", scoreErr)
			} else {
				totalScore += score
			}
		}
	}

	fmt.Println("Challenge 2 B - total score is ", totalScore)
}

func (g *Game) calculateMaxDiceScore() (int, error) {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	colors := map[string]int{
		"red":   1,
		"green": 2,
		"blue":  3,
	}

	for _, reveal := range g.Reveals {
		for _, pile := range reveal.dicePiles {

			if _, exists := colors[pile.diceColor]; !exists {
				return -1, errors.New("Unsupported dice color; can't calculate score.")
			}

			if pile.diceColor == "red" && pile.diceCount > maxRed {
				maxRed = pile.diceCount
			} else if pile.diceColor == "green" && pile.diceCount > maxGreen {
				maxGreen = pile.diceCount
			} else if pile.diceColor == "blue" && pile.diceCount > maxBlue {
				maxBlue = pile.diceCount
			} else {
				continue
			}
		}
	}

	return maxRed * maxGreen * maxBlue, nil
}
