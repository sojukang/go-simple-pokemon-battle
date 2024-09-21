package pokemon

type DamageCalculator interface {
	calculate(
		attackerPower int,
		movePower int,
		defenderDefense int,
		typeCoefficient float32,
		rankCoefficient float32,
	) int
}

type DefaultDamageCalculator struct{}

func (calculator *DefaultDamageCalculator) calculate(
	attackerPower int,
	movePower int,
	defenderDefense int,
	typeCoefficient float32,
	rankCoefficient float32,
) int {
	return int(float32(attackerPower) * float32(movePower) / float32(defenderDefense) * typeCoefficient * rankCoefficient)
}
