package pokemon

import (
	"go-simple-pokemon-battle/cmd/pokemon/move"
	"go-simple-pokemon-battle/cmd/pokemon/pokemonType"
)

type Charmander struct {
	*BasePokemon
	*BaseMoves
	Code string
}

func NewCharmander(code string, moves []move.Move) *Charmander {
	return &Charmander{
		BasePokemon: NewBasePokemon(
			100,
			65,
			50,
			65,
			pokemonType.TypeFire,
		),
		Code:      code,
		BaseMoves: NewBaseMoves(moves),
	}
}

func (pokemon Charmander) move(move move.Move, defender *BasePokemon, damageCalculator DamageCalculator) {
	pokemon.processMove(move, pokemon.BasePokemon, defender, damageCalculator)
}
