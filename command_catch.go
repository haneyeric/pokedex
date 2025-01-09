package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, params []string) error {
	_, ok := cfg.pokedex[params[1]]
	if ok {
		fmt.Printf("You already caugth %s\n", params[1])
		return nil
	}

	pokeResp, err := cfg.pokeapiClient.GetPokenmon(params[1])
	if err != nil {
		fmt.Println("Couldn't find that pokemon")
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", params[1])

	r := rand.New(rand.NewSource(time.Now().Unix()))

	if r.Intn(256) > pokeResp.BaseExperience {
		fmt.Printf("%s was caught!\n", params[1])
		cfg.pokedex[params[1]] = pokeResp
		return nil
	}

	fmt.Printf("%s escaped!\n", params[1])

	return nil
}
