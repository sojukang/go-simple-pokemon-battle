package move

import "go-simple-pokemon-battle/cmd/pokemon/pokemonType"

type 울음소리 struct {
	power    int
	moveType pokemonType.PokemonType
	name     string
	pp       int
}

func (move *울음소리) GetPP() int {
	return move.pp
}

func (move *울음소리) EffectOnAttacker() int {
	return NONE
}

func (move *울음소리) EffectOnDefender() int {
	return ATTACK_MINUS_1
}

func (move *울음소리) GetPower() int {
	return move.power
}

func new울음소리() *울음소리 {
	return &울음소리{
		power:    0,
		moveType: pokemonType.TypeNormal,
		name:     "울음소리",
		pp:       40,
	}
}
