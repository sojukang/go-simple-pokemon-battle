package pokemon

import (
	"go-simple-pokemon-battle/cmd/pokemon/move"
	"go-simple-pokemon-battle/cmd/pokemon/pokemonType"
)

type Squirtle struct {
	*BasePokemon
	Code string
}

func NewSquirtle(code string) *Squirtle {
	return &Squirtle{
		BasePokemon: NewBasePokemon(
			100,
			50,
			65,
			pokemonType.TypeFire,
		),
		Code: code,
	}
}

func (attacker Squirtle) move(move move.Move, defender *BasePokemon, damageCalculator DamageCalculator) {
	damage := damageCalculator.calculate(
		attacker.power,
		move.GetPower(),
		defender.defense,
		attacker.calculateTypeCoefficient(defender),
		attacker.calculateRankCoefficient(defender),
	)

	attacker.processMoveEffect(move.EffectOnAttacker())
	defender.TakeDamage(damage)
	defender.processMoveEffect(move.EffectOnDefender())
}
