package pokemon

type Rank struct {
	value int
}

func NewRank() *Rank {
	return &Rank{value: 0}
}

func (rank *Rank) increase() {
	rank.value++
}

func (rank *Rank) decrease() {
	rank.value--
}

func (rank *Rank) getCoefficient() float32 {
	return 1 + float32(rank.value)*0.5
}
