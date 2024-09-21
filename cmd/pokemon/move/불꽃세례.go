package move

import "go-simple-pokemon-battle/cmd/pokemon/pokemonType"

type 불꽃세례 struct {
	power    int
	moveType pokemonType.PokemonType
	name     string
	pp       int
}

func (move *불꽃세례) GetPP() int {
	return move.pp
}

func (move *불꽃세례) EffectOnAttacker() int {
	return NONE
}

func (move *불꽃세례) EffectOnDefender() int {
	return NONE
}

func (move *불꽃세례) GetPower() int {
	return move.power
}

func new불꽃세례() *불꽃세례 {
	return &불꽃세례{
		power:    40,
		moveType: pokemonType.TypeFire,
		name:     "불꽃세례",
		pp:       25,
	}
}
