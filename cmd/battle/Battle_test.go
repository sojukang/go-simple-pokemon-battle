package battle

import (
	"go-simple-pokemon-battle/cmd/Pokemon"
	"testing"
)

func Test_턴_진행_후_결과를_반환한다(t *testing.T) {
	// given
	testBattle := Battle{
		pokemonA: Pokemon.Bulbasaur{Code: "A"},
		pokemonB: Pokemon.Bulbasaur{Code: "B"},
	}

	// when
	result := testBattle.doTurn(
		TurnCommand{turnType: MOVE},
	)

	// then
	expected := TurnResult{
		turnType: MOVE,
		pokemonA: Pokemon.Bulbasaur{Code: "A"},
		pokemonB: Pokemon.Bulbasaur{Code: "B"},
	}
	if result != expected {
		t.Error("expected: ", expected, ",", "actual: ", result)
	}
}
