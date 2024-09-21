package pokemon

import (
	"go-simple-pokemon-battle/cmd/pokemon/move"
	"go-simple-pokemon-battle/cmd/pokemon/pokemonType"
)

type Charmander struct {
	*BasePokemon
	Code string
}

func NewCharmander(code string) *Charmander {
	return &Charmander{
		BasePokemon: NewBasePokemon(
			100,
			65,
			50,
			65,
			pokemonType.TypeFire,
		),
		Code: code,
	}
}

func (attacker Charmander) move(move move.Move, defender *BasePokemon, damageCalculator DamageCalculator) {
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
