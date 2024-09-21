package pokemon

import (
	"go-simple-pokemon-battle/cmd/pokemon/move"
)

type Pokemon interface {
	move(move move.Move, defender *BasePokemon, damageCalculator DamageCalculator)
}
