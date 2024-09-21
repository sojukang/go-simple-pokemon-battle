package pokemon

import (
	"go-simple-pokemon-battle/cmd/pokemon/move"
	"go-simple-pokemon-battle/cmd/pokemon/pokemonType"
)

type Bulbasaur struct {
	*BasePokemon
	Code string
}

func NewBulbasaur(code string) *Bulbasaur {
	return &Bulbasaur{
		BasePokemon: NewBasePokemon(
			100,
			65,
			65,
			45,
			pokemonType.TypeGrass,
		),
		Code: code,
	}
}

func (attacker Bulbasaur) move(move move.Move, defender *BasePokemon, damageCalculator DamageCalculator) {
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
