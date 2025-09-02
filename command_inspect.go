package main

import (
	"fmt"
	"strings"
)

func commandInspect(cfg *config, args ...string) error {
    if len(args) < 1 {
        fmt.Println("Usage: inspect <pokemon_name>")
        return nil
    }

    name := strings.ToLower(args[0])
    pokemon, exists := cfg.pokedex[name]
    if !exists {
        fmt.Printf("you have not caught that pokemon\n")
        return nil
    }

    fmt.Printf("Name: %s\n", pokemon.Name)
    fmt.Printf("Height: %d\n", pokemon.Height)
    fmt.Printf("Weight: %d\n", pokemon.Weight)

    fmt.Println("Stats:")
    for _, stat := range pokemon.Stats {
        fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
    }

    fmt.Println("Types:")
    for _, t := range pokemon.Types {
        fmt.Printf("  - %s\n", t.Type.Name)
    }

    return nil
}
