package main

import (
	"fmt"
	"os"

	"github.com/allyRx/hangman/hangman"
)

func main(){
	g := hangman.New(8, "Golang")
	fmt.Println(g)
	
	r , err := hangman.ReadGuess()
	
	if err != nil {
		fmt.Println(err) 
		os.Exit(1)
	}
	
	fmt.Println(r)
}
	
	
	