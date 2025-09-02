package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func commandCatch(cfg *config, args ...string) error{
	if len(args) < 1 {
		fmt.Println("Usage: catch <pokemon_name>")
		return nil
	}

	pokemonName := strings.ToLower(args[0])

	// Fetch pokemon from API
	pokemon, err := cfg.pokeapiClient.CatchPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	// Random chance based on base experience
	roll := rand.Intn(pokemon.BaseExperience + 50)

	if roll < 50 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}