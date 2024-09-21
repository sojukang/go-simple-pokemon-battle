package pokemon

import (
	"go-simple-pokemon-battle/cmd/pokemon/move"
	"go-simple-pokemon-battle/cmd/pokemon/type"
)

type BasePokemon struct {
	hp          int
	powerRank   *Rank
	defenseRank *Rank
	power       int
	defense     int
	pokemonType _type.Type
}

func NewBasePokemon(hp int, power int, defense int, pokemonType _type.Type) *BasePokemon {
	return &BasePokemon{hp: hp, powerRank: NewRank(), defenseRank: NewRank(), power: power, defense: defense, pokemonType: pokemonType}
}

func (basePokemon *BasePokemon) TakeDamage(damage int) {
	basePokemon.hp -= damage
}

func (basePokemon *BasePokemon) calculateRankCoefficient(defender *BasePokemon) float32 {
	return basePokemon.powerRank.getCoefficient() / defender.defenseRank.getCoefficient()
}

func (basePokemon *BasePokemon) calculateTypeCoefficient(defender *BasePokemon) float32 {
	isEffective := basePokemon.pokemonType.IsEffective(defender.pokemonType)
	isNotEffective := basePokemon.pokemonType.IsNotEffective(defender.pokemonType)

	if isEffective {
		return 2
	}

	if isNotEffective {
		return 0.5
	}

	return 1
}

func (basePokemon *BasePokemon) processMoveEffect(moveEffect int) {
	switch moveEffect {
	case move.NONE:
	case move.DEFENSE_MINUS_1:
		basePokemon.defenseRank.decrease()
	case move.DEFENSE_PLUS_1:
		basePokemon.defenseRank.increase()
	case move.ATTACK_MINUS_1:
		basePokemon.powerRank.decrease()
	case move.ATTACK_PLUS_1:
		basePokemon.powerRank.increase()
	}
}
