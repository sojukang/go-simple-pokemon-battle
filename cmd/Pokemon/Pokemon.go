package Pokemon

type Pokemon interface {
}

type Bulbasaur struct {
	*BasePokemon
	Code string
}

func NewBulbasaur(code string) *Bulbasaur {
	return &Bulbasaur{
		BasePokemon: &BasePokemon{hp: 100},
		Code:        code,
	}
}
