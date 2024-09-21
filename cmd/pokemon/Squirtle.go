package pokemon

import (
	"go-simple-pokemon-battle/cmd/pokemon/move"
	"go-simple-pokemon-battle/cmd/pokemon/pokemonType"
)

type Squirtle struct {
	*BasePokemon
	*BaseMoves
	Code string
}

func NewSquirtle(code string, moves []move.Move) *Squirtle {
	return &Squirtle{
		BasePokemon: NewBasePokemon(
			100,
			50,
			65,
			43,
			pokemonType.TypeFire,
		),
		Code:      code,
		BaseMoves: NewBaseMoves(moves),
	}
}

func (pokemon Squirtle) move(move move.Move, defender *BasePokemon, damageCalculator DamageCalculator) {
	pokemon.processMove(move, pokemon.BasePokemon, defender, damageCalculator)
}
