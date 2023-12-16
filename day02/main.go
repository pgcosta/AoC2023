package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/errors/fmt"
)

var (
	blueCubes  = 14
	redCubes   = 12
	greenCubes = 13
)

type Draw struct {
	green int
	blue  int
	red   int
}

type Game struct {
	id    int
	draws []Draw
}

func loadInput(filePath string) []string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("File reading error", err)
		return []string{}
	}

	return strings.Split(string(data), "\n")
}

func loadGames(inputLines []string) []Game {
	var games []Game
	var game Game
	for _, line := range inputLines {
		expression := regexp.MustCompile(`^\w+ ([0-9]+): (.+)`)
		regexResults := expression.FindAllStringSubmatch(line, -1)

		if len(regexResults[0]) < 1 {
			panic("Could not parse game id")
		}

		gameId, err := strconv.Atoi(regexResults[0][1])
		if err != nil {
			panic(err)
		}

		game.id = gameId
		game.draws = parseGame(regexResults[0][2])
		games = append(games, game)
	}
	return games
}

func makeDraw(draw *Draw, color string, number int) Draw {
	switch color {
	case "blue":
		draw.blue = number
	case "red":
		draw.red = number
	case "green":
		draw.green = number
	}
	return *draw
}

func parseGame(game string) []Draw {
	games := strings.Split(game, ";")
	var draw Draw
	var draws []Draw
	var color string

	for _, cubeThrows := range games {
		cubeThrows = strings.TrimSpace(cubeThrows)
		expression := regexp.MustCompile(`([0-9]+) (\w+)`)
		throw := expression.FindAllStringSubmatch(cubeThrows, -1)

		for _, t := range throw {
			color = t[2]
			number, err := strconv.Atoi(t[1])
			if err != nil {
				panic(err)
			}
			draw = makeDraw(&draw, color, number)
		}
		draws = append(draws, draw)
		draw = Draw{}
	}

	return draws
}

func isGameValid(game Game) bool {
	for _, draw := range game.draws {
		if draw.blue > blueCubes || draw.red > redCubes || draw.green > greenCubes {
			return false
		}
	}
	return true
}

func sumValidGames(games []Game) int {
	var sum int
	for _, game := range games {

		if isGameValid(game) {
			sum += game.id
		}
	}
	return sum
}

func powerMinimumSet(games []Game) int {
	var total int
	for _, game := range games {
		minimum := minimumSetOfCubes(game.draws)
		total += minimum.blue * minimum.red * minimum.green
	}
	return total
}

func minimumSetOfCubes(draws []Draw) Draw {
	var maxRed, maxBlue, maxGreen int
	for _, draw := range draws {
		if draw.red > maxRed {
			maxRed = draw.red
		}
		if draw.blue > maxBlue {
			maxBlue = draw.blue
		}
		if draw.green > maxGreen {
			maxGreen = draw.green
		}
	}

	return Draw{green: maxGreen, blue: maxBlue, red: maxRed}
}

func main() {
	var games []Game
	inputLines := loadInput("input.txt")
	games = loadGames(inputLines)
	part1 := sumValidGames(games)
	part2 := powerMinimumSet(games)
	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}
