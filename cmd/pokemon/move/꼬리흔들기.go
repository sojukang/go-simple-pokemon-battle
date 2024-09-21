package move

import "go-simple-pokemon-battle/cmd/pokemon/pokemonType"

type 꼬리흔들기 struct {
	power    int
	moveType pokemonType.PokemonType
	name     string
	pp       int
}

func (move *꼬리흔들기) GetPP() int {
	return move.pp
}

func (move *꼬리흔들기) EffectOnAttacker() int {
	return NONE
}

func (move *꼬리흔들기) EffectOnDefender() int {
	return DEFENSE_MINUS_1
}

func (move *꼬리흔들기) GetPower() int {
	return move.power
}

func new꼬리흔들기() *꼬리흔들기 {
	return &꼬리흔들기{
		power:    0,
		moveType: pokemonType.TypeNormal,
		name:     "꼬리흔들기",
		pp:       30,
	}
}
