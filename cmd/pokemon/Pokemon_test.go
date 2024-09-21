package pokemon

import (
	"go-simple-pokemon-battle/cmd/pokemon/move"
	"testing"
)

func Test_코드를_반환한다(t *testing.T) {
	// given
	testPokemon := NewBulbasaur("A")

	// when
	result := testPokemon.Code

	// then
	expected := "A"
	if result != expected {
		t.Error("expected: ", expected, ",", "actual: ", result)
	}
}

func Test_데미지를_입는다(t *testing.T) {
	// given
	testPokemon := NewBulbasaur("A")
	initialHp := testPokemon.hp

	// when
	testPokemon.TakeDamage(10)
	result := testPokemon.hp

	// then
	expected := initialHp - 10
	if result != expected {
		t.Error("expected: ", expected, ",", "actual: ", result)
	}
}

func Test_공격기술_사용시_데미지를_입힌다(t *testing.T) {
	// given
	attacker := NewBulbasaur("A")
	defender := NewBulbasaur("B")
	testDamageCalculator := TestDamageCalculator{}
	initialHp := defender.hp

	// when
	attacker.move(move.MOVE_몸통박치기,
		defender.BasePokemon,
		&testDamageCalculator,
	)
	result := defender.hp

	// then
	expected := initialHp - 10
	if result != expected {
		t.Error("expected: ", expected, ",", "actual: ", result)
	}
}

func Test_변화기술_사용시_랭크_변화시킨다(t *testing.T) {
	// given
	attacker := NewBulbasaur("A")
	defender := NewBulbasaur("B")
	damageCalculator := TestDamageCalculator{}
	initialDefenderDefenseRank := defender.defenseRank.value

	// when
	attacker.move(move.MOVE_째려보기,
		defender.BasePokemon,
		&damageCalculator,
	)
	result := defender.defenseRank.value

	// then
	expected := initialDefenderDefenseRank - 1
	if result != expected {
		t.Error("expected: ", expected, ",", "actual: ", result)
	}
}

type TestDamageCalculator struct {
}

func (damageCalculator *TestDamageCalculator) calculate(
	attackerPower int,
	movePower int,
	defenderDefense int,
	typeCoefficient float32,
	rankCoefficient float32,
) int {
	return 10
}
