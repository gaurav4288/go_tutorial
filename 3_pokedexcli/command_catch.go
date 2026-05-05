package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	baseExperience := pokemon.BaseExperience

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	// Create local RNG (modern Go way)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	roll := r.Float64()

	catchChance := 100.0 / float64(baseExperience)
	if catchChance > 1 {
		catchChance = 1
	}

	if roll < catchChance {
		fmt.Printf("🎉 You caught %s!\n", pokemon.Name)
	} else {
		fmt.Printf("💨 %s escaped!\n", pokemon.Name)
	}

	return nil
}