package move

import "go-simple-pokemon-battle/cmd/pokemon/pokemonType"

type 덩굴채찍 struct {
	power    int
	moveType pokemonType.PokemonType
	name     string
	pp       int
}

func (move *덩굴채찍) GetPP() int {
	return move.pp
}

func (move *덩굴채찍) EffectOnAttacker() int {
	return NONE
}

func (move *덩굴채찍) EffectOnDefender() int {
	return NONE
}

func (move *덩굴채찍) GetPower() int {
	return move.power
}

func new덩굴채찍() *덩굴채찍 {
	return &덩굴채찍{
		power:    35,
		moveType: pokemonType.TypeGrass,
		name:     "덩굴채찍",
		pp:       10,
	}
}
