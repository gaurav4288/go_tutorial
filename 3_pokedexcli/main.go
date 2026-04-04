package main

import (
	"time"

	"github.com/gaurav4288/3_pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
