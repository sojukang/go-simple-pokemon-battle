package Pokemon

import (
	"testing"
)

func Test_포켓몬_코드를_반환한다(t *testing.T) {
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

func Test_포켓몬_데미지를_입는다(t *testing.T) {
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
