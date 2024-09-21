package pokemon

type Pokemon interface {
	move(move Move, defender *BasePokemon, damageCalculator DamageCalculator)
}

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
			TYPE_NORMAL,
		),
		Code: code,
	}
}

func (pokemon Bulbasaur) move(move Move, defender *BasePokemon, damageCalculator DamageCalculator) {
	damage := damageCalculator.calculate(
		pokemon.power,
		move.getPower(),
		defender.defense,
		pokemon.calculateTypeCoefficient(defender),
		pokemon.calculateRankCoefficient(defender),
	)

	defender.TakeDamage(damage)
}
