package main

import (
	"fmt"
	"strings"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		fmt.Println("Usage: explore <area_name>")
		return nil
	}

	areaName := strings.ToLower(args[0])

	areaResp, err := cfg.pokeapiClient.GetLocationAreaPokemon(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokémon in %s:\n", areaName)
	for _, encounter := range areaResp.PokemonEncounters {
		fmt.Println("-", encounter.Pokemon.Name)
	}

	return nil
}