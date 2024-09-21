package move

import (
	"go-simple-pokemon-battle/cmd/pokemon/type"
)

type 째려보기 struct {
	power    int
	moveType _type.Type
	name     string
}

func (move *째려보기) EffectOnAttacker() int {
	return NONE
}

func (move *째려보기) EffectOnDefender() int {
	return DEFENSE_MINUS_1
}

func (move *째려보기) GetPower() int {
	return move.power
}

func new째려보기() *째려보기 {
	return &째려보기{
		power:    0,
		moveType: _type.TYPE_NORMAL,
		name:     "째려보기",
	}
}
