package move

type Move interface {
	GetPower() int
	EffectOnAttacker() int
	EffectOnDefender() int
	GetPP() int
}

var MOVE_몸통박치기 = new몸통박치기()
var MOVE_째려보기 = new째려보기()
