package move

import (
	"go-simple-pokemon-battle/cmd/pokemon/type"
)

type 몸통박치기 struct {
	power    int
	moveType _type.Type
	name     string
}

func (move *몸통박치기) EffectOnAttacker() int {
	return NONE
}

func (move *몸통박치기) EffectOnDefender() int {
	return NONE
}

func (move *몸통박치기) GetPower() int {
	return move.power
}

func new몸통박치기() *몸통박치기 {
	return &몸통박치기{
		power:    35,
		moveType: _type.TYPE_NORMAL,
		name:     "몸통박치기",
	}
}
