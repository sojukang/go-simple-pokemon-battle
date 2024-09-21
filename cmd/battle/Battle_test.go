package battle

import (
	"go-simple-pokemon-battle/cmd/pokemon"
	"testing"
)

func Test_턴_진행_후_결과를_반환한다(t *testing.T) {
	// given
	testBattle := Battle{
		pokemonA: pokemon.Bulbasaur{Code: "A"},
		pokemonB: pokemon.Bulbasaur{Code: "B"},
	}

	// when
	result := testBattle.doTurn(
		TurnCommand{turnType: MOVE},
	)

	// then
	expected := TurnResult{
		turnType: MOVE,
		pokemonA: pokemon.Bulbasaur{Code: "A"},
		pokemonB: pokemon.Bulbasaur{Code: "B"},
	}
	if result.turnType != MOVE {
		t.Error("expected: ", expected, ",", "actual: ", result)
	}
}
