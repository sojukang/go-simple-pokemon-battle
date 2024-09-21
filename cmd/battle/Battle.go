package battle

import "go-simple-pokemon-battle/cmd/Pokemon"

type Battle struct {
	pokemonA Pokemon.Pokemon
	pokemonB Pokemon.Pokemon
}

type TurnCommand struct {
	turnType TurnType
}

type TurnResult struct {
	turnType TurnType
	pokemonA Pokemon.Pokemon
	pokemonB Pokemon.Pokemon
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
