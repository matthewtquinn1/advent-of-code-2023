package challenges

import (
	"aoc-2023/pkg/inpututils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func GetAnswerTwoA() {
	games := inpututils.GetFileInput("2-input.txt")

	// Sum of Game IDs.
	totalScore := 0

	// Update score with Game ID when it is a winning game.
	for _, game := range strings.Split(games, "\n") {

		if isWinningGame(game) {
			re := regexp.MustCompile("[0-9]+")
			gameId, err := getGameId(game, re)
			if err != nil {
				fmt.Println("Failed to get game ID. Exiting early.")
				return
			} else {
				totalScore += gameId
			}
		}
	}

	fmt.Println("Challenge 2 - total score is ", totalScore)
}

func isWinningGame(game string) bool {
	// Get the reveals.
	reveals := strings.SplitAfter(game, ":")[1]

	// Check each reveal pile.
	for _, reveal := range strings.Split(reveals, ";") {

		// Get pile of dices.
		piles := strings.Split(reveal, ",")

		// Check each pile to make sure it follows the max count rule for each color.
		for _, pile := range piles {
			pileComponents := strings.Fields(pile)
			diceCount, _ := strconv.Atoi(pileComponents[0])
			diceColor := pileComponents[1]

			// Fail out when any dice count fails to pass.
			if !dicePassesCheck(diceColor, diceCount) {
				return false
			}
		}
	}

	// No failures detected.
	return true
}

func dicePassesCheck(diceColor string, diceCount int) bool {
	// Map of dice colors to their maximum allowed counts.
	maxCounts := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	// Check if the dice count exceeds the maximum for its color.
	if maxCount, exists := maxCounts[diceColor]; exists && diceCount > maxCount {
		return false
	}

	return true
}

func getGameId(game string, re *regexp.Regexp) (int, error) {
	// Section containing ID.
	sectionWithId := strings.SplitAfter(game, ":")[0]

	// Extract ID from section using regex.
	gameIdAsString := re.FindString(sectionWithId)

	return strconv.Atoi(gameIdAsString)
}
