package hangman

import "strings"

type Game struct {
	State        string //Game state
	Letters      []string
	FoundLetters []string
	UsedLetters  []string
	TurnLeft     int
}

func New(turns int, word string) *Game {
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := range letters {
		found[i] = "_"
	}

	//New game
	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnLeft:     turns,
	}

	return g
}
