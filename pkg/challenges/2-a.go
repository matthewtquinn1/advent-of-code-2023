package challenges

import (
	"aoc-2023/pkg/inpututils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	ID      int
	Reveals []Reveal
}

type Reveal struct {
	dicePiles []DicePile
}

type DicePile struct {
	diceColor string
	diceCount int
}

var re = regexp.MustCompile("[0-9]+") // Compile once and use everywhere

func GetAnswerTwoA() {
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
			if game.isWinningGame() {
				totalScore += game.ID
			}
		}
	}

	fmt.Println("Challenge 2 - total score is ", totalScore)
}

func (g *Game) isWinningGame() bool {
	for _, reveal := range g.Reveals {
		for _, pile := range reveal.dicePiles {
			if !dicePassesCheck(pile.diceColor, pile.diceCount) {
				// Exit early when any dice count fails to pass.
				return false
			}
		}
	}

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

func getGameId(sectionWithId string) (int, error) {
	gameIdAsString := re.FindString(sectionWithId)

	return strconv.Atoi(gameIdAsString)
}

func parseGame(gameLine string) (Game, error) {
	parts := strings.Split(gameLine, ":")
	idPart := parts[0]

	reveals := parseReveals(parts[1])

	id, err := getGameId(idPart)
	if err != nil {
		return Game{}, err
	}

	return Game{
		ID:      id,
		Reveals: reveals,
	}, nil
}

func parseReveals(revealsPart string) []Reveal {
	reveals := []Reveal{}
	for _, reveal := range strings.Split(revealsPart, ";") {
		piles := parseDicePiles(strings.Split(reveal, ","))

		reveals = append(reveals, Reveal{
			dicePiles: piles,
		})
	}

	return reveals
}

func parseDicePiles(pilesPart []string) []DicePile {
	piles := []DicePile{}
	for _, pile := range pilesPart {
		pileComponents := strings.Fields(pile)
		diceCount, _ := strconv.Atoi(pileComponents[0])

		piles = append(piles, DicePile{
			diceCount: diceCount,
			diceColor: pileComponents[1],
		})
	}

	return piles
}
