package pokemon

import (
	"go-simple-pokemon-battle/cmd/pokemon/move"
	"go-simple-pokemon-battle/cmd/pokemon/pokemonType"
)

type Bulbasaur struct {
	*BasePokemon
	*BaseMoves
	Code string
}

func NewBulbasaur(code string, moves []move.Move) *Bulbasaur {
	return &Bulbasaur{
		BasePokemon: NewBasePokemon(
			100,
			65,
			65,
			45,
			pokemonType.TypeGrass,
		),
		Code:      code,
		BaseMoves: NewBaseMoves(moves),
	}
}

func (pokemon Bulbasaur) move(move move.Move, defender *BasePokemon, damageCalculator DamageCalculator) {
	pokemon.processMove(move, pokemon.BasePokemon, defender, damageCalculator)
}
