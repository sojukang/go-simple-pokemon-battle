package move

import (
	"go-simple-pokemon-battle/cmd/pokemon/pokemonType"
)

type 째려보기 struct {
	power    int
	moveType pokemonType.PokemonType
	name     string
	pp       int
}

func (move *째려보기) GetPP() int {
	return move.pp
}

func (move *째려보기) EffectOnAttacker() int {
	return NONE
}

func (move *째려보기) EffectOnDefender() int {
	return DEFENSE_MINUS_1
}

func (move *째려보기) GetPower() int {
	return move.power
}

func new째려보기() *째려보기 {
	return &째려보기{
		power:    0,
		moveType: pokemonType.TypeNormal,
		name:     "째려보기",
		pp:       30,
	}
}
