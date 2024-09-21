package move

import "go-simple-pokemon-battle/cmd/pokemon/pokemonType"

type 물대포 struct {
	power    int
	moveType pokemonType.PokemonType
	name     string
	pp       int
}

func (move *물대포) GetPP() int {
	return move.pp
}

func (move *물대포) EffectOnAttacker() int {
	return NONE
}

func (move *물대포) EffectOnDefender() int {
	return NONE
}

func (move *물대포) GetPower() int {
	return move.power
}

func new물대포() *물대포 {
	return &물대포{
		power:    40,
		moveType: pokemonType.TypeAqua,
		name:     "물대포",
		pp:       25,
	}
}
