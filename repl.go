package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Dex struct {
	config *Config
}

func newDex() *Dex {
	conf := NewConfig()
	dex := Dex{
		config: conf,
	}
	return &dex
}

func (d *Dex) run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nPokkdex >")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := words[1:]

		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(d.config, args...)
			if err != nil {
				log.Panic(err)
			}

		} else {
			fmt.Println()
			fmt.Printf("%s command does not exist!\nTo get the list of avaliable commands use: help", commandName)
			fmt.Println()
			continue
		}

	}
}
