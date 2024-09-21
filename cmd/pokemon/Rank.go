package pokemon

const maxRank = 2

type Rank struct {
	value int
}

func NewRank() *Rank {
	return &Rank{value: 0}
}

func (rank *Rank) increase() {
	if rank.value < maxRank {
		rank.value++
	}
}

func (rank *Rank) decrease() {
	if rank.value > -maxRank {
		rank.value--
	}
}

func (rank *Rank) getCoefficient() float32 {
	if rank.value < 0 {
		return 2 / float32(2-rank.value)
	} else {
		return float32(2+(rank.value)) / 2
	}
}
