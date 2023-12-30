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
	maxCounts := make(map[string]int)
	supportedColors := map[string]bool{
		"red":   true,
		"green": true,
		"blue":  true,
	}

	for _, reveal := range g.Reveals {
		for _, pile := range reveal.dicePiles {
			if _, exists := supportedColors[pile.diceColor]; !exists {
				return -1, errors.New("Unsupported dice color; can't calculate score.")
			}

			if pile.diceCount > maxCounts[pile.diceColor] {
				maxCounts[pile.diceColor] = pile.diceCount
			}
		}
	}

	if len(maxCounts) != len(supportedColors) {
		return -1, errors.New("Not all colors are present.")
	}

	score := 1
	for _, count := range maxCounts {
		score *= count
	}

	return score, nil
}
