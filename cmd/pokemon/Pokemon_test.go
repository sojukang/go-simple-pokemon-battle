package pokemon

import (
	"go-simple-pokemon-battle/cmd/pokemon/move"
	"testing"
)

func Test_코드를_반환한다(t *testing.T) {
	// given
	testPokemon := NewBulbasaur("A", []move.Move{move.MOVE_덩굴채찍})

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
	testPokemon := NewBulbasaur("A", []move.Move{move.MOVE_덩굴채찍})
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
	attacker := NewBulbasaur("A", []move.Move{move.MOVE_몸통박치기})
	defender := NewSquirtle("B", []move.Move{move.MOVE_물대포})
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
	attacker := NewSquirtle("A", []move.Move{move.MOVE_째려보기})
	defender := NewSquirtle("B", []move.Move{move.MOVE_물대포})
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

func Test_현재_기술목록_반환한다(t *testing.T) {
	// given
	pokemon := NewBulbasaur("A", []move.Move{move.MOVE_몸통박치기, move.MOVE_덩굴채찍})

	// when
	allMoves := pokemon.Moves

	// then
	expectedLength := 2
	if len(allMoves) != expectedLength {
		t.Error("expected: ", expectedLength, "actual: ", len(allMoves))
	}
	expected몸통박치기PP := 35
	if allMoves[move.MOVE_몸통박치기] != expected몸통박치기PP {
		t.Error("expected: ", expected몸통박치기PP, "actual: ", allMoves[move.MOVE_몸통박치기])
	}
	expected덩굴채찍PP := 10
	if allMoves[move.MOVE_덩굴채찍] != expected덩굴채찍PP {
		t.Error("expected: ", expected덩굴채찍PP, "actual: ", allMoves[move.MOVE_덩굴채찍])
	}
}

func Test_기술사용시_PP_1_감소한다(t *testing.T) {
	// given
	pokemon := NewBulbasaur("A", []move.Move{move.MOVE_몸통박치기})
	defender := NewSquirtle("B", []move.Move{move.MOVE_물대포})
	initialPP := pokemon.Moves[move.MOVE_몸통박치기]

	// when
	pokemon.move(move.MOVE_몸통박치기, defender.BasePokemon, &TestDamageCalculator{})
	result := pokemon.Moves[move.MOVE_몸통박치기]

	// then
	expected := initialPP - 1
	if result != expected {
		t.Error("expected: ", expected, "actual: ", result)
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
