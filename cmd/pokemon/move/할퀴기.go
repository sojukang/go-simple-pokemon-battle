package move

import "go-simple-pokemon-battle/cmd/pokemon/pokemonType"

type 할퀴기 struct {
	power    int
	moveType pokemonType.PokemonType
	name     string
	pp       int
}

func (move *할퀴기) GetPP() int {
	return move.pp
}

func (move *할퀴기) EffectOnAttacker() int {
	return NONE
}

func (move *할퀴기) EffectOnDefender() int {
	return NONE
}

func (move *할퀴기) GetPower() int {
	return move.power
}

func new할퀴기() *할퀴기 {
	return &할퀴기{
		power:    40,
		moveType: pokemonType.TypeNormal,
		name:     "할퀴기",
		pp:       35,
	}
}
