package main

import (
	"regexp"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestRegex(t *testing.T) {
	expression := `^\w+ ([0-9]+):`
	result1 := regexp.MustCompile(expression).FindAllStringSubmatch("Game 1: 8 blue, 3 green; 1 green, 3 red;", -1)
	result100 := regexp.MustCompile(expression).FindAllStringSubmatch("Game 100: 8 blue, 3 green; 1 green, 3 red;", -1)

	assert.Equal(t, result1[0][1], "1")
	assert.Equal(t, result100[0][1], "100")
}

func TestGamesRegex(t *testing.T) {
	input := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	expression := `^.+: (.+)`
	regexResult := regexp.MustCompile(expression).FindAllStringSubmatch(input, -1)
	gamesStrings := strings.Split(regexResult[0][1], ";")

	assert.Equal(t, regexResult[0][1], "8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")
	assert.Equal(t, gamesStrings[0], "8 green, 6 blue, 20 red")
	assert.Equal(t, gamesStrings[1], " 5 blue, 4 red, 13 green")
	assert.Equal(t, gamesStrings[2], " 5 green, 1 red")
}

func TestParseGame(t *testing.T) {
	gameString := "8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	game := parseGame(gameString)

	expected := []Draw{
		Draw{green: 8, blue: 6, red: 20},
		Draw{blue: 5, red: 4, green: 13},
		Draw{green: 5, red: 1, blue: 0},
	}

	assert.Equal(t, game, expected)
}

func TestLoadGames(t *testing.T) {
	lines := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	games := loadGames(lines)
	sum := sumValidGames(games)

	expected := 8
	assert.Equal(t, sum, expected)
}

func TestPowerMinimumSet(t *testing.T) {
	lines := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	games := loadGames(lines)
	sum := powerMinimumSet(games)

	expected := 2286
	assert.Equal(t, sum, expected)
}
