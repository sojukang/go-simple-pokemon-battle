package pokemon

type BasePokemon struct {
	hp          int
	powerRank   *Rank
	defenseRank *Rank
	power       int
	defense     int
	pokemonType Type
}

func NewBasePokemon(hp int, power int, defense int, pokemonType Type) *BasePokemon {
	return &BasePokemon{hp: hp, powerRank: NewRank(), defenseRank: NewRank(), power: power, defense: defense, pokemonType: pokemonType}
}

func (basePokemon *BasePokemon) TakeDamage(damage int) {
	basePokemon.hp -= damage
}

func (basePokemon *BasePokemon) calculateRankCoefficient(defender *BasePokemon) float32 {
	return basePokemon.powerRank.getCoefficient() / defender.defenseRank.getCoefficient()
}

func (basePokemon *BasePokemon) calculateTypeCoefficient(defender *BasePokemon) float32 {
	isEffective := basePokemon.pokemonType.isEffective(defender.pokemonType)
	isNotEffective := basePokemon.pokemonType.isEffective(defender.pokemonType)

	if isEffective {
		return 2
	}

	if isNotEffective {
		return 0.5
	}

	return 1
}
