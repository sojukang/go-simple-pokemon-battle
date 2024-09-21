package move

import (
	"go-simple-pokemon-battle/cmd/pokemon/pokemonType"
)

type 몸통박치기 struct {
	power    int
	moveType pokemonType.PokemonType
	name     string
	pp       int
}

func (move *몸통박치기) GetPP() int {
	return move.pp
}

func (move *몸통박치기) EffectOnAttacker() int {
	return NONE
}

func (move *몸통박치기) EffectOnDefender() int {
	return NONE
}

func (move *몸통박치기) GetPower() int {
	return move.power
}

func new몸통박치기() *몸통박치기 {
	return &몸통박치기{
		power:    35,
		moveType: pokemonType.TypeNormal,
		name:     "몸통박치기",
		pp:       35,
	}
}
