package pokemon

type Move interface {
	getPower() int
}

var MOVE_몸통박치기 = new몸통박치기()

type 몸통박치기 struct {
	power    int
	moveType Type
	name     string
}

func (move *몸통박치기) getPower() int {
	return move.power
}

func new몸통박치기() *몸통박치기 {
	return &몸통박치기{
		power:    35,
		moveType: TYPE_NORMAL,
		name:     "몸통박치기",
	}
}
