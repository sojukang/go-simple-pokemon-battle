package pokemon

import "go-simple-pokemon-battle/cmd/pokemon/move"

type BaseMoves struct {
	Moves map[move.Move]int
}

func NewBaseMoves(moves []move.Move) *BaseMoves {
	if len(moves) > 4 {
		panic("4개 이상의 기술을 지닐 수 없습니다.")
	}
	return &BaseMoves{Moves: movesToPPMap(moves)}
}

func (baseMoves *BaseMoves) useMove(move move.Move) {
	baseMoves.Moves[move]--
}

func movesToPPMap(moves []move.Move) map[move.Move]int {
	movesByPP := make(map[move.Move]int)
	for _, m := range moves {
		movesByPP[m] = m.GetPP()
	}
	return movesByPP
}

func (baseMoves *BaseMoves) processMove(move move.Move, attacker *BasePokemon, defender *BasePokemon, damageCalculator DamageCalculator) {
	damage := damageCalculator.calculate(
		attacker.power,
		move.GetPower(),
		defender.defense,
		attacker.calculateTypeCoefficient(defender),
		attacker.calculateRankCoefficient(defender),
	)

	attacker.processMoveEffect(move.EffectOnAttacker())
	baseMoves.useMove(move)
	defender.TakeDamage(damage)
	defender.processMoveEffect(move.EffectOnDefender())
}
