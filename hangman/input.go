package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
func ReadGuess() (guess string, err error) {
	valid := false //initilize valid in false

	for !valid {
		fmt.Print("what is your letter ? ")
		guess , err :=reader.ReadString('\n')

		if err != nil {
			return guess , err
		}

		guess = strings.TrimSpace(guess)

		if len(guess) != 1 {
			fmt.Printf("Invalid letters input. letter=%v , len=%v" , guess , len(guess))
			continue
		}

		valid = true
	}

	return guess, nil
}
