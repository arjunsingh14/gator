package main

import (
	"fmt"
	"log"
	"os"
	"github.com/arjunsingh14/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	s := &state{ cfg: &cfg }

	commands := commands{ handlerMap: make(map[string]func(*state, command) error) }
	commands.register("login", handlerLogin)
	cliArgs := os.Args
	if len(cliArgs) < 2 {
		fmt.Printf("Invalid arguments")
		os.Exit(1)
	}

	name := cliArgs[1]
	commandArgs := cliArgs[2:]
	err = commands.run(s, command{ name: name, args: commandArgs})
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

}