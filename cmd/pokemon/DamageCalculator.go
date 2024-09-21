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
