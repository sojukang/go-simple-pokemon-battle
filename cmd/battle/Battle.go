package battle

import "go-simple-pokemon-battle/cmd/pokemon"

type Battle struct {
	pokemonA pokemon.Pokemon
	pokemonB pokemon.Pokemon
}

type TurnCommand struct {
	turnType TurnType
}

type TurnResult struct {
	turnType TurnType
	pokemonA pokemon.Pokemon
	pokemonB pokemon.Pokemon
}

func (battle *Battle) doTurn(command TurnCommand) TurnResult {
	return TurnResult{
		turnType: command.turnType,
		pokemonA: battle.pokemonA,
		pokemonB: battle.pokemonB,
	}
}

type TurnType string

const (
	MOVE TurnType = "MOVE"
)
